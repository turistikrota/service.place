package platform

import (
	"context"
	"time"

	"api.turistikrota.com/place/src/adapters/mongo/platform/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"api.turistikrota.com/place/src/domain/platform"
	"github.com/mixarchitecture/i18np"
)

func (r *repo) Create(ctx context.Context, platform *platform.Entity) *i18np.Error {
	n := &entity.MongoPlatform{}
	res, err := r.collection.InsertOne(ctx, n.FromEntity(platform))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	platform.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *repo) GetBySlug(ctx context.Context, slug string) (*platform.Entity, *i18np.Error) {
	filter := bson.M{entity.Fields.Slug: slug}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound("platform")
	}
	return o.ToEntity(), nil
}

func (r *repo) Update(ctx context.Context, slug string, platform *platform.Entity) *i18np.Error {
	filter := bson.M{entity.Fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		entity.Fields.Name:      platform.Name,
		entity.Fields.Slug:      platform.Slug,
		entity.Fields.Regexp:    platform.Regexp,
		entity.Fields.Prefix:    platform.Prefix,
		entity.Fields.UpdatedAt: platform.UpdatedAt,
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{entity.Fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		entity.Fields.IsActive:  false,
		entity.Fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{entity.Fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		entity.Fields.IsActive:  true,
		entity.Fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, slug string) *i18np.Error {
	filter := bson.M{entity.Fields.Slug: slug}
	setter := bson.M{"$set": bson.M{
		entity.Fields.IsDeleted: true,
		entity.Fields.UpdatedAt: time.Now(),
	}}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListAll(ctx context.Context) ([]*platform.Entity, *i18np.Error) {
	filter := bson.M{entity.Fields.IsDeleted: bson.M{
		"$ne": true,
	}, entity.Fields.IsActive: true}
	return r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoPlatform) *platform.Entity {
		return o.ToEntity()
	})
}

func (r *repo) TranslationCreate(ctx context.Context, platform string, locale platform.Locale, translations platform.Translations) *i18np.Error {
	filter := bson.M{
		entity.Fields.Slug: platform,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Translation(locale.String()): bson.M{
				entity.TranslationsFields.Name:        translations.Name,
				entity.TranslationsFields.Description: translations.Description,
				entity.TranslationsFields.Placeholder: translations.Placeholder,
			},
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) TranslationUpdate(ctx context.Context, platform string, locale platform.Locale, translations platform.Translations) *i18np.Error {
	filter := bson.M{
		entity.Fields.Slug: platform,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Translation(locale.String()): bson.M{
				entity.TranslationsFields.Name:        translations.Name,
				entity.TranslationsFields.Description: translations.Description,
				entity.TranslationsFields.Placeholder: translations.Placeholder,
			},
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) TranslationDelete(ctx context.Context, platform string, locale platform.Locale) *i18np.Error {
	filter := bson.M{entity.Fields.Slug: platform}
	setter := bson.M{
		"$unset": bson.M{
			entity.Translation(locale.String()): "",
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}
