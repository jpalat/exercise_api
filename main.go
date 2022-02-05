package main

import (
	"jpalat/exercise_api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := models.ConnectDatabase()
	checkErr(err)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("exercises", getExercise)
		// v1.GET("person/:id", getPersonById)
		// v1.POST("person", addPerson)
		// v1.PUT("person/:id", updatePerson)
		// v1.DELETE("person/:id", deletePerson)
		// v1.OPTIONS("person", options)
	}

	r.Run()
}

func getExercise(c *gin.Context) {
	persons, err := models.GetExercises(10)
	checkErr(err)

	if persons == nil {
		c.JSON(http.StatusOK, gin.H{"error": "No records found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
}
