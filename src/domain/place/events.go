package place

import (
	"api.turistikrota.com/place/src/config"
	"github.com/turistikrota/service.shared/events"
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
