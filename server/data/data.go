package data

import (
	"time"
	"fmt"
)

type PoliticalView struct {
	RightLeft float64 `json:"right_left"`
	GalTan    float64 `json:"gal_tan"`
}

type Payload struct {
	PoliticalViews map[string]PoliticalView `json:"political_views"`
	UserChoice     string                   `json:"user_choice"`
	TimeStamp      JsonTime                 `json:"time_stamp"`
	Comment        string                   `json:"comment"`
	Active         bool                     `json:"active"`
}

type JsonTime struct {
	time.Time
}

func (t JsonTime)MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))
	return []byte(stamp), nil
}