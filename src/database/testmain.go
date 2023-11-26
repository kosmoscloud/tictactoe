package database

import (
	"log"
	"os"
	"testing"
	"tictactoe-service/util"
)

func TestMain(m *testing.M) {

	log.Println("Done setting up environment variables")
	doneSettingUpEnv := make(chan bool)
	go util.SetupEnvironmentVariables("properties/test.properties", doneSettingUpEnv)
	<-doneSettingUpEnv

	os.Exit(m.Run())
}
