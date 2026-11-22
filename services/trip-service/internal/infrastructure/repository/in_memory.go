package repository

import (
	"context"
	"sync"

	"github.com/canonflow/ride-sharing/services/trip-service/internal/domain"
)

type InMemoryRepository struct {
	trips     map[string]*domain.TripModel
	rideFates map[string]*domain.RideFareModel
	rw        sync.RWMutex
}

func NewInMemoryRepository() domain.ITripRepository {
	return &InMemoryRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFates: make(map[string]*domain.RideFareModel),
	}
}

func (r *InMemoryRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.rw.Lock()
	r.trips[trip.ID.Hex()] = trip
	r.rw.Unlock()

	return trip, nil
}
