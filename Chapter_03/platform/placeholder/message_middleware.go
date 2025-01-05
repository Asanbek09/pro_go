package placeholder

import (
	//"io"
	//"errors"
	"platform/pipeline"
	//"platform/services"
	"platform/config"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

func(lc *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (c *SimpleMessageComponent) Init() {
	c.Message = c.Configuration.GetStringDefault("main:message", "Default Message")
}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext), executor templates.TemplateExecutor) {
	err := executor.ExecTemplate(ctx.ResponseWriter, "simple_message.html", c.Message)
	if (err != nil) {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}