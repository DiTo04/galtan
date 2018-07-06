package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"time"
	"github.com/DiTo04/galtan/server/data"
	"github.com/DiTo04/galtan/server/proccessing"
	"strconv"
)

var (
	port        = getEnv("PORT", "8080")
	StorageFile = getEnv("FILE_PATH", "./results.json")
	nrOfPoints	= 100
	classifier *proccessing.Classifier
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
		k, err := strconv.Atoi(mux.Vars(request)["k"])
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		rows, err := generateBitMap(store, k)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(writer).Encode(rows)
	}
}

func generateBitMap(store ResultStore, k int) ([][]string, error) {
	if classifier == nil {
		newClassifier, err := createClassifier(store)
		if err != nil {
			return nil, err
		}
		classifier = newClassifier
	}
	rows := make([][]string, nrOfPoints)
	outputChannel := make(chan rowWithNr)
	for i := range rows {
		go classifyBitMapRow(i, k, outputChannel)
	}
	for i := 0; i < nrOfPoints; i++ {
		row := <-outputChannel
		rows[row.row] = row.data
	}
	return rows, nil
}

type rowWithNr struct {
	row int
	data []string
}

func classifyBitMapRow(i int, k int, output chan<-rowWithNr)  {
	row := make([]string, nrOfPoints)
	for j := range row {
		point := convertToPoint(i, j)
		row[j] = classifier.Classify(point, k)
	}
	output <- rowWithNr{row: i,data: row }
}

func convertToPoint(i int, j int) proccessing.Point {
	x := 2*float64(j)/float64(nrOfPoints) - 1
	y := 2*float64(i)/float64(nrOfPoints) - 1
	return proccessing.Point{X:x, Y:y}
}

func createClassifier(store ResultStore) (*proccessing.Classifier, error) {
	points, err := store.GetAll();
	if err != nil {
		return nil, err
	}
	labeledPoints := convertPoints(points)
	return &proccessing.Classifier{Data: labeledPoints}, nil
}

func convertPoints(payloads []data.Payload) []proccessing.LabeledPoint {
	result := make([]proccessing.LabeledPoint, len(payloads)*9)
	length := len(payloads[0].PoliticalViews)
	for i, point := range payloads {
		result =  append(result[:i*length],convertPoint(point)...)
	}
	return result
}

func convertPoint(payload data.Payload) []proccessing.LabeledPoint {
	result := make([]proccessing.LabeledPoint, len(payload.PoliticalViews))
	i := 0
	for party, view := range payload.PoliticalViews {
		if party == "person"{
			continue
		}
		result[i] = proccessing.LabeledPoint{Label: party, Point: proccessing.Point{X: view.RightLeft, Y:view.GalTan}}
		i++
	}
	return result
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
		go updateClassifier(resultStore);
	}
}

func updateClassifier(store ResultStore) {
	classifier, _ = createClassifier(store)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

