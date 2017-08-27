package controllers

import (
	"fmt"

	foodPkg "github.com/agparadiso/nutritionTracker/food"
	"github.com/agparadiso/nutritionTracker/persistence/mongoDB"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbname          = "nutritionTracker"
	foodTable       = "foods"
	ingredientTable = "ingredients"
)

//FoodRequest is to unmarshall the request
type FoodRequest struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}

//FoodController is the controller type
type FoodController struct {
	session     *mgo.Session
	foodFetcher foodPkg.Fetcher
}

//NewFoodController returns a new controller
func NewFoodController(s *mgo.Session, ff foodPkg.Fetcher) *FoodController {
	return &FoodController{
		session:     s,
		foodFetcher: ff,
	}
}

//GetFood returns a Food
func (fc *FoodController) GetFood(c *gin.Context) {
	id := c.Params.ByName("id")

	f, err := fc.foodFetcher.FetchFood(id, fc.session)
	if err != nil {
		c.JSON(404, gin.H{"error": "Food not found"})
	}

	c.JSON(200, f)
	//http get http://localhost:8080/api/v1/food/593e4b0686ce646e7bd4907a
}

//GetAllFood returns the list of all Food
func (fc *FoodController) GetAllFood(c *gin.Context) {
	f, err := fc.foodFetcher.FetchAllFood(fc.session)
	if err != nil {
		c.JSON(404, gin.H{"error": "Food not found"})
	}

	c.JSON(200, f)
	//http get http://localhost:8080/api/v1/food/
}

//PostFood creates a new Food
func (fc *FoodController) PostFood(c *gin.Context) {
	var foodRequest = FoodRequest{}
	c.Bind(&foodRequest)

	fmt.Println(foodRequest.Ingredients)
	var food foodPkg.Food
	food.Name = foodRequest.Name

	var ingredient foodPkg.Ingredient

	ic := NewIngredientController(fc.session, mongoDB.NewIngredientFetcher())

	for _, i := range foodRequest.Ingredients {
		// Fetch ingredient
		if err := ic.session.DB(dbname).C(ingredientTable).FindId(bson.ObjectIdHex(i)).One(&ingredient); err != nil {
			c.JSON(404, gin.H{"error": "Failed to get Ingredient from DB"})
			return
		}

		food.Ingredients = append(food.Ingredients, ingredient)
	}

	// Add an Id
	food.ID = bson.NewObjectId()

	// Write the food to mongo
	fc.session.DB(dbname).C(foodTable).Insert(food)

	// Marshal provided interface into JSON structure
	if food.Name != "" {
		c.JSON(201, gin.H{"success": food})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	//http post http://localhost:8080/api/v1/food name=huevoDuro ingredients:='["593d6f5686ce6452dfe5dc7f", "593ec50286ce649d3c417dae"]'
}

// DeleteFood removes an existing Food
func (fc *FoodController) DeleteFood(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		c.JSON(404, gin.H{"error": "Food not found"})
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove food
	if err := fc.session.DB(dbname).C(foodTable).RemoveId(oid); err != nil {
		c.JSON(404, gin.H{"error": "Failed to remove Food"})
		return
	}

	// Write status
	c.JSON(200, nil)
	//http delete http://localhost:8080/api/v1/ingredient/593d6f3286ce6452dfe5dc7e
}
