package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
	"github.com/turistikrota/service.place/src/domain/place"
)

type (
	PlaceCreateCommand struct {
		FeatureUUIDs     []string
		Images           []place.Image
		Restorations     []place.Restoration
		Translations     map[place.Locale]*place.Translations
		AverageTimeSpent place.TimeSpent
		Coordinates      []float64
		IsPayed          bool
		Type             place.Type
	}
	PlaceCreateResult struct {
		UUID string
	}
	PlaceCreateHandler decorator.CommandHandler[PlaceCreateCommand, *PlaceCreateResult]
	placeCreateHandler struct {
		repo        place.Repository
		events      place.Events
		featureRepo feature.Repository
		factory     place.Factory
	}
	PlaceCreateHandlerConfig struct {
		Repo        place.Repository
		Events      place.Events
		FeatureRepo feature.Repository
		Factory     place.Factory
		CqrsBase    decorator.Base
	}
)

func NewPlaceCreateHandler(config PlaceCreateHandlerConfig) PlaceCreateHandler {
	return decorator.ApplyCommandDecorators[PlaceCreateCommand, *PlaceCreateResult](
		placeCreateHandler{
			repo:        config.Repo,
			events:      config.Events,
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
		Coordinates:      command.Coordinates,
		IsPayed:          command.IsPayed,
		Type:             command.Type,
		Restorations:     command.Restorations,
	})
	id, error := h.repo.Create(ctx, p)
	if error != nil {
		return nil, error
	}
	return &PlaceCreateResult{UUID: id}, nil
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
