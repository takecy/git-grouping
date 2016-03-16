package cmds

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// LsCmd is cmd struct
type LsCmd struct {
	Con *conf.AppConf
}

func (*LsCmd) Name() string {
	return "ls"
}

func (*LsCmd) Synopsis() string {
	return "List all group."
}

func (*LsCmd) Usage() string {
	return `ls
`
}

func (p *LsCmd) SetFlags(f *flag.FlagSet) {
}

func (p *LsCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, g := range p.Con.Lc.Groups {
		fmt.Fprintf(os.Stdout, "group:%s\n", g.Name)

		for _, r := range g.Repos {
			fmt.Fprintf(os.Stdout, "repo:%s\n", r)
		}
	}

	return subcommands.ExitSuccess
}
