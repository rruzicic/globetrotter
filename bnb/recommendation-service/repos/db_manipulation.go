package repos

import (
	"strconv"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	grpcclient "github.com/rruzicic/globetrotter/bnb/recommendation-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/recommendation-service/models"
)

func InitDBData() {
	// Initializes the database by pulling all currently available data from all the services
	// ALSO DROPS THE WHOLE BASE BEFORE DOING SO
	DropDB()

	accommodations, _ := grpcclient.GetAllAccommodations()
	for _, accommodation := range accommodations {
		CreateAccommodationNode(accommodation)
	}

	users, _ := grpcclient.GetAllUsers()
	for _, user := range users {
		CreateUserNode(user)
	}

	reviews, _ := grpcclient.GetAllReviews()
	for _, review := range reviews {
		CreateReviewRelationship(review)
	}

	reservations, _ := grpcclient.GetAllReservations()
	for _, reservation := range reservations {
		CreateReservationRelationship(reservation)
	}
}

func DropDB() error {
	session := neo4jDriver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	cypher_query := "MATCH (n) DETACH DELETE n"

	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return tx.Run(cypher_query, map[string]interface{}{})
		}); err != nil {
		return err
	}

	return nil
}

func LoadMockData() {
	for i := 0; i < 10; i++ {
		curr_str := strconv.Itoa(i)
		prev_str := strconv.Itoa(9 - i)
		nxt_str := strconv.Itoa((9 + i) % 10)

		CreateAccommodationNode(models.Accommodation{Name: "acc" + curr_str, Location: "loc" + curr_str, Price: float32(100 + i%5), MongoId: curr_str})
		CreateUserNode(models.User{Name: "user" + curr_str, MongoId: curr_str})

		CreateReviewRelationship(models.Review{Value: 2 + i%4, MongoId: curr_str, UserMongoId: curr_str, AccommodationMongoId: curr_str})
		CreateReviewRelationship(models.Review{Value: 2 + i%4, MongoId: "10" + curr_str, UserMongoId: curr_str, AccommodationMongoId: prev_str})
		CreateReviewRelationship(models.Review{Value: 2 + i%4, MongoId: "20" + curr_str, UserMongoId: curr_str, AccommodationMongoId: nxt_str})

		CreateReservationRelationship(models.Reservation{MongoId: curr_str, UserMongoId: curr_str, AccommodationMongoId: curr_str, ReservationEnd: time.Now()})
		CreateReservationRelationship(models.Reservation{MongoId: "10" + curr_str, UserMongoId: curr_str, AccommodationMongoId: prev_str, ReservationEnd: time.Now()})
		CreateReservationRelationship(models.Reservation{MongoId: "20" + curr_str, UserMongoId: curr_str, AccommodationMongoId: nxt_str, ReservationEnd: time.Now()})
	}
	for _, id := range []int{1, 3, 5, 100, 103, 105, 200, 203, 205} {
		UpdateReservationRelationship(models.Reservation{MongoId: strconv.Itoa(id), UserMongoId: "", AccommodationMongoId: "", ReservationEnd: time.Now().AddDate(0, -4, 0)})
	}
}
