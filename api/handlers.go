package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlesierse/awsdemo-gourl/store"
)

func jsonResult(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func HandleRedirectToUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := store.GetUrlItem(vars["slug"])
	if item == nil {
		http.NotFoundHandler()
	}
	http.Redirect(w, r, item.OriginalUrl, http.StatusMovedPermanently)
}

func HandleCreateUrl(w http.ResponseWriter, r *http.Request) {
	var u Url
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item := store.CreateUrlItem(u.Url)
	jsonResult(w, item)
}
