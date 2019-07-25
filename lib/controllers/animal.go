package controllers

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gin-gonic/gin"
	"github.com/pedrocelso/go-rest-service/lib/services/animal"
)

// ResponseObject is a simple mapping object
type ResponseObject map[string]interface{}

// CreateAnimal creates an Animal
func CreateAnimal(c *gin.Context) {
	var animal *animal.Animal
	var err error
	var output *animal.Animal
	ctx := appengine.NewContext(c.Request)

	if err = c.BindJSON(&animal); err == nil {
		if output, err = animal.Create(ctx, animal); err == nil {
			c.JSON(http.StatusOK, ResponseObject{"animal": output})
		}
	}

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// GetAnimal based on its specificName
func GetAnimal(c *gin.Context) {
	var err error
	var output *animal.Animal
	animalSpecificName := c.Param("animalSpecificName")
	ctx := appengine.NewContext(c.Request)

	if output, err = animal.GetBySpecificName(ctx, animalSpecificName); err == nil {
		c.JSON(http.StatusOK, output)
	}
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// GetAnimals Fectch all animals
func GetAnimals(c *gin.Context) {
	var err error
	ctx := appengine.NewContext(c.Request)
	var output []animal.Animal

	if output, err = animal.GetAnimals(ctx); err == nil {
		c.JSON(http.StatusOK, output)
	}

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// UpdateAnimal Updates an animal
func UpdateAnimal(c *gin.Context) {
	var animal *animal.Animal
	var err error
	var output *animal.Animal
	ctx := appengine.NewContext(c.Request)

	if err = c.BindJSON(&animal); err == nil {
		if output, err = animal.Update(ctx, animal); err == nil {
			c.JSON(http.StatusOK, ResponseObject{"animal": output})
		}
	}

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// DeleteAnimal deletes an animal based on its specificName
func DeleteAnimal(c *gin.Context) {
	animalSpecificName := c.Param("animalSpecificName")
	ctx := appengine.NewContext(c.Request)

	err := animal.Delete(ctx, animalSpecificName)

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
	c.JSON(http.StatusOK, ResponseObject{"result": "ok"})
}
