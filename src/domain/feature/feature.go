package feature

import "time"

type Entity struct {
	UUID         string                  `json:"uuid"`
	MainUUID     string                  `json:"main_uuid"`
	Icon         string                  `json:"icon"`
	Translations map[Locale]Translations `json:"translations"`
	IsActive     bool                    `json:"is_active"`
	IsDeleted    bool                    `json:"is_deleted"`
	UpdatedAt    time.Time               `json:"updated_at"`
	CreatedAt    time.Time               `json:"created_at"`
}

type Translations struct {
	Title       string `json:"title"`
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
