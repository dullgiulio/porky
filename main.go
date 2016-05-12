package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type dumptransport struct{}

func (t *dumptransport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n\n", dump)
	return response, err
}

func main() {
	listen := flag.String("listen", ":8888", "`HOST:PORT` to listen to")
	to := flag.String("to", "reverse.proxy", "`HOST` to reverse proxy")
	flag.Parse()
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = *to
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n\n", dump)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.Transport = &dumptransport{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(*listen, nil))
}
