package model

import "sync"

// Employee
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

// Storage
type Storage struct {
	sync.RWMutex
	Employees map[int]Employee
}

// PaginationParams
type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
