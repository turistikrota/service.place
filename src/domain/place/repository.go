package place

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
)

type I18nDetail struct {
	Locale string
	Slug   string
}

type Repository interface {
	Create(ctx context.Context, place *Entity) *i18np.Error
	Update(ctx context.Context, uuid string, place *Entity) *i18np.Error
	Disable(ctx context.Context, uuid string) *i18np.Error
	Delete(ctx context.Context, uuid string) *i18np.Error
	Enable(ctx context.Context, uuid string) *i18np.Error
	Filter(ctx context.Context, filter EntityFilter, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	View(ctx context.Context, detail I18nDetail) (*Entity, *i18np.Error)
	List(ctx context.Context, filter EntityFilter, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}
