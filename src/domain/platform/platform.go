package platform

import (
	"time"
)

type Entity struct {
	UUID         string                  `json:"uuid"`
	Name         string                  `json:"name"`
	Slug         string                  `json:"slug"`
	Regexp       string                  `json:"regexp"`
	Prefix       string                  `json:"prefix"`
	Translations map[Locale]Translations `json:"translations"`
	IsActive     bool                    `json:"is_active"`
	IsDeleted    bool                    `json:"is_deleted"`
	UpdatedAt    time.Time               `json:"updated_at"`
	CreatedAt    time.Time               `json:"created_at"`
}

type Translations struct {
	Name        string `json:"name"`
	Placeholder string `json:"placeholder"`
	Description string `json:"description"`
}

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

func (l Locale) String() string {
	return string(l)
}
