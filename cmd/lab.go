/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/blitz-cloud/ettiWatcher/types"
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
		// editor := viper.GetString("preferred_editor")
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
		date := time.Now()
		newProject := types.ProjectMetadataType{
			FrontmatterMetaDataType: types.FrontmatterMetaDataType{
				Title:              projectName,
				UniYearAndSemester: uniYearAndSemester,
				Subject:            subject,
				Date:               &date,
			},
			Lang:       projectLang,
			DirOnly:    createDirOnly,
			OpenEditor: true,
			GitEnable:  true,
		}
		utils.CreateProject(newProject)
	},
}

func init() {
	newCmd.AddCommand(labCmd)

	labCmd.Flags().BoolP("createDirOnly", "d", false, "Flag optional care indica faptul ca doar folderul ar trebui creat")
	labCmd.Flags().String("subject", "", "Flag optional care indica faptul ca ar trebui ca proiectul sa fie creat in folderul dat nu in cel prestabilit")
}
