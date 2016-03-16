package cmds

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/subcommands"
	"github.com/takecy/git-grouping/conf"
	"golang.org/x/net/context"
)

// GitCmd is cmd struct
type GitCmd struct {
	Con   *conf.AppConf
	group string
}

func (*GitCmd) Name() string {
	return "git"
}

func (*GitCmd) Synopsis() string {
	return "Execute git commands to group of repository"
}

func (*GitCmd) Usage() string {
	return `git [--gg group_name] <git command> [git options]
`
}

func (p *GitCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.group, "gg", "", "Target group")
}

func (p *GitCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Fprint(os.Stderr, "git command required.\n")
		return subcommands.ExitFailure
	}

	group := p.group
	if group == "" {
		fmt.Fprintf(os.Stderr, "group name required.\n")
		return subcommands.ExitFailure
	}

	g := p.Con.Lc.Get(group)

	gitCmd := f.Arg(1)
	gitOps := f.Args()[2:]
	cmdArgs := append([]string{gitCmd}, gitOps...)

	for _, d := range g.Repos {
		os.Chdir(d)
		cd, _ := os.Getwd()
		fmt.Fprintf(os.Stdout, "repo:%v\n", cd)

		c := exec.Command("git", cmdArgs...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()

		if err != nil {
			return subcommands.ExitFailure
		}
	}

	return subcommands.ExitSuccess
}

// hasGit is check git command
func hasGit() error {
	s, err := exec.LookPath("git")
	fmt.Fprintf(os.Stdout, "git path:/%s", s)
	return err
}

// HasHub is check hub command
func hasHub() error {
	s, err := exec.LookPath("hub")
	fmt.Fprintf(os.Stdout, "hub path:/%s", s)
	return err
}
