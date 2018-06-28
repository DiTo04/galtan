package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
)

var (
	PORT = getEnv("PORT", "8080")
	STORAGE_FILE = getEnv("FILE_PATH", "./results.json")
)

type politicalView struct {
	RightLeft float32 `json:"right_left"`
	GalTan    float32 `json:"gal_tan"`
}

type payload struct {
	PoliticalViews map[string]politicalView `json:"political_views"`
	UserChoice     string                   `json:"user_choice"`
}

func main() {
	router := mux.NewRouter()
	store := NewResultStore(STORAGE_FILE)
	router.NewRoute().
		Methods("POST").
		Path("/results").
		HandlerFunc(postResultHandler(store))
	router.HandleFunc("/healthz", allIsOkey)
	http.ListenAndServe("0.0.0.0:" + PORT, router)
}

func allIsOkey(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("okey")
}

type ResultStore interface {
	save(payload payload) error
}

func postResultHandler(resultStore ResultStore) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payLoad := payload{}
		json.NewDecoder(r.Body).Decode(&payLoad)
		fmt.Printf("%+v\n", payLoad)
		err := resultStore.save(payLoad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

