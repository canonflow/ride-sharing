package service

import (
	"context"

	"github.com/canonflow/ride-sharing/services/trip-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripService struct {
	tripRepo domain.ITripRepository
}

func NewTripService(tripRepo domain.ITripRepository) domain.ITripSerivice {
	return &TripService{
		tripRepo: tripRepo,
	}
}

func (s *TripService) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {
	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}

	return s.tripRepo.CreateTrip(ctx, t)
}
