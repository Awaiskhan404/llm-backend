package vector

type VectorRepository interface {
	GetVectorById(id int) (*Vector, error)
	GetAllVectors() ([]Vector, error)
	CreateVector(name string, description string, connection_string string) (*int32, error)
	UpdateVector(vector *Vector) error
}