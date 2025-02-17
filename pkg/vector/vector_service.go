package vector

import (
	"context"
	"errors"
	"llm-backend/pkg/lib"
	"time"
)

type VectorService struct {
	logger lib.Logger
	db     *lib.Database
}

func GetVectorService(logger lib.Logger, db *lib.Database) VectorRepository {
	return &VectorService{
		logger: logger,
		db:     db,
	}
}

func (service VectorService) GetVectorById(id int) (*Vector, error) {
	service.logger.Info("Retrieving vector with id", id)
	return nil, errors.New("Not implemented")
}

func (service VectorService) GetAllVectors() ([]Vector, error) {
	return nil, errors.New("Not implemented")
}

func (service *VectorService) CreateVector(name string, description string, connection_string string) (*int32, error) {
	service.logger.Info("Creating vector:", name)

	var id int32

	var created_at = time.Now()

	err := service.db.QueryRow(
		context.Background(),
		"INSERT INTO vector.databases (name, description, connection_string, created_at) VALUES ($1, $2, $3, $4) RETURNING id;",
		name, description, connection_string, created_at,
	).Scan(&id)

	if err != nil {
		service.logger.Fatal("Error while executing query. Err:", err)
		return nil, err
	}
	return &id, nil
}

func (service VectorService) UpdateVector(vector *Vector) error {
	return errors.New("Not implemented")
}