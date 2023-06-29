package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformDeleteCommand struct {
	AdminUUID string
	Slug      string
}

type PlatformDeleteResult struct{}

type PlatformDeleteHandler decorator.CommandHandler[PlatformDeleteCommand, *PlatformDeleteResult]

type platformDeleteHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformDeleteHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformDeleteHandler(config PlatformDeleteHandlerConfig) PlatformDeleteHandler {
	return decorator.ApplyCommandDecorators[PlatformDeleteCommand, *PlatformDeleteResult](
		platformDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformDeleteHandler) Handle(ctx context.Context, command PlatformDeleteCommand) (*PlatformDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, command.Slug)
	if err != nil {
		return nil, err
	}
	h.events.Deleted(platform.EventPlatformDeleted{
		AdminUUID: command.AdminUUID,
		Slug:      command.Slug,
	})
	return &PlatformDeleteResult{}, nil
}
