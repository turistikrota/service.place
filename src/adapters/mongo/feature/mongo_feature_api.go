package feature

import (
	"context"

	"api.turistikrota.com/place/src/domain/feature"
	"github.com/mixarchitecture/i18np"
)

func (r *repo) Create(ctx context.Context, entity *feature.Entity) *i18np.Error {
	return nil
}

func (r *repo) Update(ctx context.Context, uuid string, entity *feature.Entity) *i18np.Error {
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

func (r *repo) ListAll(ctx context.Context) ([]*feature.Entity, *i18np.Error) {
	return nil, nil
}

func (r *repo) GetByUUIDs(ctx context.Context, uuids []string) ([]*feature.Entity, *i18np.Error) {
	return nil, nil
}
