/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
		subjectsF, _ := cmd.Flags().GetBool("subjects")
		subjectF, _ := cmd.Flags().GetString("subject")
		showAllSubjectsF, _ := cmd.Flags().GetBool("all")
		openTui, _ := cmd.Flags().GetBool("open")

		if subjectsF {
			fmt.Println("Materii:")
			subjects := utils.GetSubjects()
			for i, subject := range subjects {
				fmt.Printf("\t%d. %s\n", i+1, subject)
			}
			os.Exit(0)
		}

		if subjectF != "" {
			subjects := utils.GetSubjects()
			subjectFound := false
			for _, subject := range subjects {
				if subject == subjectF {
					subjectFound = true
					break
				}
			}
			if !subjectFound {
				fmt.Println("Aceasta nu este o materie valida sau nu a fost creata sau definita")
				os.Exit(1)
			}
			projectsMetaData := utils.GetProjectsMetadata(subjectF)
			if openTui {
				items := make([]list.Item, 0)
				for _, project := range projectsMetaData {
					items = append(items, utils.Item{
						Metadata: project,
					})
				}
				utils.RUNTUI(items,
					fmt.Sprintf("TUI mode for %s", subjectF))
			} else {
				fmt.Println("Lista tuturor proiectelor:")
				for i, project := range projectsMetaData {
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tNume: %s\n", project.Title)
					fmt.Printf("\tCreat la: %s\n", project.Date)
					fmt.Printf("\tMateria: %s\n", project.Subject)
					fmt.Printf("\tDescrierea proiectului: %s\n\n", project.Description)
				}
			}
			os.Exit(0)
		}

		if showAllSubjectsF {
			subjects := utils.GetSubjects()
			projectsMetaData := make([]utils.FrontmatterMetaData, 0)
			for _, subject := range subjects {
				projectsMetaData = append(projectsMetaData, utils.GetProjectsMetadata(subject)...)
			}

			if openTui {

				items := make([]list.Item, 0)
				for _, project := range projectsMetaData {
					items = append(items, utils.Item{
						Metadata: project,
					})
				}
				utils.RUNTUI(items, "Toate proiectele")

			} else {

				fmt.Println("Lista tuturor proiectelor")
				for i, project := range projectsMetaData {
					fmt.Printf("%d.", i+1)
					fmt.Printf("\tNume: %s\n", project.Title)
					fmt.Printf("\tCreat la: %s\n", project.Date)
					fmt.Printf("\tMateria: %s\n", project.Subject)
					fmt.Printf("\tDescrierea proiectului: %s\n\n", project.Description)
				}
			}
			os.Exit(0)
		}

		projectsMetaData := utils.GetProjectsMetadata(curentSubject)
		fmt.Println("Lista tuturor proiectelor")
		for i, project := range projectsMetaData {
			fmt.Printf("%d.", i+1)
			fmt.Printf("\tNume: %s\n", project.Title)
			fmt.Printf("\tCreat la: %s\n", project.Date)
			fmt.Printf("\tMateria: %s\n", project.Subject)
			fmt.Printf("\tDescrierea proiectului: %s\n\n", project.Description)
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

	// listCmd.Flags().BoolVarP(&test)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
