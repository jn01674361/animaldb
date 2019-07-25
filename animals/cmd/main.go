package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lib/controllers"
)

func init() {
	router := gin.New()
	v1 := router.Group("/v1")

	animals := v1.Group("/animals")
	animals.POST("/", controllers.CreateAnimal)
	animals.GET("/:animalSpecificName", controllers.GetAnimal)
	animals.GET("/", controllers.Getanimals)
	animals.PUT("/:animalSpecificName", controllers.UpdateAnimal)
	animals.DELETE("/:animalSpecificName", controllers.DeleteAnimal)

	http.Handle("/", router)
}