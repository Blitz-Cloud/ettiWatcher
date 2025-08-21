/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Prin intermediul acestei comenzi utilizatorul poate sa seteze parametrii precum anul, semestrul, editorul preferat, precum si locatia unde sunt create proiectele de laborator",
	Long: `Help menu
	Disclaimere:
		1. La acest moment nu este posibila utilizarea variabilelor din bash precum $HOME
		2. Nu este implementat inca o incrementare automata a anului si semestrului de facultate 
		3. Pentru a putea deschide editorul de text trebuie ca acesta sa se afle in PATH-ul sistemului sau al utilizatorului curent
	Exemple de utilizare:
	ettiWatcher set --year 1 --semester 1 --editor code --path /home/blitz-cloud/facultate/pclp2/labs --subject pclp1`,
	Run: func(cmd *cobra.Command, args []string) {
		uniYear, _ := cmd.Flags().GetInt("year")
		semester, _ := cmd.Flags().GetInt("semester")
		subject, _ := cmd.Flags().GetString("subject")
		preferredEditor, _ := cmd.Flags().GetString("editor")
		labsLocation, _ := cmd.Flags().GetString("path")

		var isAtLeastAFlagSet bool = false

		cmd.Flags().Visit(func(f *pflag.Flag) {
			isAtLeastAFlagSet = true
		})
		if !isAtLeastAFlagSet {
			fmt.Println("Erroare: Pentru a folosi aceasta comanda trebuie sa fie setat cel putin un flag")
			cmd.Help()
			return
		}

		if uniYear != 0 {
			viper.Set("uni_year", uniYear)
		}

		if semester != 0 {
			viper.Set("semester", semester)
		}

		if subject != "" {
			viper.Set("subject", subject)
		}

		if preferredEditor != "" {
			viper.Set("preferred_editor", preferredEditor)
		}

		if labsLocation != "DEFAULT" {
			viper.Set("labs_location", labsLocation)
		}
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().IntP("year", "y", 0, "Seteaza anul de studii")
	setCmd.Flags().IntP("semester", "s", 0, "Seteaza semesterul de studii")
	setCmd.Flags().StringP("editor", "e", "", "Seteaza editorul preferat de utilizator")
	setCmd.Flags().String("subject", "", "Seteaza materia pentru care vor fi create proiectele")
	setCmd.Flags().StringP("path", "p", "DEFAULT", "Seteaza locatia unde utilizatorul doreste sa salveze proiectele de laborator, daca nu este definita o locatie pentru laboratore atunci $HOME/facultate/labs va fi folosita pe post de locatie")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
