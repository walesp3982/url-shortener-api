package handler

import (
	"net/http"
	"url-shortener/depends"
)

type redirectRequest struct {
	code string
}
type redirectResponse struct {
	url string
}

func redirect(code string) *redirectResponse {
	store_urls := depends.GetStore()

	url := store_urls.Get(code)

	if url == nil {
		return nil
	}

	return &redirectResponse{
		url: url.URL,
	}

}

func Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if len(code) == 0 {
		generateError(w, http.StatusBadRequest, "Invalid Code")
		return
	}

	redirect_url := redirect(code)

	if redirect_url == nil {
		print("auxilio")
		generateError(w, http.StatusBadRequest, "No found code in the system")
		return
	}
	println(redirect_url.url)
	http.Redirect(w, r, redirect_url.url, http.StatusMovedPermanently)
}
