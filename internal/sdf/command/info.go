package command

import (
	"context"
	"flag"
	"github.com/NoUseFreak/sdf/internal/pkg/output"

	"github.com/google/subcommands"
)

type InfoCmd struct {
}

func (*InfoCmd) Name() string     { return "info" }
func (*InfoCmd) Synopsis() string { return "Print args to stdout." }
func (*InfoCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *InfoCmd) SetFlags(f *flag.FlagSet) {
	// f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *InfoCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		output.Print("%s\n ", arg)
	}
	return subcommands.ExitSuccess
}
