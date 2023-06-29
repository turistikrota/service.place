package command

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type (
	PlaceCreateCommand struct {
		FeatureUUIDs     []string
		Images           []place.Image
		Translations     map[place.Locale]place.Translations
		AverageTimeSpent place.TimeSpent
		Review           place.Review
		Coordinates      []float64
		IsPayed          bool
	}
	PlaceCreateResult  struct{}
	PlaceCreateHandler decorator.CommandHandler[PlaceCreateCommand, *PlaceCreateResult]
	placeCreateHandler struct {
		repo        place.Repository
		featureRepo feature.Repository
		factory     place.Factory
	}
	PlaceCreateHandlerConfig struct {
		Repo        place.Repository
		FeatureRepo feature.Repository
		Factory     place.Factory
		CqrsBase    decorator.Base
	}
)

func NewPlaceCreateHandler(config PlaceCreateHandlerConfig) PlaceCreateHandler {
	return decorator.ApplyCommandDecorators[PlaceCreateCommand, *PlaceCreateResult](
		placeCreateHandler{
			repo:        config.Repo,
			featureRepo: config.FeatureRepo,
			factory:     config.Factory,
		},
		config.CqrsBase,
	)
}

func (h placeCreateHandler) Handle(ctx context.Context, command PlaceCreateCommand) (*PlaceCreateResult, *i18np.Error) {
	err := h.checkFeatureUUIDs(ctx, command.FeatureUUIDs)
	if err != nil {
		return nil, err
	}
	p := h.factory.New(place.NewConfig{
		FeatureUUIDs:     command.FeatureUUIDs,
		Images:           command.Images,
		Translations:     command.Translations,
		AverageTimeSpent: command.AverageTimeSpent,
		Review:           command.Review,
		Coordinates:      command.Coordinates,
		IsPayed:          command.IsPayed,
	})
	err = h.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return &PlaceCreateResult{}, nil
}

func (h placeCreateHandler) checkFeatureUUIDs(ctx context.Context, uuids []string) *i18np.Error {
	res, err := h.featureRepo.GetByUUIDs(ctx, uuids)
	if err != nil {
		return err
	}
	if len(res) != len(uuids) {
		notFounds := make([]string, 0, len(uuids))
		for _, uuid := range uuids {
			for _, r := range res {
				if r.UUID == uuid {
					continue
				}
				notFounds = append(notFounds, uuid)
			}
		}
		return h.factory.Errors.FeatureUUIDsNotFound(notFounds)
	}
	return nil
}
