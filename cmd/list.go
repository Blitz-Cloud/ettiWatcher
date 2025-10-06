/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/blitz-cloud/ettiWatcher/types"
	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Comanda folosita pentru a afisa toate proiectele chiar si pentru a le deschide",
	Long: `
	Exemplu de utilizare: ettiWatcher list --open --all -m`,
	Run: func(cmd *cobra.Command, args []string) {
		curentSubject := viper.GetString("subject")
		subjectF, _ := cmd.Flags().GetString("subject")
		subjectsF, _ := cmd.Flags().GetBool("subjects")
		remoteF, _ := cmd.Flags().GetString("remote")
		remotesF, _ := cmd.Flags().GetBool("remotes")
		showAllSubjectsF, _ := cmd.Flags().GetBool("all")
		openTui, _ := cmd.Flags().GetBool("open")
		remotes := utils.GetRemotes()
		projectsMetadata := make([]types.FrontmatterMetaDataType, 0)
		if len(remoteF) != 0 && showAllSubjectsF {
			log.Fatal("Este imposibil sa folosesti --all si --remote in acelasi timp.\n")
		}
		if remotesF && subjectsF {
			log.Fatal("Este imposibil sa foleseti  --remotes si --subjects in acelasi timp ")
		}
		if subjectsF {
			for _, remote := range remotes {
				fmt.Printf("Materii in folederul: %s\n", remote)
				subjects := utils.GetSubjects(filepath.Join(utils.GetRootDirectory(), remote))
				for i, subject := range subjects {
					fmt.Printf("\t%d. %s\n", i+1, subject)
				}
				fmt.Println()
			}
			os.Exit(0)
		}
		if remotesF {
			fmt.Println("In aceste foldere sunt salvare laburile(fie ca este vorba de un repo de pe github sau ce este local):")
			for i, remote := range remotes {
				fmt.Printf("\t%d. %s\n", i+1, remote)
			}
			os.Exit(0)
		}
		if showAllSubjectsF {
			for _, remote := range remotes {
				if len(subjectF) != 0 {
					projectsMetadata = append(projectsMetadata, utils.GetProjectsMetadata(filepath.Join(utils.GetRootDirectory(), remote, subjectF))...)
				} else {
					subjects := utils.GetSubjects(filepath.Join(utils.GetRootDirectory(), remote))
					for _, subject := range subjects {
						projectsMetadata = append(projectsMetadata, utils.GetProjectsMetadata(filepath.Join(utils.GetRootDirectory(), remote, subject))...)
					}
				}
			}
		} else {
			subject := curentSubject
			if len(subjectF) != 0 {
				subject = subjectF
			}
			remote := "local"
			if len(remoteF) != 0 {
				remote = remoteF
			}
			projectsMetadata = append(projectsMetadata, utils.GetProjectsMetadata(filepath.Join(utils.GetRootDirectory(), remote, subject))...)
		}

		if openTui {
			items := make([]list.Item, 0)
			for _, project := range projectsMetadata {
				items = append(items, utils.Item{
					Metadata: project,
				})
			}
			utils.RUNTUI(items, "Toate proiectele")
		} else {
			fmt.Println("Lista tuturor proiectelor")
			for i, project := range projectsMetadata {
				fmt.Printf("%d.", i+1)
				fmt.Printf("\tNume: %s\n", project.Title)
				fmt.Printf("\tCreat la: %s\n", project.Date)
				fmt.Printf("\tMateria: %s\n", project.Subject)
				fmt.Printf("\tDescrierea proiectului: %s\n\n", project.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// ettiWatcher list --subject --subjects --open --all
	listCmd.Flags().String("subject", "", "Flag optional pentru care va permite afisarea doar laboratoarelor pentru materia introdusa")
	listCmd.Flags().Bool("subjects", false, "Flag optional care daca este seta va afisa o lista cu toate materiile pentru care au fost create laboratoare")
	listCmd.Flags().Bool("open", false, "Flag optional care daca este setat permite intrarea intr un mod interactiv de unde pot fi cautate laboratore. Este incompatibil cu afisarea materiilor.\nExemplu de utilizare:\n ettiWatcher list --open ")
	listCmd.Flags().BoolP("all", "a", false, "Flag optional care permite afisarea tutror laboratoarelor fara sa tina conta de materia care a fost setata. Acest flag ofera o versabilitate foarte buna atunci cand este folosit impreuna cu --open ca mai jos:\n ettiWatcher list --all --open")
	listCmd.Flags().String("remote", "", "Flag folosit pentru a seta daca este folosit un proiect remote sau cel local")
	listCmd.Flags().Bool("remotes", false, "Flag folosit pentru a afisa toate remoteurile")

}
