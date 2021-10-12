package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	re, err := url.Parse("https://www.messenger.com")
	if err != nil {
		panic(err)
	}
	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r)
			r.Host = re.Host
			w.Header().Set("X-Proxy", "suka blyat")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(re)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
