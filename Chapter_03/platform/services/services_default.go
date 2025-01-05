package services

import (
	//"platform/logging"
	"platform/config"
	"platform/templates"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if (loadErr != nil) {
			panic(loadErr)
		}
		return
	})

	err = AddSingleton(
		func(c config.Configuration) templates.TemplateExecutor {
			templates.LoadTemplates(c)
			return &templates.LayoutTemplateProcessor{}
		})
	if (err != nil) {
		panic(err)
	}
}