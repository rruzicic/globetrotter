package repos

import (
	"time"

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

	cypher_query := "MATCH (u:User {mongoId:$userMongoId}) MATCH (a:Accommodation {mongoId:$accommodationMongoId}) CREATE (u)-[r:Reservation {reservationEnd:datetime($reservationEnd), mongoId:$mongoId}]->(a)"
	query_params := map[string]interface{}{
		"userMongoId":          reservation.UserMongoId,
		"accommodationMongoId": reservation.AccommodationMongoId,
		"reservationEnd":       reservation.ReservationEnd.Format(time.RFC3339),
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

func DeleteAccommodationNode(accommodation models.Accommodation) error {
	// also deletes all the relationships this node had
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (a:Accommodation {mongoId=$mongoId}) DETACH DELETE"
	query_params := map[string]interface{}{
		"mongoId": accommodation.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func DeleteUserNode(user models.User) error {
	// also deletes all the relationships this node had
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (a:User {mongoId=$mongoId}) DETACH DELETE"
	query_params := map[string]interface{}{
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

func DeleteReviewRelationship(review models.Review) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH ()-[r:Review {mongoId:$mongoId}]->() DELETE r"
	query_params := map[string]interface{}{
		"mongoId": review.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func DeleteReservationRelationship(reservation models.Reservation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH ()-[r:Reservation {mongoId:$mongoId}]->() DELETE r"
	query_params := map[string]interface{}{
		"mongoId": reservation.MongoId,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func UpdateAccommodationNode(accommodation models.Accommodation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (a:Accommodation {mongoId:$mongoId})" +
		"SET a.name=$name, a.location=$location, a.price=$price"
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

func UpdateUserNode(user models.User) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (u:User {mongoId:$mongoId}) SET u.name=$name"
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

func UpdateReviewRelationship(review models.Review) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH ()-[r:Review {mongoId:$mongoId}]->() SET r.value=$value"
	query_params := map[string]interface{}{
		"mongoId": review.MongoId,
		"value":   review.Value,
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}

func UpdateReservationRelationship(reservation models.Reservation) error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH ()-[r:Reservation {mongoId:$mongoId}]->() SET r.reservationEnd=datetime($reservationEnd)"
	query_params := map[string]interface{}{
		"mongoId":        reservation.MongoId,
		"reservationEnd": reservation.ReservationEnd.Format(time.RFC3339),
	}

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		}); err != nil {
		return err
	}

	return nil
}
