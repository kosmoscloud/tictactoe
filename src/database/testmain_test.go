package database

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"tictactoe-service/util"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	mockDB *sql.DB
	mock   sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	var err error

	mockDB, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	DB = mockDB

	err = util.SetupTestEnvironment()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Test environment setup successful")
	}

	os.Exit(m.Run())
}
