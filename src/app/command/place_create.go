package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type PlaceCreateCommand struct {
	FeatureUUIDs     []string
	Images           []place.Image
	Translations     map[place.Locale]place.Translations
	AverageTimeSpent place.TimeSpent
	Review           place.Review
	Coordinates      []float64
	IsPayed          bool
}

type PlaceCreateResult struct{}

type PlaceCreateHandler decorator.CommandHandler[PlaceCreateCommand, *PlaceCreateResult]

type placeCreateHandler struct {
	repo    place.Repository
	factory place.Factory
}

type PlaceCreateHandlerConfig struct {
	Repo     place.Repository
	Factory  place.Factory
	CqrsBase decorator.Base
}

func NewPlaceCreateHandler(config PlaceCreateHandlerConfig) PlaceCreateHandler {
	return decorator.ApplyCommandDecorators[PlaceCreateCommand, *PlaceCreateResult](
		placeCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeCreateHandler) Handle(ctx context.Context, command PlaceCreateCommand) (*PlaceCreateResult, *i18np.Error) {
	p := h.factory.New(place.NewConfig{
		FeatureUUIDs:     command.FeatureUUIDs,
		Images:           command.Images,
		Translations:     command.Translations,
		AverageTimeSpent: command.AverageTimeSpent,
		Review:           command.Review,
		Coordinates:      command.Coordinates,
		IsPayed:          command.IsPayed,
	})
	err := h.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return &PlaceCreateResult{}, nil
}
