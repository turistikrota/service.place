package entity

import (
	"time"

	"api.turistikrota.com/place/src/domain/place"
)

type (
	MongoPlace struct {
		UUID             string                      `bson:"_id"`
		FeatureUUIDs     []string                    `bson:"feature_uuids"`
		Images           []MongoImage                `bson:"images"`
		Translations     map[string]MongoTranslation `bson:"translations"`
		AverageTimeSpent MongoTimeSpent              `bson:"average_time_spent"`
		Review           MongoReview                 `bson:"review"`
		Coordinates      []float64                   `bson:"coordinates"`
		IsActive         bool                        `bson:"is_active"`
		IsDeleted        bool                        `bson:"is_deleted"`
		IsPayed          bool                        `bson:"is_payed"`
		UpdatedAt        time.Time                   `bson:"updated_at"`
		CreatedAt        time.Time                   `bson:"created_at"`
	}
	MongoTranslation struct {
		Title       string   `bson:"title"`
		Description string   `bson:"description"`
		Slug        string   `bson:"slug"`
		MarkdownURL string   `bson:"markdown_url"`
		Seo         MongoSeo `bson:"seo"`
	}
	MongoSeo struct {
		Title       string `bson:"title"`
		Description string `bson:"description"`
		Keywords    string `bson:"keywords"`
	}
	MongoImage struct {
		Url   string `bson:"url"`
		Order int16  `bson:"order"`
	}
	MongoTimeSpent struct {
		Min int16 `bson:"min"`
		Max int16 `bson:"max"`
	}
	MongoReview struct {
		Total        int16   `bson:"total"`
		AveragePoint float32 `bson:"average_point"`
	}
)

func (e *MongoPlace) FromEntity(entity *place.Entity) *MongoPlace {
	e.FeatureUUIDs = entity.FeatureUUIDs
	e.Images = e.fromImages(entity.Images)
	e.Translations = e.fromTranslations(entity.Translations)
	e.AverageTimeSpent = e.fromTimeSpent(entity.AverageTimeSpent)
	e.Review = e.fromReview(entity.Review)
	e.IsActive = entity.IsActive
	e.Coordinates = entity.Coordinates
	e.IsDeleted = entity.IsDeleted
	e.IsPayed = entity.IsPayed
	t := time.Now()
	e.UpdatedAt = t
	e.CreatedAt = t
	return e
}

func (e *MongoPlace) FromEntityUpdate(entity *place.Entity) *MongoPlace {
	e.FeatureUUIDs = entity.FeatureUUIDs
	e.Images = e.fromImages(entity.Images)
	e.Translations = e.fromTranslations(entity.Translations)
	e.AverageTimeSpent = e.fromTimeSpent(entity.AverageTimeSpent)
	e.Review = e.fromReview(entity.Review)
	e.IsActive = entity.IsActive
	e.IsDeleted = entity.IsDeleted
	e.Coordinates = entity.Coordinates
	e.IsPayed = entity.IsPayed
	e.UpdatedAt = time.Now()
	return e
}

func (e *MongoPlace) ToListEntity() *place.Entity {
	return &place.Entity{
		FeatureUUIDs:     e.FeatureUUIDs,
		Images:           e.toImages(),
		Translations:     e.toTranslations(),
		AverageTimeSpent: e.toTimeSpent(),
		Coordinates:      e.Coordinates,
		Review:           e.toReview(),
		IsPayed:          e.IsPayed,
	}
}

func (e *MongoPlace) ToAdminListEntity() *place.Entity {
	return &place.Entity{
		UUID:             e.UUID,
		FeatureUUIDs:     e.FeatureUUIDs,
		Images:           e.toImages(),
		Translations:     e.toTranslations(),
		AverageTimeSpent: e.toTimeSpent(),
		Review:           e.toReview(),
		Coordinates:      e.Coordinates,
		IsPayed:          e.IsPayed,
		IsActive:         e.IsActive,
		IsDeleted:        e.IsDeleted,
	}
}

