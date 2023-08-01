package repository

type Repository interface {
	All() (interface{}, error)
	// Get(interface{}) (interface{}, error)
	// Create(interface{}) (interface{}, error)
	// Update(interface{}) (interface{}, error)
	// Delete(interface{}) (interface{}, error)
}
