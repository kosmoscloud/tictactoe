package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"tictactoe-service/server/dto"

	_ "github.com/go-sql-driver/mysql"
)

func setupUserTable(DB *sql.DB) error {
	query := generateCreateTableQuery(dto.User{})
	_, err := DB.Exec(query)

	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Default().Print("Users table setup successfully!")
		return nil
	}
}

func generateCreateTableQuery(data interface{}) string {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("sql")
		keyTag := field.Tag.Get("key")

		if tag != "" {
			fields = append(fields, fmt.Sprintf("%s %s %s", tag, getFieldType(value), keyTag))
		}
	}

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS users (%s)", strings.Join(fields, ", "))

	return query
}

func getFieldType(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Int:
		return "INT"
	case reflect.String:
		return "VARCHAR(255)"
	case reflect.Struct:
		switch value.Type().String() {
		case "datetime.DateTime":
			return "DATETIME"
		case "datetime.Date":
			return "DATE"
		case "time.Time":
			return "DATETIME"
		default:
			return ""
		}
	default:
		return ""
	}
}
