package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterphanouvong/hst/models"
)

func addPeopleHandler(r *mux.Router) {
	r.HandleFunc("/person", addPerson).Methods("POST")
	r.HandleFunc("/people", updatePerson).Methods("PUT")
	r.HandleFunc("/people", getAllPeople).Methods("GET")
	r.HandleFunc("/people/cognito/{cognitoId}", getPersonByCognitoId).Methods("GET")
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	people, err := dbInstance.GetAllPersons()
	if err != nil {
		fmt.Println("we got an error - get all people")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(people)
}


func getPersonByCognitoId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cognitoId := vars["cognitoId"]
	person, err := dbInstance.GetPersonByCognitoId(cognitoId)
	if err != nil {
		fmt.Println("we got an error - get person by cognito id")
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(person)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - add person")
		fmt.Println(err.Error())
	}

	var person models.Person
	json.Unmarshal(reqBody, &person)

	// fmt.Println(person)

	newPerson, err := dbInstance.AddPerson(&person);

	if err != nil {
		fmt.Println("we got an error - add person")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(newPerson)

}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("we got an error - update person")
		fmt.Println(err.Error())
	}

	var person models.Person
	json.Unmarshal(reqBody, &person)

	updatedPerson, err := dbInstance.UpdatePerson(&person)

	if err != nil {
		fmt.Println("we got an error - update person")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(updatedPerson)
}