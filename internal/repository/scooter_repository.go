package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"main/internal/model"
)

type ScooterRepository interface {
	FindScootersByStatusAndLocation(status bool, latitude float64, longitude float64) []model.Scooter
	SetScooterStatus(scooterUuid uuid.UUID, status bool)
}

type mysqlScooterRepository struct {
	db *sql.DB
}

func (r *mysqlScooterRepository) FindScootersByStatusAndLocation(status bool, latitude float64, longitude float64) []model.Scooter {
	return []model.Scooter{}
}

func (r *mysqlScooterRepository) SetScooterStatus(scooterUuid uuid.UUID, status bool) {
	return
}

func NewScooterRepository(db *sql.DB) ScooterRepository {
	return &mysqlScooterRepository{db: db}
}
