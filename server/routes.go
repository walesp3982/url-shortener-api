package server

import (
	"net/http"
	"url-shortener/handler"
)

func AppPath(server *http.ServeMux) {
	server.HandleFunc("POST /url", handler.Shorten)
	server.HandleFunc("GET /url", handler.GetAllUrl)
	server.HandleFunc("DELETE /url/{code}", handler.DeleteCode)
	server.HandleFunc("GET /{code}", handler.Redirect)
}
