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

func TestInitDB(t *testing.T) {
	t.Run("initDB success", func(t *testing.T) {
		caught := catchFatalExit(func() {
			InitDB()
		})
		if caught {
			t.Errorf("initDB() called fatalExit()")
		}
	})

	t.Run("initDB failure", func(t *testing.T) {
		util.SetupEnvironmentVariables("../properties/empty.properties")
		caught := catchFatalExit(func() {
			InitDB()
		})
		if !caught {
			t.Errorf("initDB() did not return an error")
		}
		util.SetupTestEnvironment()
	})
}

func TestOpenSqlConnection(t *testing.T) {
	t.Run("openSqlConnection success", func(t *testing.T) {
		caught := catchFatalExit(func() {
			openSqlConnection()
		})
		if caught {
			t.Errorf("openSqlConnection() called fatalExit()")
		}
	})

	t.Run("openSqlConnection failure", func(t *testing.T) {
		sqlOpenString = func() string { return "invalid_dsn" }
		caught := catchFatalExit(func() {
			openSqlConnection()
		})
		if !caught {
			t.Errorf("openSqlConnection() did not return an error")
		}
		sqlOpenString = generateSqlOpenString
	})
}

func TestPingDatabase(t *testing.T) {
	openSqlConnection()
	t.Run("pingDatabase success", func(t *testing.T) {
		caught := catchFatalExit(func() {
			pingDatabase()
		})
		if caught {
			t.Errorf("pingDatabase() called fatalExit()")
		}
	})

	t.Run("pingDatabase failure", func(t *testing.T) {
		util.SetupEnvironmentVariables("../properties/empty.properties")
		DB.Close()
		openSqlConnection()
		caught := catchFatalExit(func() {
			pingDatabase()
		})
		if !caught {
			t.Errorf("pingDatabase() did not return an error")
		}
		util.SetupTestEnvironment()
	})
}

func TestGenerateSqlOpenString(t *testing.T) {
	t.Run("generateSqlOpenString with default environment", func(t *testing.T) {
		util.SetupEnvironmentVariables("../properties/default.properties")
		expected := "tictactoe_app:password@tcp(127.0.0.1:3306)/tictactoe_database?parseTime=true"
		actual := generateSqlOpenString()
		if actual != expected {
			t.Errorf("generateSqlOpenString() = %s; want %s", actual, expected)
		}
	})

	t.Run("generateSqlOpenString with test environment", func(t *testing.T) {
		util.SetupEnvironmentVariables("../properties/test.properties")
		expected := "tictactoe_test_app:password@tcp(127.0.0.1:3307)/tictactoe_database_test?parseTime=true"
		actual := generateSqlOpenString()
		if actual != expected {
			t.Errorf("generateSqlOpenString() = %s; want %s", actual, expected)
		}
	})

	t.Run("generateSqlOpenString with empty environment", func(t *testing.T) {
		util.SetupEnvironmentVariables("../properties/empty.properties")
		expected := ":@tcp(:)/?parseTime=true"
		actual := generateSqlOpenString()
		if actual != expected {
			t.Errorf("generateSqlOpenString() = %s; want %s", actual, expected)
		}
		util.SetupTestEnvironment()
	})

}

func catchFatalExit(f func()) (caught bool) {
	fatalExit = func(...interface{}) {
		caught = true
	}
	f()
	return
}