func (e *MongoPlace) ToViewEntity() *place.Entity {
	return &place.Entity{
		FeatureUUIDs:     e.FeatureUUIDs,
		Images:           e.toImages(),
		Translations:     e.toTranslations(),
		AverageTimeSpent: e.toTimeSpent(),
		Review:           e.toReview(),
		Coordinates:      e.Coordinates,
		IsPayed:          e.IsPayed,
		CreatedAt:        e.CreatedAt,
	}
}

func (e *MongoPlace) ToAdminViewEntity() *place.Entity {
	return &place.Entity{
		UUID:             e.UUID,
		FeatureUUIDs:     e.FeatureUUIDs,
		Images:           e.toImages(),
		Translations:     e.toTranslations(),
		AverageTimeSpent: e.toTimeSpent(),
		Review:           e.toReview(),
		IsPayed:          e.IsPayed,
		Coordinates:      e.Coordinates,
		IsActive:         e.IsActive,
		IsDeleted:        e.IsDeleted,
		UpdatedAt:        e.UpdatedAt,
		CreatedAt:        e.CreatedAt,
	}
}

func (e *MongoPlace) fromImages(images []place.Image) []MongoImage {
	mongoImages := make([]MongoImage, 0)
	for _, image := range images {
		mongoImages = append(mongoImages, MongoImage{
			Url:   image.Url,
			Order: image.Order,
		})
	}
	return mongoImages
}

func (e *MongoPlace) toImages() []place.Image {
	images := make([]place.Image, 0)
	for _, mongoImage := range e.Images {
		images = append(images, place.Image{
			Url:   mongoImage.Url,
			Order: mongoImage.Order,
		})
	}
	return images
}

func (e *MongoPlace) fromTranslations(translations map[place.Locale]place.Translations) map[string]MongoTranslation {
	mongoTranslations := make(map[string]MongoTranslation)
	for key, translation := range translations {
		mongoTranslations[key.String()] = MongoTranslation{
			Title:       translation.Title,
			Description: translation.Description,
			Slug:        translation.Slug,
			MarkdownURL: translation.MarkdownURL,
			Seo: MongoSeo{
				Title:       translation.Seo.Title,
				Description: translation.Seo.Description,
				Keywords:    translation.Seo.Keywords,
			},
		}
	}
	return mongoTranslations
}

func (e *MongoPlace) toTranslations() map[place.Locale]place.Translations {
	translations := make(map[place.Locale]place.Translations)
	for key, mongoTranslation := range e.Translations {
		translations[place.Locale(key)] = place.Translations{
			Title:       mongoTranslation.Title,
			Description: mongoTranslation.Description,
			Slug:        mongoTranslation.Slug,
			MarkdownURL: mongoTranslation.MarkdownURL,
			Seo: place.Seo{
				Title:       mongoTranslation.Seo.Title,
				Description: mongoTranslation.Seo.Description,
				Keywords:    mongoTranslation.Seo.Keywords,
			},
		}
	}
	return translations
}

func (e *MongoPlace) fromTimeSpent(timeSpent place.TimeSpent) MongoTimeSpent {
	return MongoTimeSpent{
		Min: timeSpent.Min,
		Max: timeSpent.Max,
	}
}

func (e *MongoPlace) toTimeSpent() place.TimeSpent {
	return place.TimeSpent{
		Min: e.AverageTimeSpent.Min,
		Max: e.AverageTimeSpent.Max,
	}
}

func (e *MongoPlace) fromReview(review place.Review) MongoReview {
	return MongoReview{
		Total:        review.Total,
		AveragePoint: review.AveragePoint,
	}
}

func (e *MongoPlace) toReview() place.Review {
	return place.Review{
		Total:        e.Review.Total,
		AveragePoint: e.Review.AveragePoint,
	}
}
