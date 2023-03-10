package models

import "gopkg.in/mgo.v2/bson"

type Model struct {
	Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CreatedOn  int           `json:"createdOn" bson:"created_on"`
	ModifiedOn int           `json:"modifiedOn" bson:"modified_on"`
	DeletedOn  int           `json:"deletedOn" bson:"deleted_on"`
}
