package model

type Storage interface {
	Save() (string, error)
	DetailByKeyProperty() (interface{}, error)
	IsExistedByKeyProperty() (bool, error)
	List() (interface{}, error)
	// TODO Add
	// Update() int
	// Delete() int
	// Query() int
}
