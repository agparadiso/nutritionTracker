package controllers

import (
	"github.com/agparadiso/nutritionTracker/models"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//FoodRequest is to unmarshall the request
type FoodRequest struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//FoodController is the controller type
type FoodController struct {
	session *mgo.Session
}

//NewFoodController returns a new controller
func NewFoodController(s *mgo.Session) *FoodController {
	return &FoodController{session: s}
}

//GetFood returns a Food
func (fc *FoodController) GetFood(c *gin.Context) {
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		c.JSON(404, gin.H{"error": "Food not found"})
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub Food
	f := models.Food{}

	// Fetch Food
	if err := fc.session.DB("nutritionTracker").C("foods").FindId(oid).One(&f); err != nil {
		c.JSON(404, gin.H{"error": "Failed to get Food from DB"})
		return
	}
	c.JSON(200, f)
	//http get http://localhost:8080/api/v1/food/593e4b0686ce646e7bd4907a
}

//PostFood creates a new Food
func (fc *FoodController) PostFood(c *gin.Context) {
	var foodRequest = FoodRequest{}
	c.Bind(&foodRequest)

	var food models.Food
	food.Name = foodRequest.Name

	var ingredient models.Ingredient

	ic := NewIngredientController(fc.session)
	// Fetch ingredient
	if err := ic.session.DB("nutritionTracker").C("ingredients").FindId(bson.ObjectIdHex(foodRequest.Ingredients)).One(&ingredient); err != nil {
		c.JSON(404, gin.H{"error": "Failed to get Ingredient from DB"})
		return
	}

	food.Ingredients = append(food.Ingredients, ingredient)

	// Add an Id
	food.ID = bson.NewObjectId()

	// Write the food to mongo
	fc.session.DB("nutritionTracker").C("foods").Insert(food)

	// Marshal provided interface into JSON structure
	if food.Name != "" {
		c.JSON(201, gin.H{"success": food})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	//http post http://localhost:8080/api/v1/food name=huevoDuro ingredients=593d6f5686ce6452dfe5dc7f
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
	if err := fc.session.DB("nutritionTracker").C("foods").RemoveId(oid); err != nil {
		c.JSON(404, gin.H{"error": "Failed to remove Food"})
		return
	}

	// Write status
	c.JSON(200, nil)
	//http delete http://localhost:8080/api/v1/ingredient/593d6f3286ce6452dfe5dc7e
}