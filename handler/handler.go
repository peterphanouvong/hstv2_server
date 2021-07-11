package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterphanouvong/hst/db"
	"github.com/rs/cors"
)

var dbInstance db.Database

type Query struct {
	Query string `json:"query"`
}


func HandleRequests(db db.Database) {
	dbInstance = db
	fmt.Println(dbInstance)
	fmt.Println("handle requests")
	addr := ":10000"
	r := mux.NewRouter()
	r.HandleFunc("/query", runQuery).Methods("POST")
	addPeopleHandler(r)
	addUsersHandler(r)
	addAddressesHandler(r)
	addFranchisesHandler(r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet,//http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",//or you can your header key values which you are using in your application
		},
	})

	handler := c.Handler(r)
	
	log.Fatal(http.ListenAndServe(addr, handler))
}

func runQuery (w http.ResponseWriter, r *http.Request) {
	var query Query
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &query)
	fmt.Println(string(reqBody))
	json.NewEncoder(w).Encode(query)
}