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
		editor := viper.GetString("preferred_editor")
		uniYearAndSemester := viper.GetInt("uni_year")*10 + viper.GetInt("semester")
		projectName := args[0]
		// validez argumentele posibil sa fie necesara

		// incarcarea flags

		subject, err := cmd.Flags().GetString("subject")
		if err != nil {
			log.Fatal(err)
		}

		if subject == "" {
			subject = viper.GetString("subject")
		}

		// posibila solutie pentru a rezolva si blog
		// mai jos legat de createOnlyDir
		projectLocation := utils.CreateDirectory(projectName, subject, "blog")

		// fisierele necesare pt proiect c/cpp cmake readme.md

		readmeFile := fmt.Sprintf(templates.MDTemplate, projectName, utils.GetRFC3339Time(time.Now()), "blog", "", uniYearAndSemester)

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
	newCmd.AddCommand(blogCmd)

	blogCmd.Flags().BoolP("createDirOnly", "d", false, "Flag optional care indica faptul ca doar folderul ar trebui creat")
	blogCmd.Flags().String("subject", "", "Flag optional care indica faptul ca ar trebui ca proiectul sa fie creat in folderul dat nu in cel prestabilit")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
