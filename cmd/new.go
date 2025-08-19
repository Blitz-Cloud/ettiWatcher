/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Acesta este comanda care este folostia pentru a creea noi laboratoare sau fisiere de documentatie",
	Long: `Aceasta comanda accepta 2 parametri, lab sau blog. 
	Fiecare dintre acestia sunt necesari pentru ca programul sa fie executat.
	Exemple de utilizare:
		1. Pentru a creea un laborator:
				semHelper new lab [c/cpp] [name] 

		2. Pentru a crea un fisier markdown:
				semHelper new blog [name]`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
