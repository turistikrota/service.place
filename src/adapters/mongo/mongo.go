package mongo

type Mongo interface{}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}
