package model

type Storage interface {
	Save() (int64, error)
	DetailByKeyProperty() (interface{}, error)
	IsExistedByKeyProperty() (bool, error)
	// TODO Add
	// Update() int
	// Delete() int
	// Query() int
}
