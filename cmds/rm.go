package cmds

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

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
	return "Remove repository form group"
}

func (*RmCmd) Usage() string {
	return `rm <group> <repository>
`
}

func (p *RmCmd) SetFlags(f *flag.FlagSet) {
}

func (p *RmCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
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

	err := p.Con.Lc.Remove(group, dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remove.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	ad("removed")
	ad("ls:%+v", p.Con.Lc)

	err = p.Con.Save(p.Con.Lc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "save.failed:%v\n", err)
		return subcommands.ExitFailure
	}
	ad("saved")

	return subcommands.ExitSuccess
}
