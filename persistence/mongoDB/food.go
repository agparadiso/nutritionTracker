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

func (ff *foodFetcher) FetchFood(foodID string, session *mgo.Session) (*foodPkg.Food, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(foodID) {
		return nil, errors.New("Food not found")
	}

	// Grab id
	oid := bson.ObjectIdHex(foodID)

	// Stub Food
	f := foodPkg.Food{}

	// Fetch Food
	if err := session.DB(dbname).C(foodTable).FindId(oid).One(&f); err != nil {
		return nil, errors.New("Failed to get Food from DB")
	}

	return &f, nil
}

func (ff *foodFetcher) FetchAllFood(session *mgo.Session) ([]foodPkg.Food, error) {
	// Stub Food
	f := []foodPkg.Food{}

	// Fetch Food
	if err := session.DB(dbname).C(foodTable).Find(nil).All(&f); err != nil {
		return nil, errors.New("Failed to get all Food from DB")
	}

	return f, nil
}

func NewIngredientFetcher() foodPkg.Fetcher {
	return &foodFetcher{}
}
