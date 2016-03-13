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
	Con *conf.AppConf
}

func (*GitCmd) Name() string {
	return "git"
}

func (*GitCmd) Synopsis() string {
	return "Execute git commands to group of repository"
}

func (*GitCmd) Usage() string {
	return `git <git command> [git options]
`
}

func (p *GitCmd) SetFlags(f *flag.FlagSet) {
}

func (p *GitCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Fprint(os.Stderr, "git command required.\n")
		return subcommands.ExitFailure
	}

	gitCmd := f.Arg(0)
	gitOps := f.Args()[1:]
	cmdArgs := append([]string{gitCmd}, gitOps...)
	c := exec.Command("git", cmdArgs...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()

	if err != nil {
		return subcommands.ExitFailure
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
