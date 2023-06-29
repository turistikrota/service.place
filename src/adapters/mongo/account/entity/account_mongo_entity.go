package entity

import (
	"time"

	"api.turistikrota.com/account/src/domain/account"
)

type MongoAccount struct {
	UUID          string               `bson:"_id,omitempty"`
	UserUUID      string               `bson:"user_uuid"`
	UserName      string               `bson:"user_name"`
	FullName      string               `bson:"full_name"`
	Description   string               `bson:"description"`
	Social        []MongoAccountSocial `bson:"social"`
	IsActive      bool                 `bson:"is_active"`
	CompletedRate int                  `bson:"completed_rate"`
	IsDeleted     bool                 `bson:"is_deleted"`
	IsVerified    bool                 `bson:"is_verified"`
	BirthDate     *time.Time           `bson:"birth_date"`
	CreatedAt     *time.Time           `bson:"created_at"`
	UpdatedAt     *time.Time           `bson:"updated_at"`
}

type MongoAccountSocial struct {
	Platform   string `bson:"platform"`
	Value      string `bson:"value"`
	FixedValue string `bson:"fixed_value"`
}

func (e *MongoAccount) ToEntity() *account.Entity {
	return &account.Entity{
		UUID:          e.UUID,
		UserUUID:      e.UserUUID,
		UserName:      e.UserName,
		FullName:      e.FullName,
		Description:   e.Description,
		Social:        e.ToSocial(),
		IsActive:      e.IsActive,
		CompletedRate: e.CompletedRate,
		IsDeleted:     e.IsDeleted,
		IsVerified:    e.IsVerified,
		BirthDate:     e.BirthDate,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
}

func (e *MongoAccount) ToSocial() []account.EntitySocial {
	social := make([]account.EntitySocial, len(e.Social))
	for i, v := range e.Social {
		social[i] = v.ToEntity()
	}
	return social
}

func (e *MongoAccountSocial) ToEntity() account.EntitySocial {
	return account.EntitySocial{
		Platform:   e.Platform,
		Value:      e.Value,
		FixedValue: e.FixedValue,
	}
}

func (e *MongoAccount) FromEntity(entity *account.Entity) *MongoAccount {
	e.UserUUID = entity.UserUUID
	e.UserName = entity.UserName
	e.FullName = entity.FullName
	e.Description = entity.Description
	e.Social = e.FromSocialEntity(entity.Social)
	e.IsActive = entity.IsActive
	e.CompletedRate = entity.CompletedRate
	e.IsDeleted = entity.IsDeleted
	e.IsVerified = entity.IsVerified
	e.BirthDate = entity.BirthDate
	e.CreatedAt = entity.CreatedAt
	e.UpdatedAt = entity.UpdatedAt
	return e
}

func (e *MongoAccount) FromSocialEntity(entity []account.EntitySocial) []MongoAccountSocial {
	social := make([]MongoAccountSocial, len(entity))
	for i, v := range entity {
		social[i] = e.FromSocialEntityItem(v)
	}
	return social
}

func (e *MongoAccount) FromSocialEntityItem(entity account.EntitySocial) MongoAccountSocial {
	return MongoAccountSocial{
		Platform:   entity.Platform,
		Value:      entity.Value,
		FixedValue: entity.FixedValue,
	}
}
