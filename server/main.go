package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
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

	router.NewRoute().
		Methods("POST").
		Path("/results").
		HandlerFunc(handlePostResult)
	router.HandleFunc("/healthz", allIsOkey)

	http.ListenAndServe("0.0.0.0:8080", router)
}
func allIsOkey(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("okey")
}


func handlePostResult(writer http.ResponseWriter, request *http.Request) {
	payLoad := payload{}
	json.NewDecoder(request.Body).Decode(&payLoad)
	fmt.Printf("%+v\n", payLoad)
}

