package controllers

import (
	foodPkg "github.com/agparadiso/nutritionTracker/food"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//IngredientController is the controller type
type IngredientController struct {
	session     *mgo.Session
	foodFetcher foodPkg.Fetcher
}

//NewIngredientController returns a new controller
func NewIngredientController(s *mgo.Session, ff foodPkg.Fetcher) *IngredientController {
	return &IngredientController{
		session:     s,
		foodFetcher: ff,
	}
}

//GetIngredient returns a Ingredient
func (ic *IngredientController) GetIngredient(c *gin.Context) {
	id := c.Params.ByName("id")

	i, err := ic.foodFetcher.FetchIngredient(id, ic.session)
	if err != nil {
		c.JSON(404, gin.H{"error": "Ingredient not found"})
	}
	c.JSON(200, i)
	//http get http://localhost:8080/api/v1/ingredient/593d6f5686ce6452dfe5dc7f
}

//PostIngredient creates a new Ingredient
func (ic *IngredientController) PostIngredient(c *gin.Context) {
	var ingredient foodPkg.Ingredient
	c.Bind(&ingredient)

	// Add an Id
	ingredient.ID = bson.NewObjectId()

	// Write the ingredient to mongo
	ic.session.DB(dbname).C(ingredientTable).Insert(ingredient)

	// Marshal provided interface into JSON structure
	if ingredient.Name != "" {
		c.JSON(201, gin.H{"success": ingredient})
	} else {
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	//http post http://localhost:8080/api/v1/ingredient name=egg protein:=13 carbohydrate:=11 fat:=0
}

// DeleteIngredient removes an existing Ingredient
func (ic *IngredientController) DeleteIngredient(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		c.JSON(404, gin.H{"error": "Ingredient not found"})
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove Ingredient
	if err := ic.session.DB(dbname).C(ingredientTable).RemoveId(oid); err != nil {
		c.JSON(404, gin.H{"error": "Failed to remove ingredient"})
		return
	}

	// Write status
	c.JSON(200, nil)
	//http delete http://localhost:8080/api/v1/ingredient/593d6f3286ce6452dfe5dc7e
}
