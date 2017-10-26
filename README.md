# porky - Minimal dumping proxy

Porky is a barebone proxy that dumps to stdout everything that passes by.

After compiling with `go install`, run as follows:
```
$ porky -to <SERVER-TO-PROXY> -listen 127.0.0.1:8888
```

Use 127.0.0.1:8888 as endpoint in your development application and see the exchanged
HTTP communication on screen.

## TODO

Porky doesn't yet support TLS.
