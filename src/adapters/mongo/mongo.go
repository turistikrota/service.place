package mongo

import (
	mongo_feature "github.com/turistikrota/service.place/src/adapters/mongo/feature"
	mongo_place "github.com/turistikrota/service.place/src/adapters/mongo/place"
	"github.com/turistikrota/service.place/src/domain/feature"
	"github.com/turistikrota/service.place/src/domain/place"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewFeature(factory feature.Factory, collection *mongo.Collection) feature.Repository
	NewPlace(factory place.Factory, collection *mongo.Collection) place.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewFeature(factory feature.Factory, collection *mongo.Collection) feature.Repository {
	return mongo_feature.New(factory, collection)
}

func (m *mongodb) NewPlace(factory place.Factory, collection *mongo.Collection) place.Repository {
	return mongo_place.New(factory, collection)
}
