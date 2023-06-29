package platform

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	NotFound(name string) *i18np.Error
	NameRequired() *i18np.Error
	RegexpRequired() *i18np.Error
	PrefixRequired() *i18np.Error
	InvalidValue() *i18np.Error
	TranslationAlreadyExists(locale string) *i18np.Error
	TranslationNotExists(locale string) *i18np.Error
	Failed(action string) *i18np.Error
	AlreadyExists(name string) *i18np.Error
}

type platformErrors struct{}

func newPlatformErrors() Errors {
	return &platformErrors{}
}

func (e *platformErrors) NotFound(name string) *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformNotFound, i18np.P{"Name": name})
}

func (e *platformErrors) NameRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformNameRequired)
}

func (e *platformErrors) RegexpRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformRegexpRequired)
}

func (e *platformErrors) PrefixRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformPrefixRequired)
}

func (e *platformErrors) InvalidValue() *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformInvalidValue)
}

func (e *platformErrors) TranslationAlreadyExists(locale string) *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformTranslationAlreadyExists, i18np.P{"Locale": locale})
}

func (e *platformErrors) TranslationNotExists(locale string) *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformTranslationNotExists, i18np.P{"Locale": locale})
}

func (e *platformErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformFailed, i18np.P{"Action": action})
}

func (e *platformErrors) AlreadyExists(name string) *i18np.Error {
	return i18np.NewError(I18nMessages.PlatformAlreadyExists, i18np.P{"Name": name})
}
