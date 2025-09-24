package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Blitz-Cloud/ettiHelper/types"
	"github.com/blitz-cloud/ettiWatcher/templates"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

func GetProjectsMetadata(subject string) []FrontmatterMetaDataType {
	projectsMetadataList := make([]FrontmatterMetaDataType, 0)
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

func CreateProject(metadata ProjectMetadataType) {
	projectDirectoryName := fmt.Sprintf("%s-%d-%s", strings.ReplaceAll(metadata.Title, " ", "_"), metadata.UniYearAndSemester, GetPrettyDate(time.Now()))
	projectPath := filepath.Join(GetLabsLocation(), metadata.Subject, projectDirectoryName)
	// fmt.Println("Proiectul va fi creat la aceasta locatie: " + projectPath)
	if _, err := os.Stat(projectPath); err == nil {
		fmt.Println("Acest proiect deja exista, daca alegi sa continui atunci toate datele despre acest proiect vor fi pierdute. Continui ? (y/n)")
		var userInput string
		fmt.Scanf("%s", &userInput)
		switch userInput {
		case "y":
			err := os.RemoveAll(projectPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Aceasta operatie este reversibila folosind GIT")
		case "n":
			return
		default:
			fmt.Println("Acest raspuns este invalid si va fi tratat ca si cum ar fi NU")
		}
	}

	err := os.MkdirAll(projectPath, 0766)
	if err != nil {
		log.Fatal(err)
	}

	if metadata.DirOnly {
		fmt.Printf("Pentru a accesa proiectul:\ncd %s", projectPath)
		return
	} else {

		cmakeFile := ""
		mainFile := ""
		extension := ""
		readmeFile := fmt.Sprintf(templates.MDTemplate, metadata.Title, GetRFC3339Time(time.Now()), metadata.Subject, "", metadata.UniYearAndSemester)
		projectName := strings.ReplaceAll(metadata.Title, " ", "_")
		switch metadata.Lang {
		case "c":
			cmakeFile = fmt.Sprintf(templates.CMakeForC, projectName, projectName)
			mainFile = templates.CTemplate
			extension = ".c"
		case "cpp":
			fallthrough
		case "c++":
			cmakeFile = fmt.Sprintf(templates.CMakeForCpp, projectName, projectName)
			mainFile = templates.CppTemplate
			extension = ".cpp"
			// default:
			// 	log.Fatalf("%s nu este un limbaj supportat.\n Doar c si cpp sunt variante valide", metadata.Lang)
		}

		if metadata.Subject != "blog" {
			err = os.WriteFile(filepath.Join(projectPath, "CMakeLists.txt"), []byte(cmakeFile), 0766)
			if err != nil {
				log.Fatalf("%s", err)
			}
			err = os.WriteFile(filepath.Join(projectPath, "main"+extension), []byte(mainFile), 0766)
			if err != nil {
				log.Fatalf("%s", err)
			}
			err = os.WriteFile(filepath.Join(projectPath, ".gitignore"), []byte("/build"), 0766)
			if err != nil {
				log.Fatalf("%s", err)
			}
		}

		err = os.WriteFile(filepath.Join(projectPath, "README.md"), []byte(readmeFile), 0766)
		if err != nil {
			log.Fatalf("%s", err)
		}

	}
	if metadata.GitEnable && viper.GetBool("git_enabled") {
		commitNewFilesToGitRepo()
	} else {
		addPathToRootGitIgnore(projectPath)
	}

	AddToSyncQueue(projectPath)
	if metadata.OpenEditor {
		editor := viper.GetString("preferred_editor")
		execEditor := exec.Command(editor, projectPath)
		err = execEditor.Start()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Pentru a accesa proiectul:\ncd %s", projectPath)
	}

}

func createGitRepoIfNotExists(path string) {
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {
		_, err := git.PlainInit(path, false)
		if err != nil {
			log.Fatalf("failed to initialize git repository: %v", err)
		}
	}

}

func commitNewFilesToGitRepo() {
	repoPath := GetLabsLocation()
	createGitRepoIfNotExists(repoPath)
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("failed to open repository: %v", err)
	}

	// Get the working tree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("failed to get worktree: %v", err)
	}
	_, err = worktree.Add(".")
	if err != nil {
		log.Fatalf("failed to add all files: %v", err)
	}
	// fmt.Println("All changes added to staging area.")
	// fmt.Println("Performing 'git commit -m \"Added new files\"'...")
	_, err = worktree.Commit("Added new lab", &git.CommitOptions{
		Author: &object.Signature{
			Name:  viper.GetString("git_name"),
			Email: viper.GetString("git_email"),
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("failed to commit changes: %v", err)
	}

	// fmt.Printf("Successfully committed with hash: %s\n", commitHash)
}

func addPathToRootGitIgnore(path string) {}
