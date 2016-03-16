package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/context"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/cmds"
	"github.com/takecy/git-grouping/conf"
	"github.com/tj/go-debug"
)

var (
	d = debug.Debug("ggp:main")
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
	d("load local conf")
	err := g.con.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "load.failed:%v", err)
		os.Exit(1)
	}
	d("loaded")
	d("local conf:%+v", g.con.Lc)

	subcommands.Register(subcommands.HelpCommand(), "base")
	subcommands.Register(subcommands.FlagsCommand(), "base")
	subcommands.Register(subcommands.CommandsCommand(), "base")
	subcommands.Register(&cmds.VersionCmd{Version: ver}, "ggp")
	subcommands.Register(&cmds.AddCmd{Con: g.con}, "ggp")
	subcommands.Register(&cmds.GitCmd{Con: g.con}, "ggp")
	subcommands.Register(&cmds.InfoCmd{Con: g.con}, "ggp")
	subcommands.Register(&cmds.LsCmd{Con: g.con}, "ggp")
	subcommands.Register(&cmds.RmCmd{Con: g.con}, "ggp")

	flag.Parse()

	d("registered subcommands")

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
