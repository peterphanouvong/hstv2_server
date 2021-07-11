package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterphanouvong/hst/models"
)

func addAddressesHandler(r *mux.Router) {
	r.HandleFunc("/address/{id}", getAddress).Methods("GET")
	r.HandleFunc("/address", addAddress).Methods("POST")
	r.HandleFunc("/address", updateAddress).Methods("PUT")
	// r.HandleFunc("/address/cognito/{cognitoId}", getPersonByCognitoId).Methods("GET")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST")
}

func getAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := dbInstance.GetAddress(id)
	if err != nil {
		fmt.Println("we got an error - get address")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func addAddress(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Add address")

	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - add address")
		fmt.Println(err.Error())
	}

	var address models.Address
	json.Unmarshal(reqBody, &address)

	// fmt.Println(address)

	addr, err := dbInstance.AddAddress(address)

	if err != nil {
		fmt.Println("we got an error - add address")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(addr)
}

func updateAddress(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - update address")
		fmt.Println(err.Error())
	}

	var address models.Address
	json.Unmarshal(reqBody, &address)

	updatedAddress, err := dbInstance.UpdateAddress(&address)

	if err != nil {
		fmt.Println("we got an error - update address")
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(updatedAddress)

}
// func getAllAddresses(w http.ResponseWriter, r *http.Request) {
// 	people, err := dbInstance.GetAllPersons()
// 	if err != nil {
// 		fmt.Println("we got an error")
// 		fmt.Println(err.Error())
// 	}

// 	w.Header().Set("Content-type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(people)
// }

// func getAddressById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	cognitoId := vars["cognitoId"]
// 	person, err := dbInstance.GetPersonByCognitoId(cognitoId)
// 	if err != nil {
// 		fmt.Println("we got an error")
// 		fmt.Println(err.Error())
// 	}
// 	w.Header().Set("Content-type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(person)
// }