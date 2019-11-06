package short

import (
	"fmt"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/setup"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type short struct {
	Name string
	Cmd  string
}

func RegisterShortCommands(cmd *cobra.Command) {
	for _, short := range List() {
		cmd.AddCommand(&cobra.Command{
			Use:    short.Name,
			Hidden: true,

			Run: func(cmd *cobra.Command, args []string) {
				RunShort(short.Name)
			},
		})
	}
}

func List() []short {
	var list []short

	for name := range viper.GetStringMap("short") {
		list = append(list, *Get(name))
	}

	return list
}

func Get(name string) *short {
	key := fmt.Sprintf("short.%s", name)

	cfg := viper.GetStringMapString(key)
	if cmd, ok := cfg["cmd"]; ok {
		return &short{
			Name: name,
			Cmd:  cmd,
		}
	}

	return nil
}

func Add(name, cmd string) {
	key := fmt.Sprintf("short.%s.cmd", name)
	viper.Set(key, cmd)
	setup.SaveConfig()
}

func RunShort(name string) {
	if short := Get(name); short != nil {
		output.Exec(short.Cmd)
	} else {
		output.Errorln("Could not find short %s", name)
	}
}

func Remove(name string) error {
	viper.Set("short."+name, nil)
	setup.SaveConfig()

	return fmt.Errorf("Could not remove short, do it in %s", setup.GetConfigFile())
}
