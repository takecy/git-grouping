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
	g.con.Load()

	subcommands.Register(subcommands.HelpCommand(), "base")
	subcommands.Register(subcommands.FlagsCommand(), "base")
	subcommands.Register(subcommands.CommandsCommand(), "base")
	subcommands.Register(&cmds.VersionCmd{Version: ver}, "config")
	subcommands.Register(&cmds.AddCmd{Con: g.con}, "config")
	subcommands.Register(&cmds.GitCmd{Con: g.con}, "git")
	subcommands.Register(&cmds.InfoCmd{Con: g.con}, "config")
	subcommands.Register(&cmds.LsCmd{Con: g.con}, "config")
	subcommands.Register(&cmds.RmCmd{Con: g.con}, "config")

	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
