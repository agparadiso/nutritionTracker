package main

import (
	"github.com/Typeform/andorra/pkg/log"
	"github.com/agparadiso/nutritionTracker/controllers"
	"github.com/agparadiso/nutritionTracker/persistence/mongoDB"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

// Standard library packages

func main() {
	r := gin.Default()
	ic := controllers.NewIngredientController(getSession(), mongoDB.NewIngredientFetcher())
	fc := controllers.NewFoodController(getSession())

	v1 := r.Group("api/v1")
	{
		v1.GET("/ingredient/:id", ic.GetIngredient)
		v1.POST("/ingredient", ic.PostIngredient)
		v1.DELETE("/ingredient/:id", ic.DeleteIngredient)

		v1.GET("/food/:id", fc.GetFood)
		v1.POST("/food", fc.PostFood)
		v1.DELETE("/food/:id", fc.DeleteFood)
	}

	r.Run(":8080")
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		log.Errorf("failed to get mongo running")
		panic(err)
	}
	return s
}
