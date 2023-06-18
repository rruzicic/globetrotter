package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id         *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedOn  int                 `json:"createdOn" bson:"created_on"`
	ModifiedOn int                 `json:"modifiedOn" bson:"modified_on"`
	DeletedOn  int                 `json:"deletedOn" bson:"deleted_on"`
}
