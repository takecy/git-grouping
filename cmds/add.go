package cmds

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

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
	return "Add repository to group"
}

func (*AddCmd) Usage() string {
	return `add <group name> <dir path>
`
}

func (p *AddCmd) SetFlags(f *flag.FlagSet) {
}

func (p *AddCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() < 2 {
		fmt.Fprint(os.Stderr, "not enough args.\n")
		return subcommands.ExitFailure
	}

	group := f.Arg(0)
	dirPath := f.Arg(1)

	if !path.IsAbs(dirPath) {
		aPath, err := filepath.Abs(dirPath)
		if err != nil {
			fmt.Fprint(os.Stderr, "invalid path.\n")
			return subcommands.ExitFailure
		}
		dirPath = aPath
	}

	ad("group:%s dir:%s", group, dirPath)

	if group == "" && dirPath == "" {
		fmt.Fprintf(os.Stdout, "group and directory required.\n")
		return subcommands.ExitFailure
	}

	err := p.Con.Lc.Add(group, dirPath)
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
