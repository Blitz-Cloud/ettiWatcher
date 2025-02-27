package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const ConfigTemplate = `labsLocation: %s
lastProcessed: %s
year: %d
semester: %d
editor: %s
`

var HomeDir = os.Getenv("HOME")
var ConfigDir = path.Join(HomeDir, ".semHelper")
var ConfigFileLocation = path.Join(ConfigDir, ".config")

type ConfigFile struct {
	LabsLocation  string `yaml:"labsLocation"`
	LastProcessed string `yaml:"lastProcessed"`
	UnivYear      int    `yaml:"year"`
	Semester      int    `yaml:"semester"`
	Editor        string `yaml:"editor"`
}

func (config ConfigFile) WriteConfigFile() error {
	content := fmt.Sprintf(ConfigTemplate, config.LabsLocation, config.LastProcessed, config.UnivYear, config.Semester, config.Editor)
	err := os.MkdirAll(ConfigDir, 0766)
	if err != nil {
		return fmt.Errorf("I couldn't create the folder\n%s", err)
	}
	err = os.WriteFile(ConfigFileLocation, []byte(content), 0766)
	if err != nil {
		return fmt.Errorf("I couldn't write the config file\n%s", err)
	}
	// asta este pentru test
	return nil
}

func (config *ConfigFile) ReadConfigFile() error {
	content, err := os.ReadFile(ConfigFileLocation)
	if err != nil {
		return fmt.Errorf("Error while reading the file")
	}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return fmt.Errorf("Error parsing the file")
	}
	return nil
}
