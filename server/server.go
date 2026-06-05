package server

import "net/http"

func Run() {
	server := http.NewServeMux()
	AppPath(server)
	http.ListenAndServe(":8080", server)
}
