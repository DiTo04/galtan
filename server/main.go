package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"time"
	"github.com/DiTo04/galtan/server/data"
)

var (
	port        = getEnv("PORT", "8080")
	StorageFile = getEnv("FILE_PATH", "./results.json")
)

type ResultStore interface {
	Save(payload data.Payload) error
	GetAll() ([]data.Payload, error)
}

func main() {
	router := mux.NewRouter()
	store := data.NewResultStore(StorageFile)
	router.NewRoute().
		Methods("POST").
		Path("/results").
		HandlerFunc(postResultHandler(store))
	router.HandleFunc("/results", getResultsHandler(store))
	router.HandleFunc("/results/k/{k}", getKNearestNeighbor(store))
	router.HandleFunc("/healthz", allIsOkey)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/static/")))
	http.ListenAndServe("0.0.0.0:" + port, router)
}

func getResultsHandler(store ResultStore) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		allData, err := store.GetAll()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(writer).Encode(allData)
	}
}

func getKNearestNeighbor(store ResultStore) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		c := "c"
		v := "v"
		nrOfPixels := 500
		rows := make([][]string, nrOfPixels, nrOfPixels)
		for _, row := range rows {
			for i := range row {
				if i < nrOfPixels/2 {
					row[i] = v
				} else {
					row[i] = c
				}
			}
		}
		json.NewEncoder(writer).Encode(rows)
	}
}

func allIsOkey(writer http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(writer).Encode("okey")
}

func postResultHandler(resultStore ResultStore) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payLoad := data.Payload{}
		json.NewDecoder(r.Body).Decode(&payLoad)
		payLoad.TimeStamp = data.JsonTime{Time: time.Now()}
		fmt.Printf("%+v\n", payLoad)
		err := resultStore.Save(payLoad)
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

