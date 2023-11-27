package database

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_setupUserTable_success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	t.Run("setupUserTable success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnResult(sqlmock.NewResult(1, 1))
		result := setupUserTable(db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		if result != nil {
			t.Errorf("The function should have returned nil, but returned %s instead.", result)
		}
	})
}

func Test_setupUserTable_errorReturnedFromDatabase(t *testing.T) {

	db, mock, mockerr := sqlmock.New()
	if mockerr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockerr)
	}

	dberr := errors.New("MySQL error: connection refused")

	defer db.Close()

	t.Run("setupUserTable error returned from database", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), created TIMESTAMP)")).WillReturnError(dberr)
		result := setupUserTable(db)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		if result == nil {
			t.Errorf("The function should have returned an error, but returned %s instead.", result)
		}
	})
}
