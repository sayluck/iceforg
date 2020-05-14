package model

type Storage interface {
	Save() (string, error)
	DetailByKeyProperty() (interface{}, error)
	IsExistedByKeyProperty() (bool, error)
	// TODO Add
	// Update() int
	// Delete() int
	// Query() int
}
