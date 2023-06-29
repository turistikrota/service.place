package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformUpdateCommand struct {
	AdminUUID string
	Slug      string
	Name      string
	Regexp    string
	Prefix    string
}

type PlatformUpdateResult struct{}

type PlatformUpdateHandler decorator.CommandHandler[PlatformUpdateCommand, *PlatformUpdateResult]

type platformUpdateHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformUpdateHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformUpdateHandler(config PlatformUpdateHandlerConfig) PlatformUpdateHandler {
	return decorator.ApplyCommandDecorators[PlatformUpdateCommand, *PlatformUpdateResult](
		platformUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformUpdateHandler) Handle(ctx context.Context, command PlatformUpdateCommand) (*PlatformUpdateResult, *i18np.Error) {
	p := h.factory.NewPlatform(command.Name, command.Regexp, command.Prefix)
	err := h.factory.Validate(p)
	if err != nil {
		return nil, err
	}
	if p.Name != command.Name {
		_, err = h.repo.GetBySlug(ctx, p.Name)
		if err == nil {
			return nil, h.factory.Errors.AlreadyExists("platform")
		}
	}
	err = h.repo.Update(ctx, command.Slug, p)
	if err != nil {
		return nil, err
	}
	h.events.Updated(platform.EventPlatformUpdated{
		AdminUUID: command.AdminUUID,
		Slug:      command.Slug,
		Name:      command.Name,
		Regexp:    command.Regexp,
		Prefix:    command.Prefix,
	})
	return &PlatformUpdateResult{}, nil
}
