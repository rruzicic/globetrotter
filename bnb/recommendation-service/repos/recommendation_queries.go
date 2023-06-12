package repos

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
)

// test what this returns and whether it can be directly transfered to []models.User
func GetSimilarUsers(user models.User) ([]models.User, error) {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (a:Accommodation)<-[:Reservation]-(this:User {mongoId:$mongoId})-[r1:Review]->(a)<-[r2:Review]-(other:User)" +
		"WHERE (r2.value-1) <= r1.value <= (r2.value+1) RETURN other"
	query_params := map[string]interface{}{
		"mongoId": user.MongoId,
	}

	retval, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, query_params)
		})
	if err != nil {
		return nil, err
	}
	log.Print(retval)

	return retval.([]models.User), nil
}
