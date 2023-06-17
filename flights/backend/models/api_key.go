package models

import "time"

type APIKey struct {
	Key        string    `bson:"key" json:"key"`
	Expiration time.Time `bson:"expiration" json:"expiration"`
}
