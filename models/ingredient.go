package models

import "gopkg.in/mgo.v2/bson"

type Ingredient struct {
	ID           bson.ObjectId `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Protein      int           `json:"protein" bson:"protein"`
	Carbohydrate int           `json:"carbohydrate" bson:"carbohydrate"`
	Fat          int           `json:"fat" bson:"fat"`
}
