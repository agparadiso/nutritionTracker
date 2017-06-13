package food

import mgo "gopkg.in/mgo.v2"

type Fetcher interface {
	FetchIngredient(ingredientID string, session *mgo.Session) (*Ingredient, error)
	FetchFood(foodID string) (*Food, error)
}
