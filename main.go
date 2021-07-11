package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/peterphanouvong/hst/db"
	"github.com/peterphanouvong/hst/handler"
)

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		log.Print(e)
	}
	
	dbUser, dbPassword, dbName, dbHost :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_DB_HOST")

	database, err := db.Initialize(dbUser, dbPassword, dbName, dbHost)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()
	 
	handler.HandleRequests(database)	
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}