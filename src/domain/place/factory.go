package place

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newPlaceErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
