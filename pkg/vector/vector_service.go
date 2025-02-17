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

func (service *VectorService) GetAllVectors() ([]Vector, error) {
	service.logger.Info("Retrieving all vectors")

	rows, err := service.db.Query(
		context.Background(),
		"SELECT id, name, description, connection_string, created_at::TEXT FROM vector.databases;",
	)
	if err != nil {
		service.logger.Fatal("Error while executing query. Err:", err)
		return nil, err
	}
	defer rows.Close()

	var result []Vector

	for rows.Next() {
		var vector Vector
		var createdAtStr string
		err := rows.Scan(&vector.ID, &vector.Name, &vector.Description, &vector.ConnectionString, &createdAtStr)
		if err != nil {
			service.logger.Fatal("Error while scanning row. Err:", err)
			return nil, err
		}

		vector.CreatedAt, err = time.Parse("2006-01-02", createdAtStr)
		if err != nil {
			service.logger.Fatal("Error while parsing created_at. Err:", err)
			return nil, err
		}

		result = append(result, vector)
	}

	return result, nil
}

func (service *VectorService) CreateVector(name string, description string, connection_string string) (*int32, error) {
	service.logger.Info("Creating vector:", name)

	var id int32

	err := service.db.QueryRow(
		context.Background(),
		"INSERT INTO vector.databases (name, description, connection_string, created_at) VALUES ($1, $2, $3, $4) RETURNING id;",
		name, description, connection_string, time.Now(),
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