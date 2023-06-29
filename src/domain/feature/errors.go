package feature

import "github.com/mixarchitecture/i18np"

type Errors interface {
	Failed(string) *i18np.Error
	InvalidUUID(string) *i18np.Error
}

type featureErrors struct{}

func newFeatureErrors() Errors {
	return &featureErrors{}
}

func (f *featureErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{
		"Action": action,
	})
}

func (f *featureErrors) InvalidUUID(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidUUID, i18np.P{
		"Action": action,
	})
}
