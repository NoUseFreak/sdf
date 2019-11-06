package setup

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func GetConfigFile() string {
	return path.Join(createHomeDir(), viper.GetString("profile")+".yml")
}

func createHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeCfgDir := path.Join(usr.HomeDir, ".sdf")
	if _, err := os.Stat(homeCfgDir); os.IsNotExist(err) {
		os.Mkdir(homeCfgDir, 0777)
	}
	return homeCfgDir
}

func SaveConfig() {
	err := viper.WriteConfigAs(GetConfigFile())
	if err != nil {
		output.Print("%v\n", err)
	}
}
