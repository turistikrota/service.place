package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformEnableCommand struct {
	AdminUUID string
	Slug      string
}

type PlatformEnableResult struct{}

type PlatformEnableHandler decorator.CommandHandler[PlatformEnableCommand, *PlatformEnableResult]

type platformEnableHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformEnableHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformEnableHandler(config PlatformEnableHandlerConfig) PlatformEnableHandler {
	return decorator.ApplyCommandDecorators[PlatformEnableCommand, *PlatformEnableResult](
		platformEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformEnableHandler) Handle(ctx context.Context, command PlatformEnableCommand) (*PlatformEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	h.events.Enabled(platform.EventPlatformEnabled{
		AdminUUID: command.AdminUUID,
		Slug:      command.Slug,
	})
	return &PlatformEnableResult{}, nil
}
