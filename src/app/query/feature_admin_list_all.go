package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type (
	AdminFeatureListAllQuery  struct{}
	AdminFeatureListAllResult struct {
		Features []*feature.Entity
	}
	AdminFeatureListAllHandler decorator.QueryHandler[AdminFeatureListAllQuery, *AdminFeatureListAllResult]
	adminFeatureListAllHandler struct {
		repo  feature.Repository
	}
	AdminFeatureListAllHandlerConfig struct {
		Repo     feature.Repository
		CqrsBase decorator.Base
	}
)

func NewAdminFeatureListAllHandler(config AdminFeatureListAllHandlerConfig) AdminFeatureListAllHandler {
	return decorator.ApplyQueryDecorators[AdminFeatureListAllQuery, *AdminFeatureListAllResult](
		adminFeatureListAllHandler{
			repo:  config.Repo,
		},
		config.CqrsBase,
	)
}

func (h adminFeatureListAllHandler) Handle(ctx context.Context, query AdminFeatureListAllQuery) (*AdminFeatureListAllResult, *i18np.Error) {
	res, err := h.repo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return &AdminFeatureListAllResult{
		Features: res,
	}, nil
}