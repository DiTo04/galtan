package main

import (
	"testing"
	"github.com/DiTo04/galtan/server/data"
)

func TestClassification(t *testing.T) {
	//Given
	store := data.NewResultStore("./results.json")
	//when
	generateBitMap(store, 5)

	// Then

}
