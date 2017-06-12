package models

import "gopkg.in/mgo.v2/bson"

type Food struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Ingredients []Ingredient  `json:"ingredients" bson:"ingredients"`
}
