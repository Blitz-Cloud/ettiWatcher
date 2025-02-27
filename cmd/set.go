package cmd

import (
	"fmt"

	"github.com/blitz-cloud/semHelper/config"
)

func Set(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("\tInvalid number of parameters\n\tUsing this command you can set the folder location for new projects and the text editor of your choice\n\tSee semHelper help to see examples")
	}
	conf := config.ConfigFile{}
	err := conf.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	switch args[0] {
	case "location":
		conf.LabsLocation = args[1]
		err = conf.WriteConfigFile()
	case "editor":
		conf.Editor = args[1]
		err = conf.WriteConfigFile()
	}
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
