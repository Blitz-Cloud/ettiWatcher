/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/blitz-cloud/ettiWatcher/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// syncContentCmd represents the syncContent command
var syncContentCmd = &cobra.Command{
	Use:   "syncContent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		queueToSync := viper.GetStringSlice("unsynced")
		spew.Dump(queueToSync)
		fmt.Printf("Sync started\nData to sync: %d", len(queueToSync))
		for _, path := range queueToSync {
			data := utils.GetProjectData(path)
			jsonBody, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
			}
			jsonBody = []byte(jsonBody)

			var contentType string
			if strings.Contains(path, "lab") {
				contentType = "lab"
			} else if strings.Contains(path, "blog") {
				contentType = "blog"
			}

			_, err = http.Post(utils.GetSyncServerURL()+"/"+contentType, "application/json", bytes.NewReader(jsonBody))
			if err != nil {
				log.Println(err)
			}
			// spew.Dump(resp)
			// sa modific datele de pe server si sa umblu sa syncQueue
		}
	},
}

func init() {
	adminCmd.AddCommand(syncContentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncContentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncContentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
