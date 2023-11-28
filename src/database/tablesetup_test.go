package database

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestSetupUserTable(t *testing.T) {

	dberr := errors.New("MySQL error: connection refused")

	t.Run("setupUserTable success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnResult(sqlmock.NewResult(1, 1))
		result := setupUserTable()

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		if result != nil {
			t.Errorf("The function should have returned nil, but returned %s instead.", result)
		}
	})

	t.Run("setupUserTable error returned from database", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnError(dberr)
		result := setupUserTable()

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		if result == nil {
			t.Errorf("The function should have returned an error, but returned %s instead.", result)
		}
	})
}
