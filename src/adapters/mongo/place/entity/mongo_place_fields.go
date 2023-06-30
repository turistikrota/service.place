package entity

type fields struct {
	UUID             string
	FeatureUUIDs     string
	Images           string
	Translations     string
	AverageTimeSpent string
	Review           string
	IsActive         string
	IsDeleted        string
	IsPayed          string
	UpdatedAt        string
	CreatedAt        string
	Coordinates      string
	Type             string
}

type translationFields struct {
	Title       string
	Description string
	Slug        string
	MarkdownURL string
	Seo         string
}

type seoFields struct {
	Title       string
	Description string
	Keywords    string
}

type imageFields struct {
	Url   string
	Order string
}

type timeSpentFields struct {
	Min string
	Max string
}

type reviewFields struct {
	Total        string
	AveragePoint string
}

var Fields = fields{
	UUID:             "_id",
	FeatureUUIDs:     "feature_uuids",
	Images:           "images",
	Translations:     "translations",
	AverageTimeSpent: "average_time_spent",
	Review:           "review",
	IsActive:         "is_active",
	IsDeleted:        "is_deleted",
	IsPayed:          "is_payed",
	UpdatedAt:        "updated_at",
	CreatedAt:        "created_at",
	Coordinates:      "coordinates",
	Type:             "type",
}

var TranslationFields = translationFields{
	Title:       "title",
	Description: "description",
	Slug:        "slug",
	MarkdownURL: "markdown_url",
	Seo:         "seo",
}

var SeoFields = seoFields{
	Title:       "title",
	Description: "description",
	Keywords:    "keywords",
}

var ImageFields = imageFields{
	Url:   "url",
	Order: "order",
}

var TimeSpentFields = timeSpentFields{
	Min: "min",
	Max: "max",
}

var ReviewFields = reviewFields{
	Total:        "total",
	AveragePoint: "average_point",
}

func TranslationField(lang string, field string) string {
	return Fields.Translations + "." + lang + "." + field
}

func TranslationSeoField(lang string, field string) string {
	return TranslationField(lang, TranslationFields.Seo+"."+field)
}

func TimeSpentField(field string) string {
	return Fields.AverageTimeSpent + "." + field
}

func ReviewField(field string) string {
	return Fields.Review + "." + field
}

func ImageField(field string) string {
	return Fields.Images + "." + field
}
