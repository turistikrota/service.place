package feature

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.place/src/adapters/mongo/feature/entity"
	"github.com/turistikrota/service.place/src/domain/feature"
	"github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, e *feature.Entity) *i18np.Error {
	f := &entity.MongoFeature{}
	_, err := r.collection.InsertOne(ctx, f.FromEntity(e))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, uuid string, e *feature.Entity) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID("update")
	}
	c := &entity.MongoFeature{}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	update := bson.M{
		"$set": c.FromEntityUpdate(e),
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Disable(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID("disable")
	}
	filter := bson.M{
		entity.Fields.UUID:     id,
		entity.Fields.IsActive: true,
	}
	update := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive:  false,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Enable(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID("enable")
	}
	filter := bson.M{
		entity.Fields.UUID:     id,
		entity.Fields.IsActive: false,
	}
	update := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive:  true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID("enable")
	}
	filter := bson.M{
		entity.Fields.UUID:      id,
		entity.Fields.IsDeleted: false,
	}
	update := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) ListAll(ctx context.Context) ([]*feature.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
	transformer := func(e *entity.MongoFeature) *feature.Entity {
		return e.ToListEntity()
	}
	return r.helper.GetListFilterTransform(ctx, filter, transformer, r.listOptions())
}

func (r *repo) GetByUUIDs(ctx context.Context, uuids []string) ([]*feature.Entity, *i18np.Error) {
	ids, err := mongo.TransformIds(uuids)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID("find by uuids")
	}
	filter := bson.M{
		entity.Fields.UUID: bson.M{
			"$in": ids,
		},
		entity.Fields.IsActive: true,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	transformer := func(e *entity.MongoFeature) *feature.Entity {
		return e.ToListEntity()
	}
	return r.helper.GetListFilterTransform(ctx, filter, transformer, r.listOptions())
}

func (r *repo) listOptions() *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		entity.Fields.UUID:         1,
		entity.Fields.Icon:         1,
		entity.Fields.Translations: 1,
	})
	return opts
}

func (r *repo) AdminListAll(ctx context.Context) ([]*feature.Entity, *i18np.Error) {
	filter := bson.M{}
	transformer := func(e *entity.MongoFeature) *feature.Entity {
		return e.ToEntity()
	}
	return r.helper.GetListFilterTransform(ctx, filter, transformer)
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*feature.Entity, *i18np.Error) {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID("find by uuid")
	}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	e, exist, error := r.helper.GetFilter(ctx, filter)
	if error != nil {
		return nil, r.factory.Errors.Failed("find by uuid")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return e.ToEntity(), nil
}
