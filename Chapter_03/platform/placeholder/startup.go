package placeholder

import (
	"platform/http"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"sync"
	"platform/http/handling"
	"platform/sessions"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.ServicesComponent{},
		&basic.StaticFileComponent{},
		//&SimpleMessageComponent{},
		handling.NewRouter(
			handling.HandlerEntry{"", NameHandler{}}, 
			handling.HandlerEntry{"", DayHandler{}},
			handling.HandlerEntry{"", CounterHandler{}},).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
	sessions.RegisterSessionService()
	results, err := services.Call(http.Serve, createPipeline())
	if(err == nil) {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}