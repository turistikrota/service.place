package place

import (
	"api.turistikrota.com/place/src/adapters/mongo/place/entity"
	"api.turistikrota.com/place/src/domain/place"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    place.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoPlace, *place.Entity]
}

func New(factory place.Factory, collection *mongo.Collection) place.Repository {
	validate(factory, collection)
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoPlace, *place.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoPlace {
	return &entity.MongoPlace{}
}

func validate(factory place.Factory, collection *mongo.Collection) {
	if factory.IsZero() {
		panic("categoryFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
