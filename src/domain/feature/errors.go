package feature

type Errors interface{}

type featureErrors struct{}

func newFeatureErrors() Errors {
	return &featureErrors{}
}
