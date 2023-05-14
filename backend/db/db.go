package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQL struct {
	Db *sqlx.DB
}

// Connect to db postgres
func (s *SQL) Connect() error {
	dbDriver := "postgres"
	dbUser := "quanbh"
	dbPass := "19111993"
	dbHost := "localhost"
	dbName := "ecommerce_platform"

	db, err := sqlx.Open(dbDriver, fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbName))

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")

	s.Db = db

	return nil
}

func (s *SQL) Close() error {
	err := s.Db.Close()
	if err != nil {
		return err
	}

	fmt.Println("Closed database connection")
	return nil
}
