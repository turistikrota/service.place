package entity

type fields struct {
	UUID         string
	Name         string
	Slug         string
	Regexp       string
	Prefix       string
	Translations string
	IsActive     string
	IsDeleted    string
	UpdatedAt    string
	CreatedAt    string
}

type translationsFields struct {
	Locale      string
	Name        string
	Placeholder string
	Description string
}

var Fields = fields{
	UUID:         "uuid",
	Name:         "name",
	Slug:         "slug",
	Regexp:       "regexp",
	Prefix:       "prefix",
	Translations: "translations",
	IsActive:     "is_active",
	IsDeleted:    "is_deleted",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

var TranslationsFields = translationsFields{
	Locale:      "locale",
	Name:        "name",
	Placeholder: "placeholder",
	Description: "description",
}

func TranslationField(locale string, path string) string {
	return Fields.Translations + "." + locale + "." + path
}

func Translation(locale string) string {
	return Fields.Translations + "." + locale
}
