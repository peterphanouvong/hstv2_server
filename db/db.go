package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	HOST = "localhost"
	PORT = 5432
)
var DEFAULT_DATE = time.Time{}

var ErrNoMatch = fmt.Errorf("no matching record")
type Database struct {
	Conn *pgxpool.Pool
}

func Initialize(username, password, database, dbHost string) (Database, error) {
	db := Database{}
	// db_url := `postgresql://peterhst:tomtom123@localhost:5432/hst?sslmode=disable`
	// fmt.Println(db_url)

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, database, password)
	fmt.Println(dbUri)

	conn, err := pgxpool.Connect(context.Background(), dbUri)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return db, err
	}
	db.Conn = conn

	err = db.Conn.Ping(context.Background())
	if err != nil {
		return db, err
	}

	fmt.Println("Database connection established")
	return db, nil
}