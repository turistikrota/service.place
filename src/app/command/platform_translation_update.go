package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformTranslationUpdateCommand struct {
	AdminUUID         string
	Slug              string
	Locale            platform.Locale
	LocaleName        string
	LocalePlaceholder string
	LocaleDescription string
}

type PlatformTranslationUpdateResult struct{}

type PlatformTranslationUpdateHandler decorator.CommandHandler[PlatformTranslationUpdateCommand, *PlatformTranslationUpdateResult]

type platformTranslationUpdateHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformTranslationUpdateHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformTranslationUpdateHandler(config PlatformTranslationUpdateHandlerConfig) PlatformTranslationUpdateHandler {
	return decorator.ApplyCommandDecorators[PlatformTranslationUpdateCommand, *PlatformTranslationUpdateResult](
		platformTranslationUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformTranslationUpdateHandler) Handle(ctx context.Context, command PlatformTranslationUpdateCommand) (*PlatformTranslationUpdateResult, *i18np.Error) {
	p, err := h.repo.GetBySlug(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	t, exist := p.Translations[command.Locale]
	if !exist {
		return nil, h.factory.Errors.TranslationNotExists(command.Locale.String())
	}
	if command.LocaleName == "" {
		command.LocaleName = t.Name
	}
	err = h.repo.TranslationUpdate(ctx, command.Slug, command.Locale, platform.Translations{
		Name:        command.LocaleName,
		Placeholder: command.LocalePlaceholder,
		Description: command.LocaleDescription,
	})
	if err != nil {
		return nil, err
	}
	h.events.TranslationUpdated(platform.EventPlatformTranslationUpdated{
		AdminUUID:         command.AdminUUID,
		Slug:              command.Slug,
		Locale:            command.Locale.String(),
		LocaleName:        command.LocaleName,
		LocalePlaceholder: command.LocalePlaceholder,
		LocaleDescription: command.LocaleDescription,
	})
	return &PlatformTranslationUpdateResult{}, nil
}
