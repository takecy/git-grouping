package conf

import lc "github.com/takecy/go-localconfig/conf"

var cc *lc.Conf

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

// New is init
func New() *AppConf {
	cc = lc.NewConfig("ggp", "")
	app, err := Load()
	if err != nil {
		panic(err)
	}
	return &app
}

// Load is load onfing from local
func Load() (app AppConf, err error) {
	err = cc.Load(&app)
	return
}

// Save is save config to local
func Save(c AppConf) (err error) {
	err = cc.Save(&c)
	return
}
