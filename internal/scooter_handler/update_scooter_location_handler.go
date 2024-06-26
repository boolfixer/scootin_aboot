package scooter_handler

import (
	"github.com/google/uuid"
	"main/internal/dto"
	"main/internal/http_error"
	"main/internal/repository"
)

type UpdateScooterLocationHandler struct {
	scooterRepository           repository.ScooterRepository
	scooterOccupationRepository repository.ScooterOccupationRepository
}

func (h UpdateScooterLocationHandler) Handle(
	scooterId uuid.UUID,
	userId uuid.UUID,
	scooterLocationUpdate dto.ScooterLocationUpdate,
) error {
	_, exists := h.scooterOccupationRepository.GetByScooterIdAndUserId(scooterId, userId)

	if !exists {
		return http_error.ConflictError{Message: "Scooter is not occupied by current user."}
	}

	scooter, exists := h.scooterRepository.GetByScooterId(scooterId)

	if !exists {
		return http_error.NotFoundError{ModelName: "Scooter"}
	}

	if scooterLocationUpdate.Time.Before(scooter.LocationUpdatedAt) {
		return http_error.ConflictError{Message: "Scooter location is outdated."}
	}

	err := h.scooterRepository.UpdateScooterCoordinatesByScooterId(
		scooterId,
		scooterLocationUpdate.Latitude,
		scooterLocationUpdate.Longitude,
		scooterLocationUpdate.Time,
	)

	if err != nil {
		return http_error.ConflictError{Message: "Failed to update scooter coordinates."}
	}

	return nil
}

func NewUpdateScooterLocation(
	scooterOccupationRepository repository.ScooterOccupationRepository,
	scooterRepository repository.ScooterRepository,
) *UpdateScooterLocationHandler {
	return &UpdateScooterLocationHandler{
		scooterRepository:           scooterRepository,
		scooterOccupationRepository: scooterOccupationRepository,
	}
}
