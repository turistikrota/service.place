package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlaceEnableCommand struct {
	UUID string
}

type PlaceEnableResult struct{}

type PlaceEnableHandler decorator.CommandHandler[PlaceEnableCommand, *PlaceEnableResult]

type placeEnableHandler struct {
	repo    place.Repository
	factory place.Factory
}

type PlaceEnableHandlerConfig struct {
	Repo     place.Repository
	Factory  place.Factory
	CqrsBase decorator.Base
}

func NewPlaceEnableHandler(config PlaceEnableHandlerConfig) PlaceEnableHandler {
	return decorator.ApplyCommandDecorators[PlaceEnableCommand, *PlaceEnableResult](
		placeEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeEnableHandler) Handle(ctx context.Context, command PlaceEnableCommand) (*PlaceEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &PlaceEnableResult{}, nil
}
