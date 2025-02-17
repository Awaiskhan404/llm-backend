package vector

type VectorRepository interface {
	GetVectorById(id int) (*Vector, error)
	GetAllVectors() ([]Vector, error)
}