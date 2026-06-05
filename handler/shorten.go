package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/depends"
	"url-shortener/shortener"
	"url-shortener/store"
)

type shortenRequest struct {
	Url string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"url"`
}

func getAllUrlService() []store.Shortener {
	store_urls := depends.GetStore()

	shorteners := store_urls.GetAll()

	return shorteners
}

func GetAllUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		generateError(w, http.StatusMethodNotAllowed, "Method not Allowed")
		return
	}

	// Working only method get
	urls := getAllUrlService()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(urls)
}

func shortenService(request shortenRequest) (*shortenResponse, error) {
	store_urls := depends.GetStore()

	code := shortener.ShortUrl(request.Url, 8)

	if code == nil {
		return nil, errors.New("No se pudo generar el code")
	}

	exist_code := store_urls.Get(*code)

	if exist_code != nil {
		if exist_code.URL == request.Url {
			return nil, errors.New("Url ya registrada anteriormente")
		}
		return nil, errors.New("Código ya existente")
	}
	store_urls.Save(store.Shortener{
		Code: *code,
		URL:  request.Url,
	})

	return &shortenResponse{
		Code: *code,
	}, nil
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	setHeaderToJson(w)

	if r.Method != http.MethodPost {
		generateError(w, http.StatusMethodNotAllowed, "Method no allowed")
		return
	}

	defer r.Body.Close()

	var request shortenRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if !isValidUrl(request.Url) {
		generateError(w, http.StatusBadRequest, "url inválido error")
		return
	}

	if err != nil {
		generateError(w, http.StatusBadRequest, "Json inválido")
		return
	}

	response, err := shortenService(request)

	if err != nil {
		generateError(w, http.StatusBadRequest, "No se puedo crear el code")
		return

	}
	json.NewEncoder(w).Encode(response)
}
