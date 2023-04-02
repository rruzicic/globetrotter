package models

import "time"

type API_Key struct {
	Value    string    `json:"value" bson:"value"`
	Duration time.Time `json:"duration" bson:"duration"`
}
