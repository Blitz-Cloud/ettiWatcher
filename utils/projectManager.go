package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

func GenerateDateStandard() string {
	return fmt.Sprintf("%d_%s_%d", time.Now().Day(), time.Now().Month().String()[:3], time.Now().Year())
}

func CreateDirectory(projectName, subject string) string {
	// ar trebui sa fie capabil sa verifice daca exista deja proiectul daca da sa iasa o erroare
	uniYear := viper.GetInt("uni_year")
	uniSemester := viper.GetInt("semester")
	labsLocation := viper.GetString("labs_location")
	if labsLocation == "DEFAULT" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Ceva gresit este cu sistemul tau de operare.\nVariabila care ar trebui sa indice unde se gaseste folterul utilizatorului curent nu exista.\n\tEsti pe cont propriu de aici inainte")
		}
		labsLocation = fmt.Sprintf("%s%s", userHomeDir, "/facultate/labs")
	}

	projectDirectoryName := fmt.Sprintf("%s-%d-%s", projectName, uniYear*10+uniSemester, GenerateDateStandard())
	path := fmt.Sprintf("%s/%s/%s", labsLocation, subject, projectDirectoryName)
	err := os.MkdirAll(path, 0766)
	if err != nil {
		log.Fatal(err)
	}
	return path

}
