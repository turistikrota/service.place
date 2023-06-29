package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformTranslationCreateCommand struct {
	AdminUUID         string
	Slug              string
	Locale            platform.Locale
	LocaleName        string
	LocalePlaceholder string
	LocaleDescription string
}

type PlatformTranslationCreateResult struct{}

type PlatformTranslationCreateHandler decorator.CommandHandler[PlatformTranslationCreateCommand, *PlatformTranslationCreateResult]

type platformTranslationCreateHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformTranslationCreateHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformTranslationCreateHandler(config PlatformTranslationCreateHandlerConfig) PlatformTranslationCreateHandler {
	return decorator.ApplyCommandDecorators[PlatformTranslationCreateCommand, *PlatformTranslationCreateResult](
		platformTranslationCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformTranslationCreateHandler) Handle(ctx context.Context, command PlatformTranslationCreateCommand) (*PlatformTranslationCreateResult, *i18np.Error) {
	p, err := h.repo.GetBySlug(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	_, exist := p.Translations[command.Locale]
	if exist {
		return nil, h.factory.Errors.TranslationAlreadyExists(command.Locale.String())
	}
	if command.LocaleName == "" {
		command.LocaleName = p.Name
	}
	err = h.repo.TranslationCreate(ctx, command.Slug, command.Locale, platform.Translations{
		Name:        command.LocaleName,
		Placeholder: command.LocalePlaceholder,
		Description: command.LocaleDescription,
	})
	if err != nil {
		return nil, err
	}
	h.events.TranslationCreated(platform.EventPlatformTranslationCreated{
		AdminUUID:         command.AdminUUID,
		Slug:              command.Slug,
		Locale:            command.Locale.String(),
		LocaleName:        command.LocaleName,
		LocalePlaceholder: command.LocalePlaceholder,
		LocaleDescription: command.LocaleDescription,
	})
	return &PlatformTranslationCreateResult{}, nil
}
