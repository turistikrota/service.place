package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformDisableCommand struct {
	AdminUUID string
	Slug      string
}

type PlatformDisableResult struct{}

type PlatformDisableHandler decorator.CommandHandler[PlatformDisableCommand, *PlatformDisableResult]

type platformDisableHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformDisableHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformDisableHandler(config PlatformDisableHandlerConfig) PlatformDisableHandler {
	return decorator.ApplyCommandDecorators[PlatformDisableCommand, *PlatformDisableResult](
		platformDisableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformDisableHandler) Handle(ctx context.Context, command PlatformDisableCommand) (*PlatformDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	h.events.Disabled(platform.EventPlatformDisabled{
		AdminUUID: command.AdminUUID,
		Slug:      command.Slug,
	})
	return &PlatformDisableResult{}, nil
}
