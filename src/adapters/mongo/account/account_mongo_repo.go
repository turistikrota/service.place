package account

import (
	"api.turistikrota.com/account/src/adapters/mongo/account/entity"
	"api.turistikrota.com/account/src/domain/account"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    account.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoAccount, *account.Entity]
}

func New(accountFactory account.Factory, collection *mongo.Collection) account.Repository {
	validate(accountFactory, collection)
	return &repo{
		factory:    accountFactory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoAccount, *account.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoAccount {
	return &entity.MongoAccount{}
}

func validate(accountFactory account.Factory, collection *mongo.Collection) {
	if accountFactory.IsZero() {
		panic("accountFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
