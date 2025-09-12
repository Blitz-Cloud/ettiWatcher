/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/blitz-cloud/ettiWatcher/templates"
	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// labCmd represents the lab command
var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "Comanda folosita pentru a creea mediul de dezvoltare pentru un laborator",
	Long: `Acesta comanda este folosita pentru a creea mediul de dezvoltare pentru un laborator. 
	La acest moment este supportat doar c si cpp, insa odata cu trecerea timpului vor fi introduse si alte limbaje de programare.
	De asemenea pentru a putea folosi mediul de dezvoltare este necesar ca CMAKE si un compilator de c/cpp sa fie instalate si locatiile executabilelor sa fie prezente in PATH-ul calculatorului folosit.
	Limitari:
		1. numele trebuie sa fie legat de preferat sa fie utiliazt camelCase

	
	Exemple de utilizare:
		semHelper new lab [c/cpp] [name] `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Aceasta comanda accepta doar 2 parametri.")
			cmd.Help()
			return
		}
		editor := viper.GetString("preferred_editor")
		uniYearAndSemester := viper.GetInt("uni_year")*10 + viper.GetInt("semester")
		projectLang := args[0]
		projectName := args[1]
		// validez argumentele posibil sa fie necesara

		// incarcarea flags

		createDirOnly, err := cmd.Flags().GetBool("createDirOnly")
		if err != nil {
			log.Fatal(err)
		}

		subject, err := cmd.Flags().GetString("subject")
		if err != nil {
			log.Fatal(err)
		}

		if subject == "" {
			subject = viper.GetString("subject")
		}

		// posibila solutie pentru a rezolva si blog
		// mai jos legat de createOnlyDir
		projectLocation := utils.CreateDirectory(projectName, subject, "lab")

		if createDirOnly {
			// sa rulez functia care creeaza doar folderul
			fmt.Printf("Pentru a accesa proiectul:\ncd %s", projectLocation)
			return
		}
		// fisierele necesare pt proiect c/cpp cmake readme.md

		cmakeFile := ""
		mainFile := ""
		extension := ""
		readmeFile := fmt.Sprintf(templates.MDTemplate, projectName, utils.GetRFC3339Time(time.Now()), subject, "", uniYearAndSemester)

		switch projectLang {
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
		default:
			log.Fatalf("%s nu este un limbaj supportat.\n Doar c si cpp sunt variante valide", projectLang)
		}

		err = os.WriteFile(filepath.Join(projectLocation, "CMakeLists.txt"), []byte(cmakeFile), 0766)
		if err != nil {
			log.Fatalf("%s", err)
		}
		err = os.WriteFile(filepath.Join(projectLocation, "main"+extension), []byte(mainFile), 0766)
		if err != nil {
			log.Fatalf("%s", err)
		}
		err = os.WriteFile(filepath.Join(projectLocation, "README.md"), []byte(readmeFile), 0766)
		if err != nil {
			log.Fatalf("%s", err)
		}
		err = os.Chdir(projectLocation)
		if err != nil {
			log.Fatal(err)
		}

		utils.AddToSyncQueue(projectLocation)

		execEditor := exec.Command(editor, projectLocation)
		err = execEditor.Start()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(labCmd)

	labCmd.Flags().BoolP("createDirOnly", "d", false, "Flag optional care indica faptul ca doar folderul ar trebui creat")
	labCmd.Flags().String("subject", "", "Flag optional care indica faptul ca ar trebui ca proiectul sa fie creat in folderul dat nu in cel prestabilit")
}
