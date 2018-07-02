package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"time"
)

var (
	port        = getEnv("PORT", "8080")
	StorageFile = getEnv("FILE_PATH", "./results.json")
)

type politicalView struct {
	RightLeft float32 `json:"right_left"`
	GalTan    float32 `json:"gal_tan"`
}

type payload struct {
	PoliticalViews map[string]politicalView `json:"political_views"`
	UserChoice     string                   `json:"user_choice"`
	TimeStamp      JSONTime					`json:"time_stamp"`
	Comment		   string 					`json:"comment"`
	Active 		   bool 					`json:"active"`
	
}

type JSONTime struct {
	time.Time
}

func (t JSONTime)MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))
	return []byte(stamp), nil
}

type ResultStore interface {
	save(payload payload) error
	getAll() ([]payload, error)
}

func main() {
	router := mux.NewRouter()
	store := NewResultStore(StorageFile)
	router.NewRoute().
		Methods("POST").
		Path("/results").
		HandlerFunc(postResultHandler(store))
	router.HandleFunc("/results", getResultsHandler(store))
	router.HandleFunc("/healthz", allIsOkey)
	http.ListenAndServe("0.0.0.0:" + port, router)
}

func getResultsHandler(store ResultStore) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.getAll()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(writer).Encode(data)
	}
}

func allIsOkey(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("okey")
}

func postResultHandler(resultStore ResultStore) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payLoad := payload{}
		json.NewDecoder(r.Body).Decode(&payLoad)
		payLoad.TimeStamp = JSONTime{time.Now()}
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

