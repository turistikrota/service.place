package feature

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.place/src/config"
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
