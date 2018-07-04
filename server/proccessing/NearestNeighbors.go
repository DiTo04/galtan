package proccessing

import (
	"math"
	"sort"
)

type Point struct {
	X float64
	Y float64
}

type LabeledPoint struct {
	Point
	Label string
}

type Classifier struct {
	Data []LabeledPoint
}

type distancePoint struct {
	distance float64
	label string
}

func(c *Classifier) Classify(point Point, k int) string {
	distanceChannel := make(chan distancePoint)
	go calculateDistances(c.Data, point, distanceChannel)
	kClosest := saveKNearest(k, distanceChannel)
	votes := mapByLabel(k, kClosest)
	return getHighestVote(votes)

}

func calculateDistances(data []LabeledPoint, point Point, result chan<-distancePoint) {
	for i := 0; i < len(data); i++ {
		result <- calculateDistance(point, data[i])
	}
	close(result)
}

func calculateDistance(a Point, b LabeledPoint) distancePoint {
	return distancePoint{
		label: b.Label,
		distance: math.Sqrt(math.Pow(b.X- a.X, 2) + math.Pow(b.Y- a.Y,2)),
	}
}

func saveKNearest(k int, distanceChannel <-chan distancePoint) []*distancePoint {
	kClosest := make([]*distancePoint, k)
	for distance := range distanceChannel {
		for i := 0; i < k; i++ {
			if kClosest[i] == nil || distance.distance < kClosest[i].distance {
				closerPoint := distance
				kClosest = append(kClosest[:i], append([]*distancePoint{&closerPoint}, kClosest[i:k-1]...)...)
				break
			}
		}
	}
	return kClosest
}

func mapByLabel(k int, kClosest []*distancePoint) map[string]int {
	resultMap := make(map[string]int)
	for i := 0; i < k; i++ {
		closePoint := kClosest[i]
		resultMap[closePoint.label] = resultMap[closePoint.label] + 1
	}
	return resultMap
}

func getHighestVote(labelMap map[string]int) string {
	pl := make(PairList, len(labelMap))
	i := 0
	for k, v := range labelMap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl[0].Key
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
