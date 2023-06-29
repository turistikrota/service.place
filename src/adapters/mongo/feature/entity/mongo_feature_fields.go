package entity

type fields struct {
	UUID         string
	Icon         string
	Translations string
	IsActive     string
	IsDeleted    string
	UpdatedAt    string
	CreatedAt    string
}

type translationFields struct {
	Title       string
	Description string
}

var Fields = fields{
	UUID:         "_id",
	Icon:         "icon",
	Translations: "translations",
	IsActive:     "is_active",
	IsDeleted:    "is_deleted",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

var TranslationFields = translationFields{
	Title:       "title",
	Description: "description",
}

func TranslationField(locale string, field string) string {
	return Fields.Translations + "." + locale + "." + field
}
