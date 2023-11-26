package util

import (
	"os"
	"strings"

	"github.com/magiconair/properties"
)

func ReadProperties(path string) (map[string]string, error) {
	p := properties.MustLoadFile(path, properties.UTF8)
	return p.Map(), nil
}

func SetupEnvironmentVariables(path string, done chan bool) error {
	p := properties.MustLoadFile(path, properties.UTF8)
	for k, v := range p.Map() {
		k = strings.ToUpper(k)
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	done <- true
	return nil
}
