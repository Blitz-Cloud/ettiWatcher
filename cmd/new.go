package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/blitz-cloud/semHelper/config"
	"github.com/blitz-cloud/semHelper/templates"
	"github.com/davecgh/go-spew/spew"
)

func New(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("\tInvalid number of args\n\tUSAGE: semHelper new lab name or semHelper new blog name\n")
	}
	conf := config.ConfigFile{}
	err := conf.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	switch args[0] {
	case "lab":
		spew.Dump(args)
		if len(args) < 3 {
			return fmt.Errorf("\tUSAGE: semHelper new lab c/cpp/c++ name")
		}
		dirName := fmt.Sprintf("%s-%d_%s_%d", strings.ReplaceAll(args[2], " ", "_"), time.Now().Day(), time.Now().Month().String()[:3], time.Now().Year())
		labLocation := path.Join(conf.LabsLocation, dirName)
		err := os.MkdirAll(labLocation, 0766)
		if err != nil {
			return fmt.Errorf("Dont be stupid this is probably not a valid path\n%s", err)
		}
		cmakeFile := ""
		mainFile := ""
		extension := ""
		if args[1] == "c" {
			cmakeFile = fmt.Sprintf(templates.CMakeForC, args[2], args[2])
			mainFile = templates.CTemplate
			extension = ".c"
		} else if args[1] == "cpp" || args[1] == "c++" {
			cmakeFile = fmt.Sprintf(templates.CMakeForCpp, args[2], args[2])
			mainFile = templates.CppTemplate
			extension = ".cpp"
		}

		err = os.WriteFile(filepath.Join(labLocation, "CMakeLists.txt"), []byte(cmakeFile), 0766)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		err = os.WriteFile(filepath.Join(labLocation, "main"+extension), []byte(mainFile), 0766)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		fmt.Println(conf.Editor)
		cmd := exec.Command(conf.Editor, labLocation)
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		return nil

	case "blog":
	}

	return nil
}
