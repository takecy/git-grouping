package conf

import (
	"github.com/takecy/git-grouping/conf"
	"github.com/takecy/go-localconfig/conf"
)

var lc *localconf.Conf

type AppConf struct {
	// config directory path
	Dir string
}

type Group struct {
	// group id
	ID string

	// group name
	Name string

	// grouped repositories
	Repos []string
}

func New() (c *AppConf) {
	conf.New()

	c = &Conf{}
	return
}

func (c *Conf) Load() (app AppConf, err error) {
	err = lconf.Load(&app)
	return
}

func (c *Conf) Save(c AppConf) (err error) {
	err = lconf.Save(&c)
	return
}
