package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	//Model
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string        `json:"firstName" bson:"first_name"`
	LastName  string        `json:"lastName" bson:"last_name"`
	EMail     string        `json:"email" bson:"email"`
	//Address   Address
}
