package platform

import (
	"api.turistikrota.com/account/src/config"
	"github.com/turistikrota/service.shared/events"
)

type Events interface {
	Created(platform EventPlatformCreated)
	Updated(platform EventPlatformUpdated)
	Deleted(platform EventPlatformDeleted)
	Disabled(platform EventPlatformDisabled)
	Enabled(platform EventPlatformEnabled)
	TranslationCreated(platform EventPlatformTranslationCreated)
	TranslationUpdated(platform EventPlatformTranslationUpdated)
	TranslationRemoved(platform EventPlatformTranslationRemoved)
}

type (
	EventPlatformCreated struct {
		AdminUUID string `json:"admin_uuid"`
		Name      string `json:"name"`
		Regexp    string `json:"regexp"`
		Prefix    string `json:"prefix"`
	}
	EventPlatformUpdated struct {
		AdminUUID string `json:"admin_uuid"`
		Slug      string `json:"slug"`
		Name      string `json:"name"`
		Regexp    string `json:"regexp"`
		Prefix    string `json:"prefix"`
	}
	EventPlatformDisabled struct {
		AdminUUID string `json:"admin_uuid"`
		Slug      string `json:"slug"`
	}
	EventPlatformEnabled struct {
		AdminUUID string `json:"admin_uuid"`
		Slug      string `json:"slug"`
	}
	EventPlatformDeleted struct {
		AdminUUID string `json:"admin_uuid"`
		Slug      string `json:"slug"`
	}
	EventPlatformTranslationCreated struct {
		AdminUUID         string `json:"admin_uuid"`
		Slug              string `json:"slug"`
		Locale            string `json:"locale"`
		LocaleName        string `json:"locale_name"`
		LocalePlaceholder string `json:"locale_placeholder"`
		LocaleDescription string `json:"locale_description"`
	}
	EventPlatformTranslationUpdated struct {
		AdminUUID         string `json:"admin_uuid"`
		Slug              string `json:"slug"`
		Locale            string `json:"locale"`
		LocaleName        string `json:"locale_name"`
		LocalePlaceholder string `json:"locale_placeholder"`
		LocaleDescription string `json:"locale_description"`
	}
	EventPlatformTranslationRemoved struct {
		AdminUUID string `json:"admin_uuid"`
		Slug      string `json:"slug"`
		Locale    string `json:"locale"`
	}
)

type accountEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &accountEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
	}
}

func (e *accountEvents) Created(platform EventPlatformCreated) {
	_ = e.publisher.Publish(e.topics.Platform.Created, platform)
}

func (e *accountEvents) Updated(platform EventPlatformUpdated) {
	_ = e.publisher.Publish(e.topics.Platform.Updated, platform)
}

func (e *accountEvents) Disabled(platform EventPlatformDisabled) {
	_ = e.publisher.Publish(e.topics.Platform.Disabled, platform)
}

func (e *accountEvents) Enabled(platform EventPlatformEnabled) {
	_ = e.publisher.Publish(e.topics.Platform.Enabled, platform)
}

func (e *accountEvents) TranslationCreated(platform EventPlatformTranslationCreated) {
	_ = e.publisher.Publish(e.topics.Platform.TranslationCreated, platform)
}

func (e *accountEvents) TranslationUpdated(platform EventPlatformTranslationUpdated) {
	_ = e.publisher.Publish(e.topics.Platform.TranslationUpdated, platform)
}

func (e *accountEvents) TranslationRemoved(platform EventPlatformTranslationRemoved) {
	_ = e.publisher.Publish(e.topics.Platform.TranslationRemoved, platform)
}

func (e *accountEvents) Deleted(platform EventPlatformDeleted) {
	_ = e.publisher.Publish(e.topics.Platform.Deleted, platform)
}
