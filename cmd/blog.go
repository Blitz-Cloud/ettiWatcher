/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// blogCmd represents the blog command
var blogCmd = &cobra.Command{
	Use:   "blog",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("Aceasta comanda accepta doar un parametru.")
			cmd.Help()
			return
		}
		uniYearAndSemester := viper.GetInt("uni_year")*10 + viper.GetInt("semester")
		projectName := args[0]

		date := time.Now()
		newProject := utils.ProjectMetadataType{
			FrontmatterMetaDataType: utils.FrontmatterMetaDataType{
				Title:              projectName,
				UniYearAndSemester: uniYearAndSemester,
				Subject:            "blog",
				Date:               &date,
			},
			Lang:       "",
			DirOnly:    false,
			OpenEditor: true,
			GitEnable:  true,
		}
		utils.CreateProject(newProject)
	},
}

func init() {
	newCmd.AddCommand(blogCmd)

	// blogCmd.Flags().BoolP("createDirOnly", "d", false, "Flag optional care indica faptul ca doar folderul ar trebui creat")
	// blogCmd.Flags().String("subject", "", "Flag optional care indica faptul ca ar trebui ca proiectul sa fie creat in folderul dat nu in cel prestabilit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
