package cmds

import (
	"flag"
	"encoding/json"
	"fmt"
	"os"

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
	return "Show configuration."
}

func (*InfoCmd) Usage() string {
	return `info <group>
`
}

func (p *InfoCmd) SetFlags(f *flag.FlagSet) {
}

func (p *InfoCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	b, err := json.Marshal(p.Con.Lc)
	if err != nil {
		return subcommands.ExitFailure
	}
	fmt.Fprintf(os.Stdout, "grouping info: %s\n", string(b))
	return subcommands.ExitSuccess
}
