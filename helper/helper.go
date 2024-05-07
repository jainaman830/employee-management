package helper

import "project/employee-management/model"

// NewStorage creates a new instance of Storage.
func NewStorage() *model.Storage {
	return &model.Storage{
		Employees: make(map[int]model.Employee),
	}
}

// GetMaxID returns the maximum ID from the storage.
func GetMaxID(s *model.Storage) int {
	maxID := 0
	for id := range s.Employees {
		if id > maxID {
			maxID = id
		}
	}
	return maxID
}
