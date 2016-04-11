package cmds

import (
	"flag"
	"fmt"
	"html/template"
	"os"

	"github.com/fatih/color"
	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

var helpers = template.FuncMap{
	"magenta": color.MagentaString,
	"yellow":  color.YellowString,
	"green":   color.GreenString,
	"black":   color.BlackString,
	"white":   color.WhiteString,
	"blue":    color.BlueString,
	"cyan":    color.CyanString,
	"red":     color.RedString,
}

const itemTmpl = `
Groups:
{{range .}}
   {{.Name | green}} ({{.ID | cyan}})
    {{range .Repos}}
    - {{. | blue}}{{end}}
    
{{end}}
`

// LsCmd is cmd struct
type LsCmd struct {
	Con *conf.AppConf
}

func (*LsCmd) Name() string {
	return "ls"
}

func (*LsCmd) Synopsis() string {
	return "List all group."
}

func (*LsCmd) Usage() string {
	return `ls
`
}

func (p *LsCmd) SetFlags(f *flag.FlagSet) {
}

func (p *LsCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := t().Execute(os.Stdout, p.Con.Lc.Groups)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err.render.template.\n")
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func t() *template.Template {
	return template.Must(template.New("group").Funcs(helpers).Parse(itemTmpl))
}
