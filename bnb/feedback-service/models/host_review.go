package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HostReview struct {
	Model  `bson:",inline"`
	Rating int                 `json:"rating" bson:"rating"`
	UserId *primitive.ObjectID `json:"userId" bson:"user_id"` //reviewer
	HostId *primitive.ObjectID `json:"hostId" bson:"host_id"` //reviewed
}
