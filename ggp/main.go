package main

import (
	"flag"

	"github.com/takecy/git-grouping/cli"
	"github.com/takecy/git-grouping/conf"
)

var (
	verbose = flag.Bool("v", false, "show verbose log")
)

type ggp struct {
	// verbose logging
	v bool

	// client
	c *cli.Cli

	// configuration
	con *conf.Conf
}

func main() {
	flag.Parse()

	err := (&ggp{
		v:   *verbose,
		c:   cli.New(),
		con: conf.New(),
	}).run()

	if err != nil {
		panic(err)
	}
}

func (g *ggp) run() (err error) {

	return
}
