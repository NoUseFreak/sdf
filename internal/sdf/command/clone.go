package command

import (
	"context"
	"flag"
	"github.com/NoUseFreak/sdf/internal/pkg/navigate"
	"github.com/NoUseFreak/sdf/internal/pkg/output"
	"github.com/NoUseFreak/sdf/internal/sdf/clone"
	"github.com/google/subcommands"
)

type CloneCmd struct {
}

func (*CloneCmd) Name() string     { return "clone" }
func (*CloneCmd) Synopsis() string { return "Clone those projects in a maintainable structure." }
func (*CloneCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *CloneCmd) SetFlags(f *flag.FlagSet) {
	// f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *CloneCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) < 1 {
		output.Error("Must provide a repo")
		return subcommands.ExitFailure
	}

	repoInput := f.Args()[0]
	targetDir, err := clone.CloneRepo(repoInput)

	if err != nil {
		output.Error("%v\n", err)
		return subcommands.ExitFailure
	}

	navigate.Chdir(targetDir)

	return subcommands.ExitSuccess
}
