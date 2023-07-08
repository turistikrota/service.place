package event_stream

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.place/src/app"
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
