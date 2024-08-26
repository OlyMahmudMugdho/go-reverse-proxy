package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	mux := http.NewServeMux()

	target, _ := url.Parse("http://localhost:8082/")

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(target)
			pr.Out.Host = pr.In.Host
		},
	}

	mux.HandleFunc("GET /ping", Ping)
	mux.HandleFunc("GET /", proxy.ServeHTTP)
	mux.HandleFunc("POST /", proxy.ServeHTTP)

	log.Printf("server is running on port %v", 8080)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(`server crashed`)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
