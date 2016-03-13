package cmds

import (
	"flag"

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
	return subcommands.ExitSuccess
}
