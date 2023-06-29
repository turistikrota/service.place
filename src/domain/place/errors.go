package place

type Errors interface{}

type placeErrors struct{}

func newPlaceErrors() Errors {
	return &placeErrors{}
}
