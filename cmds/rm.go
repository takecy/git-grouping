package cmds

import (
	"flag"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// RmCmd is cmd struct
type RmCmd struct {
	Con *conf.AppConf
}

func (*RmCmd) Name() string {
	return "rm"
}

func (*RmCmd) Synopsis() string {
	return "Remove form group"
}

func (*RmCmd) Usage() string {
	return `rm <group> <repository>`
}

func (p *RmCmd) SetFlags(f *flag.FlagSet) {
}

func (p *RmCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
