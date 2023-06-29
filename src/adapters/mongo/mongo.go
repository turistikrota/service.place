package mongo

import (
	mongo_account "api.turistikrota.com/place/src/adapters/mongo/account"
	mongo_platform "api.turistikrota.com/place/src/adapters/mongo/platform"
	"api.turistikrota.com/place/src/domain/account"
	"api.turistikrota.com/place/src/domain/platform"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewAccount(accountFactory account.Factory, collection *mongo.Collection) account.Repository
	NewPlatform(platformFactory platform.Factory, collection *mongo.Collection) platform.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewAccount(accountFactory account.Factory, collection *mongo.Collection) account.Repository {
	return mongo_account.New(accountFactory, collection)
}

func (m *mongodb) NewPlatform(platformFactory platform.Factory, collection *mongo.Collection) platform.Repository {
	return mongo_platform.New(platformFactory, collection)
}
