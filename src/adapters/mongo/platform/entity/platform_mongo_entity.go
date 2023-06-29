package entity

import (
	"time"

	"api.turistikrota.com/account/src/domain/platform"
)

type MongoPlatform struct {
	UUID         string                                `bson:"uuid"`
	Name         string                                `bson:"name"`
	Slug         string                                `bson:"slug"`
	Regexp       string                                `bson:"regexp"`
	Prefix       string                                `bson:"prefix"`
	Translations map[platform.Locale]MongoTranslations `bson:"translations"`
	IsActive     bool                                  `bson:"is_active"`
	IsDeleted    bool                                  `bson:"is_deleted"`
	UpdatedAt    time.Time                             `bson:"updated_at"`
	CreatedAt    time.Time                             `bson:"created_at"`
}

type MongoTranslations struct {
	Name        string `bson:"name"`
	Placeholder string `bson:"placeholder"`
	Description string `bson:"description"`
}

func (e *MongoPlatform) ToEntity() *platform.Entity {
	return &platform.Entity{
		UUID:         e.UUID,
		Name:         e.Name,
		Slug:         e.Slug,
		Regexp:       e.Regexp,
		Prefix:       e.Prefix,
		Translations: e.ToTranslations(),
		IsActive:     e.IsActive,
		IsDeleted:    e.IsDeleted,
		UpdatedAt:    e.UpdatedAt,
		CreatedAt:    e.CreatedAt,
	}
}

func (e *MongoPlatform) ToTranslations() map[platform.Locale]platform.Translations {
	translations := make(map[platform.Locale]platform.Translations, 0)
	for i, t := range e.Translations {
		translations[i] = t.ToEntity()
	}
	return translations
}

func (e *MongoTranslations) ToEntity() platform.Translations {
	return platform.Translations{
		Name:        e.Name,
		Placeholder: e.Placeholder,
		Description: e.Description,
	}
}

func (e *MongoPlatform) FromEntity(entity *platform.Entity) *MongoPlatform {
	return &MongoPlatform{
		UUID:         entity.UUID,
		Name:         entity.Name,
		Slug:         entity.Slug,
		Regexp:       entity.Regexp,
		Prefix:       entity.Prefix,
		Translations: e.FromTranslations(entity.Translations),
		IsActive:     entity.IsActive,
		IsDeleted:    entity.IsDeleted,
		UpdatedAt:    entity.UpdatedAt,
		CreatedAt:    entity.CreatedAt,
	}
}

func (e *MongoPlatform) FromTranslations(translations map[platform.Locale]platform.Translations) map[platform.Locale]MongoTranslations {
	mongoTranslations := make(map[platform.Locale]MongoTranslations, 0)
	for i, t := range translations {
		mongoTranslations[i] = e.FromTranslation(t)
	}
	return mongoTranslations
}

func (e *MongoPlatform) FromTranslation(translation platform.Translations) MongoTranslations {
	return MongoTranslations{
		Name:        translation.Name,
		Placeholder: translation.Placeholder,
		Description: translation.Description,
	}
}
