package animal

import (
	"fmt"

	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	index           = `Animal`
	invalidAnimalData = `error: invalid Animal data`
)

// Animal defines animal attributes
type Animal struct {
    Name  string  `json:"name"`
	SpecificName  string `json:"specific_name"`
	//Class string `json:"class"`
    //Habitat string `json:"habitat"`
}

// Create an animal
func Create(c context.Context, animal *Animal) (*Animal, error) {
	var output *Animal
	if animal == nil || animal.SpecificName == `` {
		return nil, fmt.Errorf(invalidAnimalData)
	}

	output, _ = GetBySpecificName(c, animal.SpecificName)

	if output == nil {
		key := datastore.NewKey(c, index, animal.SpecificName, 0, nil)
		insKey, err := datastore.Put(c, key, animal)

		if err != nil {
			log.Errorf(c, "ERROR INSERTING Animal: %v", err.Error())
			return nil, err
		}

		output, err = GetBySpecificName(c, insKey.StringID())
		if err != nil {
			log.Errorf(c, "ERROR GETTING Animal OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	log.Infof(c, "Animal was previously saved: %v", animal.SpecificName)
	return output, nil
}

// GetBySpecificName an Animal based on its SpecificName
func GetBySpecificName(c context.Context, SpecificName string) (*Animal, error) {
	if SpecificName == `` {
		return nil, fmt.Errorf(invalidAnimalData)
	}
	AnimalKey := datastore.NewKey(c, index, SpecificName, 0, nil)
	var animal Animal
	err := datastore.Get(c, AnimalKey, &animal)

	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(`Animal '%v' not found`, SpecificName)
		}
		return nil, err
	}
	return &animal, nil
}

// GetAnimals Fetches all Animals
func GetAnimals(c context.Context) ([]Animal, error) {
	var output []Animal
	q := datastore.NewQuery(index)
	_, err := q.GetAll(c, &output)

	if err != nil {
		log.Errorf(c, "error fetching all Animals")
		return nil, err
	}

	if len(output) <= 0 {
		return nil, fmt.Errorf("no Animals found")
	}
	return output, nil
}

// Update Animal data
func Update(c context.Context, animal *Animal) (*Animal, error) {
	if animal == nil || animal.SpecificName == `` {
		return nil, fmt.Errorf(invalidAnimalData)
	}

	output, _ := GetBySpecificName(c, animal.SpecificName)
	if output != nil {
		key := datastore.NewKey(c, index, animal.SpecificName, 0, nil)
		insKey, err := datastore.Put(c, key, animal)

		if err != nil {
			log.Errorf(c, "ERROR UPDATING Animal: %v", err.Error())
			return nil, err
		}

		output, err = GetBySpecificName(c, insKey.StringID())
		if err != nil {
			log.Errorf(c, "ERROR GETTING Animal OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	return nil, fmt.Errorf(`Animal '%v' not found`, animal.SpecificName)
}

// Delete an Animal based on its SpecificName.
func Delete(c context.Context, SpecificName string) error {
	var output *Animal
	output, _ = GetBySpecificName(c, SpecificName)

	if output != nil {
		log.Infof(c, "Deleting Animal: %v", SpecificName)
		key := datastore.NewKey(c, index, SpecificName, 0, nil)
		err := datastore.Delete(c, key)

		if err != nil {
			log.Errorf(c, "ERROR DELETING Animal: %v", err.Error())
			return err
		}
		return nil
	}
	return fmt.Errorf("Animal '%v' don't exist on the database", SpecificName)
}