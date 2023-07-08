package feature

import (
	"github.com/turistikrota/service.place/src/adapters/mongo/feature/entity"
	"github.com/turistikrota/service.place/src/domain/feature"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    feature.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoFeature, *feature.Entity]
}

func New(factory feature.Factory, collection *mongo.Collection) feature.Repository {
	validate(factory, collection)
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoFeature, *feature.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoFeature {
	return &entity.MongoFeature{}
}

func validate(factory feature.Factory, collection *mongo.Collection) {
	if factory.IsZero() {
		panic("categoryFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
