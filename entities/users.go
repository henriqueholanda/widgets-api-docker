package entities

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Avatar      string `json:"gravatar" bson:"gravatar"`
}

type Users []User
