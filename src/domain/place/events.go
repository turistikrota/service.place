package place

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.place/src/config"
)

type Events interface{}

type placeEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &placeEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
	}
}
