package account

import (
	"context"
	"time"

	"api.turistikrota.com/place/src/adapters/mongo/account/entity"
	"api.turistikrota.com/place/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, account *account.Entity) (*account.Entity, *i18np.Error) {
	n := &entity.MongoAccount{}
	res, err := r.collection.InsertOne(ctx, n.FromEntity(account))
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	account.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return account, nil
}

func (r *repo) ProfileView(ctx context.Context, u account.UserUnique) (*account.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserName: u.Name,
		entity.Fields.IsActive: true,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	opts := options.FindOne().SetProjection(bson.M{
		entity.Fields.UserName:    1,
		entity.Fields.FullName:    1,
		entity.Fields.Description: 1,
		entity.Fields.Social:      1,
		entity.Fields.IsVerified:  1,
		entity.Fields.CreatedAt:   1,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToEntity(), nil
}

func (r *repo) Get(ctx context.Context, u account.UserUnique) (*account.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToEntity(), nil
}

func (r *repo) Exist(ctx context.Context, u account.UserUnique) (bool, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return o != nil, nil
}

func (r *repo) SocialAdd(ctx context.Context, u account.UserUnique, social *account.EntitySocial) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
	}
	setter := bson.M{
		"$addToSet": bson.M{
			entity.Fields.Social: bson.M{
				entity.SocialFields.Platform:   social.Platform,
				entity.SocialFields.Value:      social.Value,
				entity.SocialFields.FixedValue: social.FixedValue,
			},
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) SocialUpdate(ctx context.Context, u account.UserUnique, social *account.EntitySocial) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID:                           u.UUID,
		entity.Fields.UserName:                           u.Name,
		entity.SocialField(entity.SocialFields.Platform): social.Platform,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.SocialFieldInArray(entity.SocialFields.Value):      social.Value,
			entity.SocialFieldInArray(entity.SocialFields.FixedValue): social.FixedValue,
			entity.Fields.UpdatedAt:                                   time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) SocialRemove(ctx context.Context, u account.UserUnique, platform string) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID:                           u.UUID,
		entity.Fields.UserName:                           u.Name,
		entity.SocialField(entity.SocialFields.Platform): platform,
	}
	setter := bson.M{
		"$pull": bson.M{
			entity.Fields.Social: bson.M{
				entity.SocialFields.Platform: platform,
			},
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Update(ctx context.Context, u account.UserUnique, account *account.Entity) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.UserName:    account.UserName,
			entity.Fields.FullName:    account.FullName,
			entity.Fields.Description: account.Description,
			entity.Fields.BirthDate:   account.BirthDate,
			entity.Fields.UpdatedAt:   account.UpdatedAt,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.IsActive: true,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive:  false,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.IsActive: false,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive:  true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListMy(ctx context.Context, userUUID string) ([]*account.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserUUID: userUUID,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	transformer := func(acc *entity.MongoAccount) *account.Entity {
		return acc.ToEntity()
	}
	return r.helper.GetListFilterTransform(ctx, filter, transformer)
}
