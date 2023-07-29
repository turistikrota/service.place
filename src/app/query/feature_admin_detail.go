package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.place/src/domain/feature"
)

type (
	AdminFeatureDetailQuery struct {
		UUID string
	}
	AdminFeatureDetailResult struct {
		Feature *feature.Entity
	}
	AdminFeatureDetailHandler decorator.QueryHandler[AdminFeatureDetailQuery, *AdminFeatureDetailResult]
	adminFeatureDetailHandler struct {
		repo feature.Repository
	}
	AdminFeatureDetailHandlerConfig struct {
		Repo     feature.Repository
		CqrsBase decorator.Base
	}
)

func NewAdminFeatureDetailHandler(config AdminFeatureDetailHandlerConfig) AdminFeatureDetailHandler {
	return decorator.ApplyQueryDecorators[AdminFeatureDetailQuery, *AdminFeatureDetailResult](
		adminFeatureDetailHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h adminFeatureDetailHandler) Handle(ctx context.Context, query AdminFeatureDetailQuery) (*AdminFeatureDetailResult, *i18np.Error) {
	res, err := h.repo.GetByUUID(ctx, query.UUID)
	if err != nil {
		return nil, err
	}
	return &AdminFeatureDetailResult{
		Feature: res,
	}, nil
}
