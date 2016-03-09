package cmds

import (
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

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
	return `add <group name> <dir path>`
}

func (p *AddCmd) SetFlags(f *flag.FlagSet) {
}

func (p *AddCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		arg = strings.ToUpper(arg)
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
