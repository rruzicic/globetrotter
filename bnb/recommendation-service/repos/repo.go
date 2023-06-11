package repos

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
)

func CreateAccommodation(accommodation models.Accommodation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	log.Println("session mad")
	defer session.Close()

	cypher_query := "CREATE (:Accommodation {name:$name, location:$location, price:$price, mongoId:$mongoId})"
	query_params := map[string]interface{}{
		"name":     accommodation.Name,
		"location": accommodation.Location,
		"price":    accommodation.Price,
		"mongoId":  accommodation.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}
	log.Print("Query done")

	return nil
}
