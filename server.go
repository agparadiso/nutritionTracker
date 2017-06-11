package main

import (
	"github.com/Typeform/andorra/pkg/log"
	"github.com/agparadiso/nutritionTracker/controllers"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

// Standard library packages

func main() {
	r := gin.Default()
	ic := controllers.NewIngredientController(getSession())

	v1 := r.Group("api/v1")
	{
		v1.GET("/ingredient/:id", ic.GetIngredient)
		v1.POST("/ingredient", ic.PostIngredient)
		v1.DELETE("/ingredient/:id", ic.DeleteIngredient)
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
