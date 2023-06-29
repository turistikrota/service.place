package platform

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Repository interface {
	GetBySlug(ctx context.Context, slug string) (*Entity, *i18np.Error)
	Create(ctx context.Context, platform *Entity) *i18np.Error
	Update(ctx context.Context, slug string, platform *Entity) *i18np.Error
	Disable(ctx context.Context, slug string) *i18np.Error
	Delete(ctx context.Context, slug string) *i18np.Error
	Enable(ctx context.Context, slug string) *i18np.Error
	ListAll(ctx context.Context) ([]*Entity, *i18np.Error)
	TranslationCreate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error
	TranslationUpdate(ctx context.Context, platform string, locale Locale, translations Translations) *i18np.Error
	TranslationDelete(ctx context.Context, platform string, locale Locale) *i18np.Error
}
