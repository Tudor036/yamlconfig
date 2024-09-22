package yamlconfig

import (
	"errors"
	"reflect"
	"testing"
)

const (
	hostname = "localhost"
	port     = 8080
	path     = "./data.sqlite"
)

type ConfigType struct {
	Server struct {
		Hostname string `yaml:"hostname"`
		Port     int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`
}

func TestLoadYaml(t *testing.T) {
	yaml, err := LoadYaml[ConfigType]("./test.yaml")

	if err != nil {
		t.Error(err)
		return
	}

	expected := new(ConfigType)
	expected.Server.Hostname = hostname
	expected.Server.Port = port
	expected.Database.Path = path

	if !reflect.DeepEqual(yaml, expected) {
		t.Error(errors.New("Values don't match."))
		return
	}
}
