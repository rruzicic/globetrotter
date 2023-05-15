package services

import (
	"github.com/rruzicic/globetrotter/bnb/feedback-service/dtos"
	//grpcclient "github.com/rruzicic/globetrotter/bnb/feedback-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHostReview(hostReviewDTO dtos.CreateHostReviewDTO) (*models.HostReview, error) {
	userId, err := primitive.ObjectIDFromHex(hostReviewDTO.UserId)
	if err != nil {
		return nil, err
	}

	hostId, err := primitive.ObjectIDFromHex(hostReviewDTO.HostId)
	if err != nil {
		return nil, err
	}

	hostReview := models.HostReview{
		UserId: &userId,
		HostId: &hostId,
		Rating: hostReviewDTO.Rating,
	}

	//TODO check if user had a previous reservation with this host.

	return repos.CreateHostReview(hostReview)
}

func GetHostReviewById(id string) (*models.HostReview, error) {
	return repos.GetHostReviewById(id)
}

func GetHostReviewsByUserId(id string) ([]models.HostReview, error) {
	return repos.GetHostReviewsByUserId(id)
}

func GetHostReviewsByHostId(id string) ([]models.HostReview, error) {
	return repos.GetHostReviewsByHostId(id)
}

func DeleteHostReview(id string) error {
	return repos.DeleteHostReview(id)
}

func UpdateHostReview(hostReviewDTO dtos.CreateHostReviewDTO) error {
	userId, err := primitive.ObjectIDFromHex(hostReviewDTO.UserId)
	if err != nil {
		return err
	}

	hostId, err := primitive.ObjectIDFromHex(hostReviewDTO.HostId)
	if err != nil {
		return err
	}

	hostReview := models.HostReview{
		UserId: &userId,
		HostId: &hostId,
		Rating: hostReviewDTO.Rating,
	}

	return repos.UpdateHostReview(hostReview)
}

//===================================================================================
//===================================================================================

func CreateAccommodationReview(accommodationReviewDTO dtos.CreateAccommodationReviewDTO) (*models.AccommodationReview, error) {
	userId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.UserId)
	if err != nil {
		return nil, err
	}

	accommodationId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.AccommodationId)
	if err != nil {
		return nil, err
	}

	accommodationReview := models.AccommodationReview{
		UserId:          &userId,
		AccommodationId: &accommodationId,
		Rating:          accommodationReviewDTO.Rating,
	}

	//TODO check if user had a previous reservation with this accommodation.

	return repos.CreateAccommodationReview(accommodationReview)
}

func GetAccommodationtReviewById(id string) (*models.AccommodationReview, error) {
	return repos.GetAccommodationReviewById(id)
}

func GetAccommodationReviewsByUserId(id string) ([]models.AccommodationReview, error) {
	return repos.GetAccommodationReviewsByUserId(id)
}

func GetAccommodationReviewsByAccommodationId(id string) ([]models.AccommodationReview, error) {
	return repos.GetAccommodationReviewsByAccommodationId(id)
}

func DeleteAccommodationReview(id string) error {
	return repos.DeleteHostReview(id)
}

func UpdateAccommodationReview(accommodationReviewDTO dtos.CreateAccommodationReviewDTO) error {
	userId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.UserId)
	if err != nil {
		return err
	}

	accommodationId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.AccommodationId)
	if err != nil {
		return err
	}

	accommodationReview := models.AccommodationReview{
		UserId:          &userId,
		AccommodationId: &accommodationId,
		Rating:          accommodationReviewDTO.Rating,
	}

	return repos.UpdateAccommodationReview(accommodationReview)
}
