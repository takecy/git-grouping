package cmds

import (
	"flag"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// InfoCmd is cmd struct
type InfoCmd struct {
	Con *conf.AppConf
}

func (*InfoCmd) Name() string {
	return "info"
}

func (*InfoCmd) Synopsis() string {
	return "Show git grouping information"
}

func (*InfoCmd) Usage() string {
	return `info <group>`
}

func (p *InfoCmd) SetFlags(f *flag.FlagSet) {
}

func (p *InfoCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
