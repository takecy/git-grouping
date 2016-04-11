package cmds

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"github.com/tj/go-debug"
	"golang.org/x/net/context"
)

var rmgd = debug.Debug("ggp:cmds:rmg")

// RmgCmd is cmd struct
type RmgCmd struct {
	Con *conf.AppConf
}

func (*RmgCmd) Name() string {
	return "rmg"
}

func (*RmgCmd) Synopsis() string {
	return "Remove group"
}

func (*RmgCmd) Usage() string {
	return `rmg <group>
`
}

func (p *RmgCmd) SetFlags(f *flag.FlagSet) {
}

func (p *RmgCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	group := f.Arg(0)
	if group == "" {
		fmt.Fprintf(os.Stderr, "group required.\n")
		return subcommands.ExitFailure
	}

	rmgd("group:%s", group)

	err := p.Con.Lc.RemoveGroup(group)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remove.group.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	rmgd("group.removed")
	rmgd("ls:%+v", p.Con.Lc)

	err = p.Con.Save(p.Con.Lc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "save.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	rmgd("saved")

	return subcommands.ExitSuccess
}
