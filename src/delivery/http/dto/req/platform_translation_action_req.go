package req

import (
	"api.turistikrota.com/account/src/app/command"
	"api.turistikrota.com/account/src/domain/platform"
)

type PlatformTranslationActionRequest struct {
	Slug              string
	Locale            platform.Locale
	LocaleName        string `json:"locale_name,omitempty" validate:"omitempty"`
	LocaleDescription string `json:"description" validate:"required"`
	LocalePlaceholder string `json:"placeholder" validate:"required"`
}

func (r *PlatformTranslationActionRequest) LoadDetail(detail *PlatformTranslationRequest) {
	r.Slug = detail.Slug
	r.Locale = platform.Locale(detail.Locale)
}

func (r *PlatformTranslationActionRequest) ToCreateCommand(adminUUID string) command.PlatformTranslationCreateCommand {
	return command.PlatformTranslationCreateCommand{
		AdminUUID:         adminUUID,
		Slug:              r.Slug,
		Locale:            r.Locale,
		LocaleName:        r.LocaleName,
		LocaleDescription: r.LocaleDescription,
		LocalePlaceholder: r.LocalePlaceholder,
	}
}

func (r *PlatformTranslationActionRequest) ToUpdateCommand(adminUUID string) command.PlatformTranslationUpdateCommand {
	return command.PlatformTranslationUpdateCommand{
		AdminUUID:         adminUUID,
		Slug:              r.Slug,
		Locale:            r.Locale,
		LocaleName:        r.LocaleName,
		LocaleDescription: r.LocaleDescription,
		LocalePlaceholder: r.LocalePlaceholder,
	}
}
