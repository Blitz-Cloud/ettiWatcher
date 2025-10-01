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
	"strings"

	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
		gitURL := args[0]
		lab := args[1]
		//https://github.com/Blitz-Cloud/ettiContent{.git}/[subject]/[title]-[uniYearAndSemester]-[date]
		if gitURL[len(gitURL)-1] != '/' {
			gitURL = gitURL + "/"
		}
		gitUrlSplit := strings.Split(gitURL, "/")
		spew.Dump(gitUrlSplit)
		repoName := gitUrlSplit[len(gitUrlSplit)-3] + "_" + gitUrlSplit[len(gitUrlSplit)-2]
		if _, err := os.Stat(filepath.Join(utils.GetRootDirectory(), repoName)); os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(utils.GetRootDirectory(), repoName), 0766)
		}

		utils.CloneRepo(filepath.Join(utils.GetRootDirectory(), repoName), gitURL)

		editor := viper.GetString("preferred_editor")
		fmt.Println(filepath.Join(utils.GetRootDirectory(), repoName, lab))
		execEditor := exec.Command(editor, filepath.Join(utils.GetRootDirectory(), repoName, lab))
		err := execEditor.Start()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
