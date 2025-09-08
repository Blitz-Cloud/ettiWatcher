package utils

import "github.com/spf13/viper"

func AddToSyncQueue(path string) {
	currentUnsynced := viper.GetStringSlice("unsynced")
	currentUnsynced = append(currentUnsynced, path)
	viper.Set("unsynced", currentUnsynced)
	viper.WriteConfig()
}

func GetSyncServerURL() string {
	if viper.GetString("env") == "prod" {
		return "https://ettih.blitzcloud.me/api/admin"
	} else {
		return "http://localhost:3000/api/admin"
	}
}
