package platform

import (
	"api.turistikrota.com/place/src/adapters/mongo/platform/entity"
	"api.turistikrota.com/place/src/domain/platform"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    platform.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoPlatform, *platform.Entity]
}

func New(platformFactory platform.Factory, collection *mongo.Collection) platform.Repository {
	validate(platformFactory, collection)
	return &repo{
		factory:    platformFactory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoPlatform, *platform.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoPlatform {
	return &entity.MongoPlatform{}
}

func validate(platformFactory platform.Factory, collection *mongo.Collection) {
	if platformFactory.IsZero() {
		panic("platformFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
