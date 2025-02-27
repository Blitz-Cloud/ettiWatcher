package cmd

import (
	"fmt"

	"github.com/blitz-cloud/semHelper/config"
)

func Increment(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Invalid number of args\nPlease enter what you want to increment{year or semester}\n")
	}
	conf := config.ConfigFile{}
	err := conf.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	switch args[0] {
	case "year":
		conf.UnivYear++
		err = conf.WriteConfigFile()
	case "semester":
		conf.Semester++
		err = conf.WriteConfigFile()
	}
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
