package place

import "time"

type Entity struct {
	UUID             string                   `json:"uuid"`
	FeatureUUIDs     []string                 `json:"featureUUIDs"`
	Images           []Image                  `json:"images"`
	Translations     map[Locale]*Translations `json:"translations"`
	AverageTimeSpent TimeSpent                `json:"averageTimeSpent"`
	Review           Review                   `json:"review"`
	Restorations     []Restoration            `json:"restorations"`
	Coordinates      []float64                `json:"coordinates"`
	IsActive         bool                     `json:"is_active"`
	IsDeleted        bool                     `json:"is_deleted"`
	IsPayed          bool                     `json:"is_payed"`
	Type             Type                     `json:"type"`
	UpdatedAt        time.Time                `json:"updated_at"`
	CreatedAt        time.Time                `json:"created_at"`
}

type Image struct {
	Url   string `json:"url"`
	Order int16  `json:"order"`
}

type TimeSpent struct {
	Min int16 `json:"min"`
	Max int16 `json:"max"`
}

type Restoration struct {
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type Review struct {
	Total        int16   `json:"total"`
	AveragePoint float32 `json:"averagePoint"`
}

type Translations struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	MarkdownURL string `json:"markdownUrl"`
	Seo         Seo    `json:"seo"`
}

type Seo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
}

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

func (l Locale) String() string {
	return string(l)
}

type Type string

const (
	TypeEating    Type = "eating"
	TypeCoffee    Type = "coffee"
	TypeBar       Type = "bar"
	TypeBeach     Type = "beach"
	TypeAmaze     Type = "amaze"
	TypeShopping  Type = "shopping"
	TypeTransport Type = "transport"
	TypeCulture   Type = "culture"
	TypeNature    Type = "nature"
	TypeHealth    Type = "health"
	TypeSport     Type = "sport"
	TypeNightlife Type = "nightlife"
	TypeGarden    Type = "garden"
	TypeTemple    Type = "temple"
	TypeMuseum    Type = "museum"
	TypeAntique   Type = "antique"
	TypePark      Type = "park"
	TypeThemePark Type = "themePark"
	TypeOther     Type = "other"
)

func (t Type) String() string {
	return string(t)
}

func (t Type) IsType() bool {
	return t == TypeEating || t == TypeCoffee || t == TypeBar || t == TypeBeach || t == TypeAmaze || t == TypeShopping || t == TypeTransport || t == TypeCulture || t == TypeNature || t == TypeHealth || t == TypeSport || t == TypeNightlife || t == TypeGarden || t == TypeTemple || t == TypeMuseum || t == TypeAntique || t == TypePark || t == TypeThemePark || t == TypeOther
}
