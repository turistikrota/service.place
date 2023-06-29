package feature

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newFeatureErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
