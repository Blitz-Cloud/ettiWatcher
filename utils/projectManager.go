package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Blitz-Cloud/ettiHelper/types"
	"github.com/spf13/viper"
)

func GetPrettyDate(refTime time.Time) string {

	return fmt.Sprintf("%02d_%s_%d", refTime.Day(), refTime.Month().String()[:3], refTime.Year())
}

func GetRFC3339Time(refTime time.Time) string {
	return refTime.UTC().Local().Format(time.RFC3339)
}

func GetLabsLocation() string {
	labsLocation := viper.GetString("labs_location")
	if labsLocation == "DEFAULT" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Ceva gresit este cu sistemul tau de operare.\nVariabila care ar trebui sa indice unde se gaseste folterul utilizatorului curent nu exista.\n\tEsti pe cont propriu de aici inainte")
		}
		labsLocation = fmt.Sprintf("%s%s", userHomeDir, "/facultate")
	}
	return labsLocation
}

func GetSubjects() []string {

	subjectsDirs, err := os.ReadDir(GetLabsLocation())
	if err != nil {
		log.Fatal(err)
	}
	subjects := make([]string, 0)
	for _, subjectDir := range subjectsDirs {
		subjects = append(subjects, subjectDir.Name())
	}

	return subjects
}

func GetProjectsMetadata(subject string) []FrontmatterMetaData {
	projectsMetadataList := make([]FrontmatterMetaData, 0)
	DirCrawler(GetLabsLocation()+"/"+subject, func(path string, file os.DirEntry) {
		if file.Name() == "README.md" {
			content, err := os.ReadFile(path + "/README.md")
			if err != nil {
				log.Fatal(err)
			}
			metadata, _ := ParseMdString(string(content))
			projectsMetadataList = append(projectsMetadataList, metadata)
		}
	}, func(paht string, dir os.DirEntry) {})
	return projectsMetadataList
}

func GetProjectData(path string) types.Lab {
	readme, err := os.ReadFile(filepath.Join(path, "README.md"))
	if err != nil {
		log.Fatal(err)
	}
	programFile := []byte("")
	if _, err := os.Stat(filepath.Join(path, "main.cpp")); err == nil {
		programFile, err = os.ReadFile(filepath.Join(path, "main.cpp"))
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(filepath.Join(path, "main.c")); err == nil {
		programFile, err = os.ReadFile(filepath.Join(path, "main.c"))
		if err != nil {
			log.Fatal(err)
		}
	}

	metadata, content := ParseMdString(string(readme))
	if metadata.Subject != "blog" {
		content = fmt.Sprintf("%s\n\n Codul sursa:\n```cpp\n%s\n```", content, string(programFile))
	}
	data := types.Lab{
		Title:              metadata.Title,
		Description:        metadata.Description,
		Date:               metadata.Date,
		Tags:               "",
		Subject:            metadata.Subject,
		UniYearAndSemester: uint(metadata.UniYearAndSemester),
		Content:            content,
	}
	return data
}

func CreateDirectory(projectName, subject, projectType string) string {
	// ar trebui sa fie capabil sa verifice daca exista deja proiectul daca da sa iasa o erroare
	uniYear := viper.GetInt("uni_year")
	uniSemester := viper.GetInt("semester")
	var path string
	projectDirectoryName := fmt.Sprintf("%s-%d-%s", projectName, uniYear*10+uniSemester, GetPrettyDate(time.Now()))

	if projectType == "lab" {
		path = fmt.Sprintf("%s/%s/%s", GetLabsLocation()+"/labs", subject, projectDirectoryName)
	} else if projectType == "blog" {
		path = fmt.Sprintf("%s/%s/%s", GetLabsLocation(), "blog", projectDirectoryName)
	}

	err := os.MkdirAll(path, 0766)
	if err != nil {
		log.Fatal(err)
	}
	return path

}
