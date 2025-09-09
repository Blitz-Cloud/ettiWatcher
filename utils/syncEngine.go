package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

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

func UpdateSyncTimeStamp() error {

	client := &http.Client{}
	req, err := http.NewRequest("POST", GetSyncServerURL()+"/last-sync", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("admin_token")))
	_, err = client.Do(req)
	return err
}
