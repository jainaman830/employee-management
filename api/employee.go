package api

import (
	"encoding/json"
	"net/http"
	"project/employee-management/helper"
	"project/employee-management/model"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request, s *model.Storage) {
	var user model.Employee
	//fetching payload
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid payload : "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(user.Name) == "" {
		http.Error(w, "Provide employee name", http.StatusBadRequest)
		return
	} else if strings.TrimSpace(user.Position) == "" {
		http.Error(w, "Provide employee position", http.StatusBadRequest)
		return
	} else if user.Salary <= 0 {
		http.Error(w, "Salary must be greater than 0", http.StatusBadRequest)
		return
	}
	// Store User information
	s.Lock()
	defer s.Unlock()
	// Get the maximum ID from the storage and increment it to assign a new ID
	maxID := helper.GetMaxID(s)
	user.ID = maxID + 1
	s.Employees[user.ID] = user
	//final response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User Registered")
	return
}
func GetByID(w http.ResponseWriter, r *http.Request, s *model.Storage) {
	urlParams := mux.Vars(r)
	id, ok := urlParams["id"]
	if !ok {
		http.Error(w, "Id can not be blank", http.StatusBadRequest)
		return

	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}
	employee, empOk := s.Employees[idInt]
	if !empOk {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	//final response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	return
}
func Update(w http.ResponseWriter, r *http.Request, s *model.Storage) {
	var user model.Employee
	//fetching payload
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid payload : "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(user.Name) == "" {
		http.Error(w, "Provide employee name", http.StatusBadRequest)
		return
	} else if strings.TrimSpace(user.Position) == "" {
		http.Error(w, "Provide employee position", http.StatusBadRequest)
		return
	} else if user.Salary <= 0 {
		http.Error(w, "Salary must be greater than 0", http.StatusBadRequest)
		return
	} else if user.ID < 0 {
		http.Error(w, "Provid valid Id", http.StatusBadRequest)
		return
	}
	_, empOk := s.Employees[user.ID]
	if !empOk {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	// Update User information
	s.Lock()
	defer s.Unlock()
	user.ID = 1
	s.Employees[user.ID] = user
	//final response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")
	return
}
func Delete(w http.ResponseWriter, r *http.Request, s *model.Storage) {
	param := struct {
		Id int
	}{}
	//fetching payload
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, "Invalid payload : "+err.Error(), http.StatusBadRequest)
		return
	}
	if param.Id < 0 {
		http.Error(w, "Provid valid Id", http.StatusBadRequest)
		return
	}
	_, empOk := s.Employees[param.Id]
	if !empOk {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	s.Lock()
	defer s.Unlock()

	delete(s.Employees, param.Id)
	//final response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Delete Success")
	return
}

func EmployeeList(w http.ResponseWriter, r *http.Request, s *model.Storage) {
	// Default page
	params := model.PaginationParams{
		Page:  1,
		Limit: 10,
	}

	pageStr := r.URL.Query().Get("page")

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		params.Page = page
	}
	start := (params.Page - 1) * params.Limit
	end := start + params.Limit

	var employeesList []model.Employee
	s.RLock()
	defer s.RUnlock()
	for _, emp := range s.Employees {
		employeesList = append(employeesList, emp)
	}
	if len(employeesList) == 0 {
		http.Error(w, "No employee found", http.StatusNotFound)
	}
	sort.Slice(employeesList, func(i, j int) bool {
		return employeesList[i].ID < employeesList[j].ID
	})
	if start >= len(employeesList) {
		employeesList = []model.Employee{}
	} else if end > len(employeesList) {
		employeesList = employeesList[start:]
	} else {
		employeesList = employeesList[start:end]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employeesList)
}
