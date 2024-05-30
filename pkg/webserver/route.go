package webserver

import "net/http"

func Route(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, handler)
}
