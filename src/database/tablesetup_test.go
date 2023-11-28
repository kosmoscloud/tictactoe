package database

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestSetupUserTable(t *testing.T) {

	t.Run("setupUserTable success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnResult(sqlmock.NewResult(1, 1))
		err := setupUserTable()

		if err != nil {
			t.Errorf("The function should have returned nil, but returned %v instead.", err)
		}
	})

	t.Run("setupUserTable error returned from database", func(t *testing.T) {
		dberr := errors.New("MySQL error: connection refused")
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnError(dberr)
		err := setupUserTable()

		if err == nil {
			t.Error("The function should have returned an error, but returned nil instead.")
		}
	})
}
