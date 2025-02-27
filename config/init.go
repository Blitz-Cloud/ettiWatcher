package config

import "path"

var conf = ConfigFile{}

func InitConfig() {
	err := conf.ReadConfigFile()
	if err != nil {
		labsLocation := path.Join(HomeDir, "facultate/labs")

		conf = ConfigFile{labsLocation, "1-Jan-1970", 1, 1, "code"}
		conf.WriteConfigFile()
	}
}
