package repos

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
)

func CreateAccommodationNode(accommodation models.Accommodation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
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

	return nil
}

func CreateUserNode(user models.User) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "CREATE (:User {name:$name, mongoId:$mongoId})"
	query_params := map[string]interface{}{
		"name":    user.Name,
		"mongoId": user.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func CreateReviewRelationship(review models.Review) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (u:User {mongoId:$userMongoId}) MATCH (a:Accommodation {mongoId:$accommodationMongoId}) CREATE (u)-[r:Review {value:$value, mongoId:$mongoId}]->(a)"
	query_params := map[string]interface{}{
		"userMongoId":          review.UserMongoId,
		"accommodationMongoId": review.AccommodationMongoId,
		"value":                review.Value,
		"mongoId":              review.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func CreateReservationRelationship(reservation models.Reservation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (u:User {mongoId:$userMongoId}) MATCH (a:Accommodation {mongoId:$accommodationMongoId}) CREATE (u)-[r:Reservation {mongoId:$mongoId}]->(a)"
	query_params := map[string]interface{}{
		"userMongoId":          reservation.UserMongoId,
		"accommodationMongoId": reservation.AccommodationMongoId,
		"mongoId":              reservation.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}
