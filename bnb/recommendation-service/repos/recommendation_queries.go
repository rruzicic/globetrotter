package repos

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
)

func GetSimilarUsers(user models.User) ([]models.User, error) {
	// use this query as entry to recommendations
	// it returns the users most similar to this one
	// then use them to get the highest rated accommodations that this user would "like"
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

func GetHighlyRatedAccommodationsOfUserGroup(users []models.User) ([]models.Accommodation, error) {
	// use this query to get the accommodations that a certain user might want to see
	// first use the get similar users and then use those users in this query
	// then use these accommodations for the filter recent lowly rated accommodations query
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	userIdList := []string{}
	for _, user := range users {
		userIdList = append(userIdList, user.MongoId)
	}

	cypher_query := "MATCH (u:User)-[:Reservation]->(a:Accommodation)<-[r:Review]-(u) WHERE u.mongoId IN $userMongoIdList AND r.value IN [4, 5] RETURN a"
	query_params := map[string]interface{}{
		"userMongoIdList": userIdList,
	}

	accommodationRecords, err := session.
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

	accommodations := []models.Accommodation{}
	for _, record := range accommodationRecords.([]*neo4j.Record) {
		accommodation_map := record.Values[0].(neo4j.Node).GetProperties()
		accommodation := models.Accommodation{
			Name:     accommodation_map["name"].(string),
			Location: accommodation_map["location"].(string),
			Price:    accommodation_map["price"].(float32),
			MongoId:  accommodation_map["mongoId"].(string),
		}
		accommodations = append(accommodations, accommodation)
	}

	return accommodations, nil
}

func FilterRecentLowlyRatedAccommodations(accommodations []models.Accommodation) ([]models.Accommodation, error) {
	// use this query to filter accommodations that were rated below 4 more than 5 times in the past 3 months
	// then use the result of this query in the orderer query to get the best 10 by price
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	accommodationIds := []string{}
	for _, accommodation := range accommodations {
		accommodationIds = append(accommodationIds, accommodation.MongoId)
	}

	cypher_query := "MATCH (u:User)-[re:Reservation]->(a:Accommodation)<-[r:Review]-(u) " +
		"WHERE a.mongoId IN $accommodationMongoIdList" +
		"WITH a, re, collect(r.value) AS reviews " +
		"WHERE size([review IN reviews WHERE review < 3]) < 5 AND re.reservationEnd >= datetime($threeMonthsAgo)" +
		"RETURN a"
	query_params := map[string]interface{}{
		"accommodationMongoIdList": accommodationIds,
		"threeMonthsAgo":           time.Now().AddDate(0, -3, 0).Format(time.RFC3339),
	}

	accommodationRecords, err := session.
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

	filtered_accommodations := []models.Accommodation{}
	for _, record := range accommodationRecords.([]*neo4j.Record) {
		accommodation_map := record.Values[0].(neo4j.Node).GetProperties()
		accommodation := models.Accommodation{
			Name:     accommodation_map["name"].(string),
			Location: accommodation_map["location"].(string),
			Price:    accommodation_map["price"].(float32),
			MongoId:  accommodation_map["mongoId"].(string),
		}
		filtered_accommodations = append(filtered_accommodations, accommodation)
	}

	return filtered_accommodations, nil
}

func GetTenLowestPricedAccommodations(accommodations []models.Accommodation) ([]models.Accommodation, error) {
	// use this query to get the 10 lowest priced accommodations out of the passed ones
	// use these 10 to recommend to the user
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	accommodationIds := []string{}
	for _, accommodation := range accommodations {
		accommodationIds = append(accommodationIds, accommodation.MongoId)
	}

	cypher_query := "MATCH (a:Accommodation)<-[r:Review]-(:User)" +
		"WHERE a.mongoId IN $accommodationMongoIdList" +
		"RETURN DISTINCT a" +
		"ORDER BY a.price DESC" +
		"LIMIT 10"
	query_params := map[string]interface{}{
		"accommodationMongoIdList": accommodationIds,
	}

	accommodationRecords, err := session.
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

	ordered_accommodations := []models.Accommodation{}
	for _, record := range accommodationRecords.([]*neo4j.Record) {
		accommodation_map := record.Values[0].(neo4j.Node).GetProperties()
		accommodation := models.Accommodation{
			Name:     accommodation_map["name"].(string),
			Location: accommodation_map["location"].(string),
			Price:    accommodation_map["price"].(float32),
			MongoId:  accommodation_map["mongoId"].(string),
		}
		ordered_accommodations = append(ordered_accommodations, accommodation)
	}

	return ordered_accommodations, nil
}

func GetRecommendedAccommodations(user models.User) ([]models.Accommodation, error) {
	users, err := GetSimilarUsers(user)
	if err != nil {
		return nil, err
	}

	highly_rated_accommodations, err := GetHighlyRatedAccommodationsOfUserGroup(users)
	if err != nil {
		return nil, err
	}

	filtered_accommodations, err := FilterRecentLowlyRatedAccommodations(highly_rated_accommodations)
	if err != nil {
		return nil, err
	}

	accommodations, err := GetTenLowestPricedAccommodations(filtered_accommodations)
	if err != nil {
		return nil, err
	}

	return accommodations, nil
}
