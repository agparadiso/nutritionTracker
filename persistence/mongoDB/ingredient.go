package mongoDB

import (
	"errors"

	foodPkg "github.com/agparadiso/nutritionTracker/food"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbname          = "nutritionTracker"
	foodTable       = "foods"
	ingredientTable = "ingredients"
)

type foodFetcher struct{}

//FetchIngredient from mongoDB
func (ff *foodFetcher) FetchIngredient(ingredientID string, session *mgo.Session) (*foodPkg.Ingredient, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(ingredientID) {
		return nil, errors.New("Ingredient not found")
	}

	// Grab id
	oid := bson.ObjectIdHex(ingredientID)

	// Stub ingredient
	i := foodPkg.Ingredient{}

	// Fetch ingredient
	if err := session.DB(dbname).C(ingredientTable).FindId(oid).One(&i); err != nil {
		return nil, errors.New("Failed to get Ingredient from DB")
	}

	return &i, nil
}

func (ff *foodFetcher) FetchFood(foodID string) (*foodPkg.Food, error) {
	f := foodPkg.Food{}
	return &f, nil
}

func NewIngredientFetcher() foodPkg.Fetcher {
	return &foodFetcher{}
}
