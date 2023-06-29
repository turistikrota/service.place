package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlatformCreateCommand struct {
	AdminUUID string
	Name      string
	Regexp    string
	Prefix    string
}

type PlatformCreateResult struct{}

type PlatformCreateHandler decorator.CommandHandler[PlatformCreateCommand, *PlatformCreateResult]

type platformCreateHandler struct {
	repo    platform.Repository
	factory platform.Factory
	events  platform.Events
}

type PlatformCreateHandlerConfig struct {
	Repo     platform.Repository
	Factory  platform.Factory
	Events   platform.Events
	CqrsBase decorator.Base
}

func NewPlatformCreateHandler(config PlatformCreateHandlerConfig) PlatformCreateHandler {
	return decorator.ApplyCommandDecorators[PlatformCreateCommand, *PlatformCreateResult](
		platformCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h platformCreateHandler) Handle(ctx context.Context, command PlatformCreateCommand) (*PlatformCreateResult, *i18np.Error) {
	p := h.factory.NewPlatform(command.Name, command.Regexp, command.Prefix)
	err := h.factory.Validate(p)
	if err != nil {
		return nil, err
	}
	_, err = h.repo.GetBySlug(ctx, p.Slug)
	if err == nil {
		return nil, h.factory.Errors.AlreadyExists("platform")
	}
	err = h.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	h.events.Created(platform.EventPlatformCreated{
		AdminUUID: command.AdminUUID,
		Name:      command.Name,
		Regexp:    command.Regexp,
		Prefix:    command.Prefix,
	})
	return &PlatformCreateResult{}, nil
}
