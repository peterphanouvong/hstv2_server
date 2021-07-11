package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterphanouvong/hst/models"
)

func addFranchisesHandler(r *mux.Router) {
	// r.HandleFunc("/franchises/{id}", getFranchise).Methods("GET")
	r.HandleFunc("/franchises", getFranchiseList).Methods("GET")
	r.HandleFunc("/franchises", addFranchise).Methods("POST")
	// r.HandleFunc("/address/cognito/{cognitoId}", getPersonByCognitoId).Methods("GET")
}

func getFranchiseList(w http.ResponseWriter, r *http.Request) {
	franchises, err := dbInstance.GetFranchiseList()
	if err != nil {
		fmt.Println("we got an error - Get Franchise List")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(franchises)
}

// func getFranchise(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	user, err := dbInstance.GetAddress(id)
// 	if err != nil {
// 		fmt.Println("we got an error - get address")
// 		fmt.Println(err.Error())
// 	}

// 	w.Header().Set("Content-type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(user)
// }

func addFranchise(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Add address")

	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - add address")
		fmt.Println(err.Error())
	}

	var franchise models.Franchise
	json.Unmarshal(reqBody, &franchise)

	// fmt.Println(address)

	addr, err := dbInstance.AddFranchise(franchise)

	if err != nil {
		fmt.Println("we got an error - add franchise")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(addr)
}

// func updateFranchise(w http.ResponseWriter, r *http.Request) {
// 	reqBody, err := ioutil.ReadAll(r.Body)

// 	if err != nil {
// 		fmt.Println("we got an error - update address")
// 		fmt.Println(err.Error())
// 	}

// 	var address models.Address
// 	json.Unmarshal(reqBody, &address)

// 	updatedAddress, err := dbInstance.UpdateAddress(&address)

// 	if err != nil {
// 		fmt.Println("we got an error - update address")
// 		fmt.Println(err.Error())
// 	}
// 	w.Header().Set("Content-type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(updatedAddress)

// }
