package entity

import (
	"time"

	"github.com/turistikrota/service.place/src/domain/feature"
)

type MongoFeature struct {
	UUID         string                       `bson:"_id,omitempty"`
	Icon         string                       `bson:"icon"`
	Translations map[string]MongoTranslations `bson:"translations"`
	IsActive     bool                         `bson:"is_active"`
	IsDeleted    bool                         `bson:"is_deleted"`
	UpdatedAt    time.Time                    `bson:"updated_at"`
	CreatedAt    time.Time                    `bson:"created_at"`
}

type MongoTranslations struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
}

func (e *MongoFeature) FromEntity(entity *feature.Entity) *MongoFeature {
	e.Icon = entity.Icon
	e.Translations = e.fromTranslations(entity.Translations)
	e.IsActive = entity.IsActive
	e.IsDeleted = entity.IsDeleted
	t := time.Now()
	entity.UpdatedAt = t
	entity.CreatedAt = t
	return e
}

func (e *MongoFeature) FromEntityUpdate(entity *feature.Entity) *MongoFeature {
	e.Icon = entity.Icon
	e.Translations = e.fromTranslations(entity.Translations)
	e.IsActive = entity.IsActive
	e.IsDeleted = entity.IsDeleted
	entity.UpdatedAt = time.Now()
	return e
}

func (e *MongoFeature) ToEntity() *feature.Entity {
	return &feature.Entity{
		UUID:         e.UUID,
		Icon:         e.Icon,
		Translations: e.toTranslations(),
		IsActive:     e.IsActive,
		IsDeleted:    e.IsDeleted,
		UpdatedAt:    e.UpdatedAt,
		CreatedAt:    e.CreatedAt,
	}
}

func (e *MongoFeature) ToListEntity() *feature.Entity {
	return &feature.Entity{
		UUID:         e.UUID,
		Icon:         e.Icon,
		Translations: e.toTranslations(),
	}
}

func (e *MongoFeature) fromTranslations(translations map[feature.Locale]feature.Translations) map[string]MongoTranslations {
	mongoTranslations := make(map[string]MongoTranslations)
	for locale, translation := range translations {
		mongoTranslations[locale.String()] = MongoTranslations{
			Title:       translation.Title,
			Description: translation.Description,
		}
	}
	return mongoTranslations
}

func (e *MongoFeature) toTranslations() map[feature.Locale]feature.Translations {
	translations := make(map[feature.Locale]feature.Translations)
	for locale, translation := range e.Translations {
		translations[feature.Locale(locale)] = feature.Translations{
			Title:       translation.Title,
			Description: translation.Description,
		}
	}
	return translations
}
