package main

import (
	"context"
	"log"
	"time"

	"github.com/canonflow/ride-sharing/services/trip-service/internal/domain"
	"github.com/canonflow/ride-sharing/services/trip-service/internal/infrastructure/repository"
	"github.com/canonflow/ride-sharing/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()
	inMemoryRepo := repository.NewInMemoryRepository()
	svc := service.NewTripService(inMemoryRepo)

	fare := &domain.RideFareModel{
		UserID: "42",
	}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	// Keep the program runnin for now
	for {
		time.Sleep(time.Second)
	}
}
