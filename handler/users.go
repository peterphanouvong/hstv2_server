package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterphanouvong/hst/models"
)

func addUsersHandler(r *mux.Router) {
	r.HandleFunc("/users/type/{user_type_id}", getUsersByType).Methods("GET")
	r.HandleFunc("/users/cognito/{cognitoId}", getUserByCognitoId).Methods("GET")
	r.HandleFunc("/users", addUser).Methods("POST")
	r.HandleFunc("/users", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", getUserById).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		fmt.Println("we got an error - get all users")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(users)
}

func getUsersByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_type_id := vars["user_type_id"]

	users, err := dbInstance.GetUsersByType(user_type_id);

	if err != nil {
		fmt.Println("we got an error - get users by type")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(users)
}

func getUserByCognitoId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cognitoId := vars["cognitoId"]
	user, err := dbInstance.GetUserByCognitoId(cognitoId)
	if err != nil {
		fmt.Println("we got an error - get user by cognito id")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := dbInstance.GetUserById(id)
	if err != nil {
		fmt.Println("we got an error - get user by id")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBody,err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - add user")
		fmt.Println(err.Error())
	}

	var user models.User
	json.Unmarshal(reqBody, &user)

	newUser, err := dbInstance.AddUser(user)

	if err != nil {
		fmt.Println("we got an error - add user")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(newUser)

}

func updateUser(w http.ResponseWriter, r * http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("we got an error - update user")
		fmt.Println(err.Error())
	}

	var user models.User
	json.Unmarshal(reqBody, &user)

	updatedUser, err := dbInstance.UpdateUser(user)
	if err != nil {
		fmt.Println("we got an error - update user")
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(updatedUser)
}