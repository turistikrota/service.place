package event_stream

import (
	"api.turistikrota.com/place/src/app"
	"github.com/turistikrota/service.shared/events"
)

type Server struct {
	app    app.Application
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
	}
}

func (s Server) Load() {
	// s.engine.Subscribe(s.Topics.Created, s.ListenEmptyCreated)
}
