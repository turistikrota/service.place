package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformTranslationRemoveCommand struct {
	AdminUUID string
	Slug      string
	Locale    platform.Locale
}

type PlatformTranslationRemoveResult struct{}

type PlatformTranslationRemoveHandler decorator.CommandHandler[PlatformTranslationRemoveCommand, *PlatformTranslationRemoveResult]

type platformTranslationRemoveHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformTranslationRemoveHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformTranslationRemoveHandler(config PlatformTranslationRemoveHandlerConfig) PlatformTranslationRemoveHandler {
	return decorator.ApplyCommandDecorators[PlatformTranslationRemoveCommand, *PlatformTranslationRemoveResult](
		platformTranslationRemoveHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformTranslationRemoveHandler) Handle(ctx context.Context, command PlatformTranslationRemoveCommand) (*PlatformTranslationRemoveResult, *i18np.Error) {
	p, err := h.repo.GetBySlug(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	_, exist := p.Translations[command.Locale]
	if !exist {
		return nil, h.factory.Errors.TranslationNotExists(command.Locale.String())
	}
	err = h.repo.TranslationDelete(ctx, command.Slug, command.Locale)
	if err != nil {
		return nil, err
	}
	h.events.TranslationRemoved(platform.EventPlatformTranslationRemoved{
		AdminUUID: command.AdminUUID,
		Slug:      command.Slug,
		Locale:    command.Locale.String(),
	})
	return &PlatformTranslationRemoveResult{}, nil
}
