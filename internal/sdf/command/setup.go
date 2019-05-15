package command

import (
	"context"
	"flag"
	"github.com/NoUseFreak/sdf/internal/pkg/output"

	"github.com/NoUseFreak/sdf/internal/sdf/setup"
	"github.com/google/subcommands"
	"github.com/spf13/viper"
)

type SetupCmd struct {
}

func (*SetupCmd) Name() string     { return "setup" }
func (*SetupCmd) Synopsis() string { return "Setup a maintainable structure." }
func (*SetupCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *SetupCmd) SetFlags(f *flag.FlagSet) {
	// f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *SetupCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	homeCfgDir := setup.CreateHomeDir()
	setup.CreateConfig(homeCfgDir)

	output.Print("Completed setting up your %s profile", viper.GetString("profile"))

	return subcommands.ExitSuccess
}
