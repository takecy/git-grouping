package main

import (
	"flag"
	"os"

	"golang.org/x/net/context"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/cmds"
	"github.com/takecy/git-grouping/conf"
)

var (
	verbose = flag.Bool("v", false, "Print verbose log")
	ver     = "0.0.1"
)

type ggp struct {
	// verbose logging
	v bool

	// configuration
	con *conf.AppConf
}

func main() {
	(&ggp{
		v:   *verbose,
		con: conf.New(),
	}).run()
}

func (g *ggp) run() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&cmds.VersionCmd{Version: ver}, "")
	subcommands.Register(&cmds.AddCmd{}, "")
	subcommands.Register(&cmds.ConfigCmd{}, "")
	subcommands.Register(&cmds.ExecCmd{}, "")
	subcommands.Register(&cmds.InfoCmd{}, "")
	subcommands.Register(&cmds.LsCmd{}, "")
	subcommands.Register(&cmds.RmCmd{}, "")

	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
