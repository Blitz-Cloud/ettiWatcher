/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
		fmt.Println("lab called")
	},
}

func init() {
	newCmd.AddCommand(labCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
