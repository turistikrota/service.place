package mongo

import (
	mongo_feature "api.turistikrota.com/place/src/adapters/mongo/feature"
	"api.turistikrota.com/place/src/domain/feature"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewFeature(factory feature.Factory, collection *mongo.Collection) feature.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewFeature(factory feature.Factory, collection *mongo.Collection) feature.Repository {
	return mongo_feature.New(factory, collection)
}
