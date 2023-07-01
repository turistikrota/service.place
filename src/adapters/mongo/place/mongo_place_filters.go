package place

import (
	"api.turistikrota.com/place/src/adapters/mongo/place/entity"
	"api.turistikrota.com/place/src/domain/place"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) baseFilter() bson.M {
	return bson.M{
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
}

func (r *repo) filterByTypes(list []bson.M, filter place.EntityFilter) []bson.M {
	if len(filter.Types) > 0 {
		list = append(list, bson.M{
			entity.Fields.Type: bson.M{
				"$in": filter.Types,
			},
		})
	}
	return list
}

func (r *repo) filterByQuery(list []bson.M, filter place.EntityFilter) []bson.M {
	if filter.Query != "" {
		list = append(list, bson.M{
			"$or": []bson.M{
				{
					entity.TranslationField(filter.Locale, entity.TranslationFields.Title): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
				{
					entity.TranslationField(filter.Locale, entity.TranslationFields.Description): bson.M{
						"$regex":   filter.Query,
						"$options": "i",
					},
				},
			},
		})
	}
	return list
}

func (r *repo) filterByLocation(list []bson.M, filter place.EntityFilter) []bson.M {
	if filter.Coordinates != nil && len(filter.Coordinates) == 2 {
		distance := filter.Distance
		if distance == 0 {
			distance = 100000 // 100 km
		}
		list = append(list, bson.M{
			entity.Fields.Coordinates: bson.M{
				"$near": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": filter.Coordinates,
					},
					"$maxDistance": distance,
				},
			},
		})
	}
	return list
}

func (r *repo) filterIsPayed(list []bson.M, filter place.EntityFilter) []bson.M {
	if filter.IsPayed != nil {
		list = append(list, bson.M{
			entity.Fields.IsPayed: *filter.IsPayed,
		})
	}
	return list
}

func (r *repo) filterReview(list []bson.M, filter place.EntityFilter) []bson.M {
	reviewFilter := make([]bson.M, 0)
	if filter.MinReview != nil && filter.MaxReview != nil {
		reviewFilter = append(reviewFilter, bson.M{
			entity.ReviewField(entity.ReviewFields.Total): bson.M{
				"$gte": *filter.MinReview,
				"$lte": *filter.MaxReview,
			},
		})
	}
	if filter.MinAveragePoint != nil && filter.MaxAveragePoint != nil {
		reviewFilter = append(reviewFilter, bson.M{
			entity.ReviewField(entity.ReviewFields.AveragePoint): bson.M{
				"$gte": *filter.MinAveragePoint,
				"$lte": *filter.MaxAveragePoint,
			},
		})
	}
	if len(reviewFilter) > 0 {
		list = append(list, reviewFilter...)
	}
	return list
}

func (r *repo) filterTimeSpent(list []bson.M, filter place.EntityFilter) []bson.M {
	if filter.AverageTimeSpent != nil {
		list = append(list, bson.M{
			entity.TimeSpentField(entity.TimeSpentFields.Min): bson.M{
				"$lte": *filter.AverageTimeSpent,
			},
			entity.TimeSpentField(entity.TimeSpentFields.Max): bson.M{
				"$gte": *filter.AverageTimeSpent,
			},
		})
	}
	return list
}

func (r *repo) filterFeatureUUIDs(list []bson.M, filter place.EntityFilter) []bson.M {
	if len(filter.FeatureUUIDs) > 0 {
		list = append(list, bson.M{
			entity.Fields.FeatureUUIDs: bson.M{
				"$in": filter.FeatureUUIDs,
			},
		})
	}
	return list
}
