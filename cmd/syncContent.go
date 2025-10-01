/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"path/filepath"

	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/spf13/cobra"
)

// syncContentCmd represents the syncContent command
var syncContentCmd = &cobra.Command{
	Use:   "syncGit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.UpdateAndPushGitRepo(filepath.Join(utils.GetRootDirectory(), "local"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	adminCmd.AddCommand(syncContentCmd)

}
