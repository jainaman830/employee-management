package main

import (
	"fmt"
	"net/http"
	"project/employee-management/api"
	"project/employee-management/helper"

	"github.com/gorilla/mux"
)

func main() {
	//
	store := helper.NewStorage()
	//router creation
	server := mux.NewRouter().StrictSlash(true)
	server.HandleFunc("/employee/create", func(w http.ResponseWriter, r *http.Request) {
		api.Register(w, r, store)
	}).Methods("POST")
	server.HandleFunc("/employee/getbyid/{id}", func(w http.ResponseWriter, r *http.Request) {
		api.GetByID(w, r, store)
	})
	server.HandleFunc("/employee/update", func(w http.ResponseWriter, r *http.Request) {
		api.Update(w, r, store)
	}).Methods("PUT")
	server.HandleFunc("/employee/delete", func(w http.ResponseWriter, r *http.Request) {
		api.Delete(w, r, store)
	}).Methods("DELETE")
	server.HandleFunc("/employee/list", func(w http.ResponseWriter, r *http.Request) {
		api.EmployeeList(w, r, store)
	})
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", server)
}
