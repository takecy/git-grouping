package cmds

import (
	"flag"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// ConfigCmd is cmd struct
type ConfigCmd struct {
	Con *conf.AppConf
}

func (*ConfigCmd) Name() string {
	return "config"
}

func (*ConfigCmd) Synopsis() string {
	return "Command configuration"
}

func (*ConfigCmd) Usage() string {
	return `config`
}

func (p *ConfigCmd) SetFlags(f *flag.FlagSet) {
}

func (p *ConfigCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
