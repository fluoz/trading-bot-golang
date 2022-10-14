package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFromYAML(path string, target interface{}) error {
	yf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}
