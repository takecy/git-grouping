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

var ad = debug.Debug("ggp:cmds:add")

// AddCmd is cmd struct
type AddCmd struct {
	Con *conf.AppConf
}

func (*AddCmd) Name() string {
	return "add"
}

func (*AddCmd) Synopsis() string {
	return "Add group to repository"
}

func (*AddCmd) Usage() string {
	return `add <group name> <dir path>
`
}

func (p *AddCmd) SetFlags(f *flag.FlagSet) {
}

func (p *AddCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	group := f.Arg(0)
	dir := f.Arg(1)

	ad("group:%s dir:%s", group, dir)

	if group == "" && dir == "" {
		fmt.Fprintf(os.Stdout, "group and directory required.\n")
		return subcommands.ExitFailure
	}

	err := p.Con.Lc.Add(group, dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "add.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	ad("added")
	ad("ls:%+v", p.Con.Lc)

	err = p.Con.Save(p.Con.Lc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "save.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	ad("saved")

	return subcommands.ExitSuccess
}
