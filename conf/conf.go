package conf

import (
	"errors"
	"fmt"
	"time"

	lc "github.com/takecy/go-localconfig/conf"
	"github.com/tj/go-debug"
)

var cd = debug.Debug("ggp:config")

// AppConf is app configuration
type AppConf struct {
	// local configuration
	Lc LocalConf

	// lc
	cc *lc.Conf
}

// LocalConf is local configuration
type LocalConf struct {
	Groups []Group `json:"groups"`
}

// Get is get by name
func (gs *LocalConf) Get(name string) (g Group) {
	if gs == nil {
		return
	}

	for _, g := range gs.Groups {
		if name == g.Name {
			return g
		}
	}

	return
}

// Add is add repo to group
func (gs *LocalConf) Add(group, repo string) (err error) {
	if gs == nil {
		return errors.New("null.receiver")
	}

	cd("group:%s repo:%s", group, repo)

	for i, g := range gs.Groups {
		if group == g.Name {
			cd("matched:%s", g.Name)
			g.Repos = append(g.Repos, repo)
			gs.Groups[i].Repos = g.Repos
			return
		}
	}

	gs.Groups = append(gs.Groups, Group{ID: fmt.Sprintf("id-%d", time.Now().Unix()), Name: group, Repos: []string{repo}})
	return
}

// Remove is remove repo from group
func (gs *LocalConf) Remove(group, repo string) (err error) {
	if gs == nil {
		return errors.New("null.receiver")
	}

	cd("group:%s repo:%s", group, repo)

	for _, g := range gs.Groups {
		if group == g.Name {
			nrepos := make([]string, 0, len(g.Repos))
			for _, j := range g.Repos {
				if j == repo {
					cd("matched.continue:%s", j)
					continue
				}
				nrepos = append(nrepos, j)
			}
			g.Repos = nrepos
			return
		}
	}

	return
}

// RemoveGroup is remove group
func (gs *LocalConf) RemoveGroup(group string) (err error) {
	if gs == nil {
		return errors.New("null.receiver")
	}

	cd("group:%s", group)

	ngroups := make([]Group, 0, len(gs.Groups))
	for _, g := range gs.Groups {
		if g.Name == group {
			cd("matched.continue:%s", g.Name)
			continue
		}
		ngroups = append(ngroups, g)
	}

	gs.Groups = ngroups

	return
}

// Group is repository group
type Group struct {
	// group id
	ID string `json:"id"`

	// group name
	Name string `json:"name"`

	// grouped repositories
	Repos []string `json:"repos"`
}

// New is init
func New() *AppConf {
	return &AppConf{
		cc: lc.NewConfig("ggp", ""),
	}
}

// Load is load onfing from local
func (a *AppConf) Load() (err error) {
	err = a.cc.Load(&a.Lc)
	return
}

// Save is save config to local
func (a *AppConf) Save(c LocalConf) (err error) {
	err = a.cc.Save(&c)
	return
}
