package command

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/config"
)

var rootCmd = &cobra.Command{
	Use:   "sdf",
	Short: "Lazy developer toolbox",
	Long:  `Doing stuff the easy way`,
}

func init() {
	rand.Seed(time.Now().UnixNano())

	rootCmd.SetOutput(os.Stderr)
	rootCmd.PersistentFlags().StringP("verbosity", "v", logrus.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := initLogs(os.Stderr, cmd.Flags().Lookup("verbosity").Value.String()); err != nil {
			return err
		}
		return nil
	}

	cf := config.InitConfig()
	cf.ReadConfig(viper.GetString("profile"))
}

func initLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}
}
