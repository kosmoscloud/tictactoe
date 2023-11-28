package util

import (
	"os"
	"strings"

	"github.com/magiconair/properties"
)

func ReadProperties(path string) (map[string]string, error) {
	// todo: figure out whether MustLoadFile is an optimal choice (it panics on error)
	p := properties.MustLoadFile(path, properties.UTF8)
	return p.Map(), nil
}

func SetupEnvironmentVariables(path string) error {
	// todo: figure out whether MustLoadFile is an optimal choice (it panics on error)
	p := properties.MustLoadFile(path, properties.UTF8)
	for k, v := range p.Map() {
		k = strings.ToUpper(k)
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetupTestEnvironment() error {
	p := properties.MustLoadFile("../properties/test.properties", properties.UTF8)
	for k, v := range p.Map() {
		k = strings.ToUpper(k)
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
