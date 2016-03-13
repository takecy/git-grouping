package cmds

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// VersionCmd is cmd struct
type VersionCmd struct {
	Con     *conf.AppConf
	Version string
}

func (*VersionCmd) Name() string {
	return "version"
}

func (*VersionCmd) Synopsis() string {
	return "Print version"
}

func (*VersionCmd) Usage() string {
	return `version
`
}

func (p *VersionCmd) SetFlags(f *flag.FlagSet) {
}

func (p *VersionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("git-grouping:v%s \nruntime:%s\n", p.Version, runtime.Version())
	return subcommands.ExitSuccess
}
