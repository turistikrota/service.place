package feature

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Repository interface {
	Create(ctx context.Context, feature *Entity) *i18np.Error
	Update(ctx context.Context, uuid string, feature *Entity) *i18np.Error
	Disable(ctx context.Context, uuid string) *i18np.Error
	Delete(ctx context.Context, uuid string) *i18np.Error
	Enable(ctx context.Context, uuid string) *i18np.Error
	ListAll(ctx context.Context) ([]*Entity, *i18np.Error)
	GetByUUIDs(ctx context.Context, uuids []string) ([]*Entity, *i18np.Error)
}
