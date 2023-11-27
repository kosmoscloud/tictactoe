package database

import (
	"log"
	"os"
	"testing"
	"tictactoe-service/util"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	err := util.SetupTestEnvironment()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Test environment setup successful")
	}
	os.Exit(m.Run())
}

func TestDbConnection(t *testing.T) {
	t.Run("InitDB success", func(t *testing.T) {
		err := InitDB()
		if err != nil {
			t.Errorf("InitDB() returned an error: %s", err)
		}
	})

	t.Run("CloseDB success", func(t *testing.T) {
		err := CloseDB()
		if err != nil {
			t.Errorf("CloseDB() returned an error: %s", err)
		}
	})
}
