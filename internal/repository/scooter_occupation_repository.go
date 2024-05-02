package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/model"
)

type ScooterOccupationRepository interface {
	Create(scooterUuid uuid.UUID, userUuid uuid.UUID) (created bool)
	DeleteByScooterUuidAndUserUuid(scooterUuid uuid.UUID, userUuid uuid.UUID) (recordDeleted bool)
	GetByScooterIdAndUserId(scooterUuid uuid.UUID, userUuid uuid.UUID) (model.ScooterOccupation, bool)
}

type mysqlScooterOccupationRepository struct {
	db *sql.DB
}

func (r mysqlScooterOccupationRepository) Create(scooterUuid uuid.UUID, userUuid uuid.UUID) (created bool) {
	scooterUuidAsBinary, _ := scooterUuid.MarshalBinary()
	userUuidAsBinary, _ := userUuid.MarshalBinary()

	query := "INSERT INTO scooters_occupations (scooter_id, user_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, scooterUuidAsBinary, userUuidAsBinary)

	if err != nil {
		return false
	}

	return true
}

func (r mysqlScooterOccupationRepository) DeleteByScooterUuidAndUserUuid(
	scooterUuid uuid.UUID,
	userUuid uuid.UUID,
) (recordDeleted bool) {

	scooterUuidAsBinary, _ := scooterUuid.MarshalBinary()
	userUuidAsBinary, _ := userUuid.MarshalBinary()

	query := "DELETE FROM scooters_occupations WHERE scooter_id = ? AND user_id = ?"
	result, err := r.db.Exec(query, scooterUuidAsBinary, userUuidAsBinary)

	if err != nil {
		panic(err)
	}

	deletedRowsCount, _ := result.RowsAffected()

	return deletedRowsCount == 1
}

func (r mysqlScooterOccupationRepository) GetByScooterIdAndUserId(
	scooterUuid uuid.UUID,
	userUuid uuid.UUID,
) (model.ScooterOccupation, bool) {
	scooterUuidAsBinary, _ := scooterUuid.MarshalBinary()
	userUuidAsBinary, _ := userUuid.MarshalBinary()

	query := "SELECT * FROM scooters_occupations WHERE scooter_id = ? AND user_id = ?"

	var scooterOccupation model.ScooterOccupation
	err := r.db.QueryRow(query, scooterUuidAsBinary, userUuidAsBinary).Scan(
		&scooterOccupation.Id,
		&scooterOccupation.ScooterId,
		&scooterOccupation.UserId,
	)

	if err != nil {
		return scooterOccupation, false
	}

	return scooterOccupation, true
}

func NewScooterOccupationRepository(db *sql.DB) ScooterOccupationRepository {
	return &mysqlScooterOccupationRepository{db: db}
}
