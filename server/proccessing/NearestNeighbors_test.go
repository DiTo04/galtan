package proccessing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestCalculateDistance(t *testing.T) {
	// Given
	data := []LabeledPoint{
		{Point: Point{1,1}, Label:"plus"},
	}
	distanceChannel := make(chan distancePoint)
	// When
	go calculateDistances(data, Point{0, 0}, distanceChannel)
	// Then
	distance := <- distanceChannel
	assert.Equal(t, distance.distance,math.Sqrt(2))
	assert.Equal(t, "plus", distance.label)
}

func TestSaveKNearest(t *testing.T) {
	// Given
	data := []LabeledPoint{
		{Point: Point{1,0}, Label:"plus"},
		{Point: Point{0,1}, Label:"plus"},
		{Point: Point{1,1}, Label:"plus"},
		{Point: Point{-1, 0}, Label:"minus"},
		{Point: Point{0, -1}, Label:"minus"},
		{Point: Point{-1, -1}, Label:"minus"},
	}
	channel := make(chan distancePoint)
	go fillChannel(channel, data)
	// When
	result := saveKNearest(3, channel)

	// Then
	assert.Equal(t, 3, len(result))
	for i := 0; i < 3; i++ {
		assert.Equal(t, &distancePoint{label:data[i].Label, distance:float64(i)}, result[i])
	}

}

func fillChannel(channel chan<- distancePoint, data []LabeledPoint) {
	for i := 0; i < 6; i++ {
		channel <- distancePoint{
			label: data[i].Label,
			distance: float64(i),
		}
	}
	close(channel)
}

func TestClassify(t *testing.T) {
	// Given
	data := []LabeledPoint{
		{Point: Point{1,0}, Label:"plus"},
		{Point: Point{0,1}, Label:"plus"},
		{Point: Point{1,1}, Label:"plus"},
		{Point: Point{-1, 0}, Label:"minus"},
		{Point: Point{0, -1}, Label:"minus"},
		{Point: Point{-1, -1}, Label:"minus"},
	}
	target := &Classifier{Data:data}

	// When
	result := target.Classify(Point{0.5, 0.5}, 3)

	// Then
	assert.Equal(t, "plus", result)
}



