package setup

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
)

func CreateHomeDir() string {
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

func CreateConfig(homeCfgDir string) {
	output.Print("%v\n", viper.GetString("profile"))
	err := viper.WriteConfigAs(path.Join(homeCfgDir, viper.GetString("profile")+".yml"))
	if err != nil {
		output.Print("%v\n", err)
	}
}
