package conf

import lc "github.com/takecy/go-localconfig/conf"

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
func (a *AppConf) Load() (app LocalConf, err error) {
	err = a.cc.Load(&app)
	return
}

// Save is save config to local
func (a *AppConf) Save(c LocalConf) (err error) {
	err = a.cc.Save(&c)
	return
}
