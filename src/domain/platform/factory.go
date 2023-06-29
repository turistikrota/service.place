package platform

import (
	"regexp"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/ssibrahimbas/slug"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newPlatformErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewPlatform(name string, regexp string, prefix string) *Entity {
	t := time.Now()
	return &Entity{
		Name:         name,
		Slug:         slug.New(name),
		Regexp:       regexp,
		Prefix:       prefix,
		IsActive:     false,
		IsDeleted:    false,
		Translations: map[Locale]Translations{},
		UpdatedAt:    t,
		CreatedAt:    t,
	}
}

func (f Factory) Validate(platform *Entity) *i18np.Error {
	if platform.Name == "" {
		f.Errors.NameRequired()
	}
	if platform.Regexp == "" {
		f.Errors.RegexpRequired()
	}
	if platform.Prefix == "" {
		f.Errors.PrefixRequired()
	}
	return nil
}

func (f Factory) ValidatePlatformValue(platform *Entity, value string) *i18np.Error {
	if platform.Regexp != "" {
		if matched, _ := regexp.MatchString(platform.Regexp, f.FixPlatformValue(platform, value)); !matched {
			return f.Errors.InvalidValue()
		}
	}
	return nil
}

func (f Factory) FixPlatformValue(platform *Entity, value string) string {
	return platform.Prefix + value
}
