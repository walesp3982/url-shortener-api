package handler

import (
	"errors"
	"net/http"
	"url-shortener/depends"
)

func deleteCode(code string) error {
	store_urls := depends.GetStore()

	if store_urls.Get(code) == nil {
		return errors.New("Url not found")
	}

	store_urls.Delete(code)
	return nil
}

func DeleteCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		generateError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	code := r.PathValue("code")

	if len(code) <= 0 {
		generateError(w, http.StatusBadRequest, "Requiered Code to delete")
	}

	deleteCode(code)

	w.WriteHeader(http.StatusNoContent)
}
