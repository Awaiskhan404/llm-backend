package vector

type VectorRepository interface {
	GetVectorById(id int) (*Vector, error)
	GetAllVectors() ([]Vector, error)
	CreateVector(name string, description string, connection_string string) (*int32, error)
	UpdateVector(
		id int,
		name string,
		description string,
		connection_string string,
	) (*Vector, error)
}