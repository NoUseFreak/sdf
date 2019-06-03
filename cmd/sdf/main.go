package main

import (
	"github.com/NoUseFreak/sdf/internal/sdf/command"
)

func main() {
	command.Execute()
}

// package main

// import (
// // 	"context"
// // 	"flag"
// // 	"os"

// // 	"github.com/NoUseFreak/sdf/internal/sdf/command"
// // 	"github.com/NoUseFreak/sdf/internal/sdf/config"
// // 	"github.com/google/subcommands"
// )

// func main() {

// 	// sdfConfig := config.InitConfig()

// 	// subcommands.Register(subcommands.HelpCommand(), "")
// 	// subcommands.Register(subcommands.FlagsCommand(), "")
// 	// subcommands.Register(subcommands.CommandsCommand(), "")
// 	// subcommands.Register(&command.InfoCmd{}, "")
// 	// subcommands.Register(&command.CloneCmd{}, "")
// 	// subcommands.Register(&command.SetupCmd{}, "")

// 	// flag.StringVar(&sdfConfig.Profile, "profile", "main", "Profile")

// 	// flag.Parse()
// 	// sdfConfig.ReadConfig(sdfConfig.Profile)

// 	// ctx := context.Background()
// 	// os.Exit(int(subcommands.Execute(ctx)))
// }
