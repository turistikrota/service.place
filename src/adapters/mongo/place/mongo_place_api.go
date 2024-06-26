package place

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.place/src/adapters/mongo/place/entity"
	"github.com/turistikrota/service.place/src/domain/place"
	"github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, e *place.Entity) (string, *i18np.Error) {
	p := &entity.MongoPlace{}
	res, err := r.collection.InsertOne(ctx, p.FromEntity(e))
	if err != nil {
		return "", r.factory.Errors.Failed("create")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *repo) Update(ctx context.Context, uuid string, e *place.Entity) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	p := &entity.MongoPlace{}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	update := bson.M{
		"$set": p.FromEntityUpdate(e),
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Disable(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
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
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
		entity.Fields.IsActive: bson.M{
			"$ne": true,
		},
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
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
			entity.Fields.IsActive:  false,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Filter(ctx context.Context, filter place.EntityFilter, listConfig list.Config) (*list.Result[*place.Entity], *i18np.Error) {
	anyFilter := r.filterToBson(filter, false)
	transformer := func(e *entity.MongoPlace) *place.Entity {
		return e.ToListEntity()
	}
	l, err := r.helper.GetListFilterTransform(ctx, anyFilter, transformer, r.sort(r.filterOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, anyFilter)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, r.baseFilter())
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*place.Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) List(ctx context.Context, filter place.EntityFilter, listConfig list.Config) (*list.Result[*place.Entity], *i18np.Error) {
	anyFilter := r.filterToBson(filter, false)
	transformer := func(e *entity.MongoPlace) *place.Entity {
		return e.ToAdminListEntity()
	}
	l, err := r.helper.GetListFilterTransform(ctx, anyFilter, transformer, r.sort(r.listOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, anyFilter)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, r.baseFilter())
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*place.Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) AdminList(ctx context.Context, filter place.EntityFilter, listConfig list.Config) (*list.Result[*place.Entity], *i18np.Error) {
	anyFilter := r.filterToBson(filter, true)
	transformer := func(e *entity.MongoPlace) *place.Entity {
		return e.ToAdminListEntity()
	}
	l, err := r.helper.GetListFilterTransform(ctx, anyFilter, transformer, r.sort(r.listOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, anyFilter)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, r.baseFilter())
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*place.Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) View(ctx context.Context, detail place.I18nDetail) (*place.Entity, *i18np.Error) {
	filter := bson.M{
		entity.TranslationField(detail.Locale, entity.TranslationFields.Slug): detail.Slug,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter, r.viewOptions())
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return e.ToViewEntity(), nil
}

func (r *repo) AdminView(ctx context.Context, uuid string) (*place.Entity, *i18np.Error) {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	e, exist, error := r.helper.GetFilter(ctx, filter)
	if error != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return e.ToAdminViewEntity(), nil
}

func (r *repo) filterToBson(filter place.EntityFilter, admin bool) bson.M {
	list := make([]bson.M, 0)
	if admin {
		list = append(list, r.adminBaseFilter())
	} else {
		list = append(list, r.baseFilter())
	}
	list = r.filterByQuery(list, filter)
	list = r.filterFeatureUUIDs(list, filter)
	list = r.filterByLocation(list, filter)
	list = r.filterIsPayed(list, filter)
	list = r.filterTimeSpent(list, filter)
	list = r.filterReview(list, filter)
	list = r.filterByTypes(list, filter)
	listLen := len(list)
	if listLen == 0 {
		return bson.M{}
	}
	if listLen == 1 {
		return list[0]
	}
	return bson.M{
		"$and": list,
	}
}

func (r *repo) viewOptions() *options.FindOneOptions {
	opts := &options.FindOneOptions{}
	opts.SetProjection(bson.M{
		entity.Fields.UUID:             1,
		entity.Fields.FeatureUUIDs:     1,
		entity.Fields.Images:           1,
		entity.Fields.Translations:     1,
		entity.Fields.AverageTimeSpent: 1,
		entity.Fields.Review:           1,
		entity.Fields.Restorations:     1,
		entity.Fields.IsPayed:          1,
		entity.Fields.UpdatedAt:        1,
		entity.Fields.Type:             1,
		entity.Fields.Coordinates:      1,
	})
	return opts
}

func (r *repo) filterOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		entity.Fields.Images:           1,
		entity.Fields.Translations:     1,
		entity.Fields.AverageTimeSpent: 1,
		entity.Fields.Review:           1,
		entity.Fields.IsPayed:          1,
		entity.Fields.Type:             1,
		entity.Fields.Coordinates:      1,
	}).SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}

func (r *repo) listOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	return opts.SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
}
