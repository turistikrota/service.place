package place

import (
	"context"

	"api.turistikrota.com/place/src/domain/place"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/types/list"
)

func (r *repo) Create(ctx context.Context, e *place.Entity) *i18np.Error {
	return nil
}

func (r *repo) Update(ctx context.Context, uuid string, e *place.Entity) *i18np.Error {
	return nil
}

func (r *repo) Disable(ctx context.Context, uuid string) *i18np.Error {
	return nil
}

func (r *repo) Enable(ctx context.Context, uuid string) *i18np.Error {
	return nil
}

func (r *repo) Delete(ctx context.Context, uuid string) *i18np.Error {
	return nil
}

func (r *repo) Filter(ctx context.Context, filter place.EntityFilter, listConfig list.Config) (*list.Result[*place.Entity], *i18np.Error) {
	return nil, nil
}

func (r *repo) View(ctx context.Context, detail place.I18nDetail) (*place.Entity, *i18np.Error) {
	return nil, nil
}
