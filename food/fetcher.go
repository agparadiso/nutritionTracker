package food

import mgo "gopkg.in/mgo.v2"

type Fetcher interface {
	FetchIngredient(ingredientID string, session *mgo.Session) (*Ingredient, error)
	FetchFood(foodID string, session *mgo.Session) (*Food, error)
	FetchAllFood(session *mgo.Session) ([]Food, error)
}
