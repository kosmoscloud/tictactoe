package database

import (
	"fmt"
	"log"
	"regexp"
	"testing"
	"tictactoe-service/server/dto"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetUser(t *testing.T) {
	t.Run("getUser successful", func(t *testing.T) {
		currentTime := time.Now()
		rows := sqlmock.NewRows([]string{"id", "username", "created"}).AddRow(1, "test", currentTime)
		// todo: figure out why can't we use hardcoded id value (such as id=1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnRows(rows)
		user, err := GetUser(1)
		if err != nil {
			t.Errorf("getUser failed: %v", err)
		}
		if !validateUserFieldsValues(user, 1, "test", currentTime) {
			t.Errorf("getUser failed: %v", err)
		}
	})

	t.Run("getUser failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnError(fmt.Errorf("some error"))
		user, err := GetUser(1)
		if err == nil {
			t.Error("getUser failed: expected error")
		}
		if user != nil {
			t.Errorf("getUser failed: expected nil user, got %v", user)
		}
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("createUser successful", func(t *testing.T) {
		currentTime := time.Now()
		rows := sqlmock.NewRows([]string{"id", "username", "created"}).AddRow(1, "test", currentTime)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, created) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
		// todo: figure out why can't we use hardcoded id value (such as id=1)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnRows(rows)
		user, err := CreateUser("test")
		if err != nil {
			t.Errorf("createUser failed: %v", err)
		}
		if !validateUserFieldsValues(user, 1, "test", currentTime) {
			t.Errorf("createUser failed: %v", err)
		}
	})

	t.Run("createUser db exec failed", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, created) VALUES (?, ?)")).WillReturnError(fmt.Errorf("some error"))
		user, err := CreateUser("test")
		if err == nil {
			t.Error("createUser failed: expected error")
		}
		if user != nil {
			t.Errorf("createUser failed: expected nil user, got %v", user)
		}
	})

	t.Run("createUser parse last inserted id failed", func(t *testing.T) {
		// todo: how to test parsing last inserted id failed?
	})

	t.Run("createUser get user failed", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, created) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnError(fmt.Errorf("some error"))
		user, err := CreateUser("test")
		if err == nil {
			t.Error("createUser failed: expected error")
		}
		if user != nil {
			t.Errorf("createUser failed: expected nil user, got %v", user)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("updateUser successful", func(t *testing.T) {
		currentTime := time.Now()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET username=? WHERE id=?")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "created"}).AddRow(1, "test", currentTime))
		user, err := UpdateUser(1, "test")
		if err != nil {
			t.Errorf("test failed: %v", err)
		}
		if !validateUserFieldsValues(user, 1, "test", currentTime) {
			t.Errorf("test failed: %v", err)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("deleteUser successful", func(t *testing.T) {
		currentTime := time.Now()
		rows := sqlmock.NewRows([]string{"id", "username", "created"}).AddRow(1, "test", currentTime)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id=?")).WillReturnRows(rows)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id=?")).WillReturnResult(sqlmock.NewResult(1, 1))
		user, err := DeleteUser(1)
		if err != nil {
			t.Errorf("test failed: %v", err)
		}
		if !validateUserFieldsValues(user, 1, "test", currentTime) {
			t.Errorf("test failed: %v", err)
		}
	})
}

func validateUserFieldsValues(user *dto.User, id int, username string, created time.Time) bool {
	if user.UserId != fmt.Sprint(id) {
		log.Printf("user.UserId: %v does not equal id: %v", user.UserId, id)
		return false
	}
	if user.Username != username {
		log.Printf("user.Username: %v does not equal username: %v", user.Username, username)
		return false
	}
	if user.CreatedDate != created {
		log.Printf("user.CreatedDate: %v does not equal created: %v", user.CreatedDate, created)
		return false
	}
	return true
}
