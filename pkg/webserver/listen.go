package webserver

import "net/http"

func Listen(port string, callback func()) {
	callback()
	http.ListenAndServe(port, nil)
}
