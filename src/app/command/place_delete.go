package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	PlaceDeleteCommand struct {
		UUID string
	}
	PlaceDeleteResult  struct{}
	PlaceDeleteHandler decorator.CommandHandler[PlaceDeleteCommand, *PlaceDeleteResult]
	placeDeleteHandler struct {
		repo    place.Repository
		events  place.Events
		factory place.Factory
	}
	PlaceDeleteHandlerConfig struct {
		Repo     place.Repository
		Events   place.Events
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceDeleteHandler(config PlaceDeleteHandlerConfig) PlaceDeleteHandler {
	return decorator.ApplyCommandDecorators[PlaceDeleteCommand, *PlaceDeleteResult](
		placeDeleteHandler{
			repo:    config.Repo,
			events:  config.Events,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeDeleteHandler) Handle(ctx context.Context, command PlaceDeleteCommand) (*PlaceDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, command.UUID)
	if err != nil {
		return nil, err
	}
	return &PlaceDeleteResult{}, nil
}
