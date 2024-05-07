package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/employee-management/api"
	"project/employee-management/helper"
	"project/employee-management/model"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	storage := helper.NewStorage()
	employee := model.Employee{Name: "Aman", Position: "SDE", Salary: 50000}

	payload, _ := json.Marshal(employee)
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(payload))
	rec := httptest.NewRecorder()

	api.Register(rec, req, storage)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "User Registered", response)
	assert.Len(t, storage.Employees, 1)
}

func TestGetByID(t *testing.T) {
	storage := helper.NewStorage()
	employee := model.Employee{Name: "Aman", Position: "SDE", Salary: 50000}
	storage.Employees[1] = employee

	req, _ := http.NewRequest("GET", "/employee/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rec := httptest.NewRecorder()

	api.GetByID(rec, req, storage)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response model.Employee
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "Aman", response.Name)
}

func TestUpdate(t *testing.T) {
	storage := helper.NewStorage()
	employee := model.Employee{Name: "Aman", Position: "SDE", Salary: 50000}
	storage.Employees[1] = employee

	updatedEmployee := model.Employee{ID: 1, Name: "Aman Jain", Position: "SDE 3", Salary: 60000}
	payload, _ := json.Marshal(updatedEmployee)
	req, _ := http.NewRequest("PUT", "/update", bytes.NewReader(payload))
	rec := httptest.NewRecorder()

	api.Update(rec, req, storage)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "Success", response)

	assert.Equal(t, "Aman Jain", storage.Employees[1].Name)
}

func TestDelete(t *testing.T) {
	storage := helper.NewStorage()
	employee := model.Employee{Name: "Aman Jain", Position: "SDE", Salary: 50000}
	storage.Employees[1] = employee

	payload, _ := json.Marshal(struct{ Id int }{Id: 1})
	req, _ := http.NewRequest("DELETE", "/delete", bytes.NewReader(payload))
	rec := httptest.NewRecorder()

	api.Delete(rec, req, storage)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response string
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(t, "Delete Success", response)

	_, ok := storage.Employees[1]
	assert.False(t, ok)
}
