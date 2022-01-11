package dummy

import "go-dummyapp-moysklad/internal/app"

type Dummy struct {
	app.BaseApp
	InfoMessage string `yaml:"info_message"`
	Store       string `yaml:"store"`
}

func (d *Dummy) GetBaseApp() *app.BaseApp {
	return &d.BaseApp
}
