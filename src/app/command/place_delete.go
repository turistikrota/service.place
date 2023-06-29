package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlaceDeleteCommand struct{}

type PlaceDeleteResult struct{}

type PlaceDeleteHandler decorator.CommandHandler[PlaceDeleteCommand, *PlaceDeleteResult]

type placeDeleteHandler struct {
	repo    place.Repository
	factory place.Factory
}

type PlaceDeleteHandlerConfig struct {
	Repo     place.Repository
	Factory  place.Factory
	CqrsBase decorator.Base
}

func NewPlaceDeleteHandler(config PlaceDeleteHandlerConfig) PlaceDeleteHandler {
	return decorator.ApplyCommandDecorators[PlaceDeleteCommand, *PlaceDeleteResult](
		placeDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeDeleteHandler) Handle(ctx context.Context, command PlaceDeleteCommand) (*PlaceDeleteResult, *i18np.Error) {
	return &PlaceDeleteResult{}, nil
}
