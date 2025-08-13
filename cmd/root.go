/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "semHelper",
	Short: "Un tool CLI creat pentru a automatiza crearea mediului de dezvoltare pentru laboratoarele de la facultate si nu numai",
	Long: `SemHelper este un cli conceput pentru a usura crearea mediului de dezvoltare necesar pentru laboratoarele de la facultate
	Acest CLI are mult mai multe capabilitati decat cele care sunt menite pentru studenti, deoarece acesta este o componenta complementara site-ului web https://ettih.blitzcloud.me`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.AddConfigPath("$HOME/.ettiWatcher/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			userHomeDir, err := os.UserHomeDir()
			cobra.CheckErr(err)
			err = os.MkdirAll(userHomeDir+"/.ettiWatcher", os.FileMode(0777))
			cobra.CheckErr(err)
			err = viper.WriteConfigAs(userHomeDir + "/.ettiWatcher/config.yaml")
			cobra.CheckErr(err)
		}

	}
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("uni-year", 0)
	viper.SetDefault("semester", 0)
	viper.SetDefault("preferred_editor", "")
	viper.SetDefault("labs_location", "DEFAULT")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.semHelper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
