package account

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type UserUnique struct {
	UUID string
	Name string
}

type Repository interface {
	Create(ctx context.Context, account *Entity) (*Entity, *i18np.Error)
	ProfileView(ctx context.Context, u UserUnique) (*Entity, *i18np.Error)
	Get(ctx context.Context, u UserUnique) (*Entity, *i18np.Error)
	Exist(ctx context.Context, u UserUnique) (bool, *i18np.Error)
	SocialAdd(ctx context.Context, u UserUnique, social *EntitySocial) *i18np.Error
	SocialUpdate(ctx context.Context, u UserUnique, social *EntitySocial) *i18np.Error
	SocialRemove(ctx context.Context, u UserUnique, platform string) *i18np.Error
	Update(ctx context.Context, u UserUnique, account *Entity) *i18np.Error
	Disable(ctx context.Context, u UserUnique) *i18np.Error
	Enable(ctx context.Context, u UserUnique) *i18np.Error
	Delete(ctx context.Context, u UserUnique) *i18np.Error
	ListMy(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error)
}
