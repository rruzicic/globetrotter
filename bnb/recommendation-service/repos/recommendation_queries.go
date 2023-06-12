package repos

import (
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

	userRecords, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			retval, err := tx.Run(cypher_query, query_params)
			if err != nil {
				return nil, err
			}

			return retval.Collect()
		})
	if err != nil {
		return nil, err
	}

	users := []models.User{}
	for _, record := range userRecords.([]*neo4j.Record) {
		user_map := record.Values[0].(neo4j.Node).GetProperties()
		// could have named all properties the same as they are in the models structs
		// to be able to then marshall a json based on the map from above, then to unmarshall that json into a struct
		// but thats too much of a hassle and a bit unsafe i think
		users = append(users, models.User{Name: user_map["name"].(string), MongoId: user_map["mongoId"].(string)})
	}

	return users, nil
}
