package cmds

import (
	"flag"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// ExecCmd is cmd struct
type ExecCmd struct {
	Con *conf.AppConf
}

func (*ExecCmd) Name() string {
	return "exec"
}

func (*ExecCmd) Synopsis() string {
	return "Exec git commands"
}

func (*ExecCmd) Usage() string {
	return `exec <git command> [git options]`
}

func (p *ExecCmd) SetFlags(f *flag.FlagSet) {
}

func (p *ExecCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
