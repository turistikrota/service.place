package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceUpdateCommand struct {
		UUID             string
		FeatureUUIDs     []string
		Images           []place.Image
		Translations     map[place.Locale]place.Translations
		AverageTimeSpent place.TimeSpent
		Review           place.Review
		Coordinates      []float64
		IsPayed          bool
	}
	PlaceUpdateResult  struct{}
	PlaceUpdateHandler decorator.CommandHandler[PlaceUpdateCommand, *PlaceUpdateResult]
	placeUpdateHandler struct {
		repo    place.Repository
		factory place.Factory
	}
	PlaceUpdateHandlerConfig struct {
		Repo     place.Repository
		Factory  place.Factory
		CqrsBase decorator.Base
	}
)

func NewPlaceUpdateHandler(config PlaceUpdateHandlerConfig) PlaceUpdateHandler {
	return decorator.ApplyCommandDecorators[PlaceUpdateCommand, *PlaceUpdateResult](
		placeUpdateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeUpdateHandler) Handle(ctx context.Context, command PlaceUpdateCommand) (*PlaceUpdateResult, *i18np.Error) {
	p := h.factory.New(place.NewConfig{
		FeatureUUIDs:     command.FeatureUUIDs,
		Images:           command.Images,
		Translations:     command.Translations,
		AverageTimeSpent: command.AverageTimeSpent,
		Review:           command.Review,
		Coordinates:      command.Coordinates,
		IsPayed:          command.IsPayed,
	})
	err := h.repo.Update(ctx, command.UUID, p)
	if err != nil {
		return nil, err
	}
	return &PlaceUpdateResult{}, nil
}
