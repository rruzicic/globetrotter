package repos

import (
	"context"
	"log"

	"github.com/rruzicic/globetrotter/bnb/search-service/models"
	"gopkg.in/mgo.v2/bson"
)

func GetAll() []models.Accommodation{
	result := []models.Accommodation{}
	cursor, err := accommodationCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Panic("could not find users err: ", err.Error())
		return []models.Accommodation{}
	}
	cursor.All(context.TODO(), &result)
	return result
}