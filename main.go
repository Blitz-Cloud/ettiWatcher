package main

import (
	"github.com/blitz-cloud/semHelper/cmd"
	"github.com/blitz-cloud/semHelper/config"
)

func main() {
	config.InitConfig()
	cmd.RootCmd()
}
