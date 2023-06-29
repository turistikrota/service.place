package feature

import (
	"api.turistikrota.com/place/src/config"
	"github.com/turistikrota/service.shared/events"
)

type Events interface{}

type featureEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &featureEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
	}
}
