package place

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newPlaceErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	FeatureUUIDs     []string
	Images           []Image
	Translations     map[Locale]Translations
	AverageTimeSpent TimeSpent
	Review           Review
	Coordinates      []float64
	IsPayed          bool
}

func (f Factory) New(config NewConfig) *Entity {
	return &Entity{
		FeatureUUIDs:     config.FeatureUUIDs,
		Images:           config.Images,
		Translations:     config.Translations,
		AverageTimeSpent: config.AverageTimeSpent,
		Review:           config.Review,
		Coordinates:      config.Coordinates,
		IsPayed:          config.IsPayed,
		IsActive:         true,
		IsDeleted:        false,
	}
}
