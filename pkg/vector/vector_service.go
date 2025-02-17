package vector

import (
	"errors"
	"llm-backend/pkg/lib"
)

type VectorService struct {
	logger lib.Logger
	db     *lib.Database
}

func GetVectorService(logger lib.Logger, db *lib.Database) VectorRepository {
	return VectorService{
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