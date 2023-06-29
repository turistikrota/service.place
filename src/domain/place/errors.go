package place

import "github.com/mixarchitecture/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID() *i18np.Error
	NotFound() *i18np.Error
}

type placeErrors struct{}

func newPlaceErrors() Errors {
	return &placeErrors{}
}

func (e *placeErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{
		"Action": action,
	})
}

func (e *placeErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidUUID)
}

func (e *placeErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}