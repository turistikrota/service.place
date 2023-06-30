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
	Coordinates      []float64
	IsPayed          bool
	Type             Type
}

func (f Factory) New(config NewConfig) *Entity {
	return &Entity{
		FeatureUUIDs:     config.FeatureUUIDs,
		Images:           config.Images,
		Translations:     config.Translations,
		AverageTimeSpent: config.AverageTimeSpent,
		Review: Review{
			Total:        0,
			AveragePoint: 0,
		},
		Coordinates: config.Coordinates,
		IsPayed:     config.IsPayed,
		IsActive:    true,
		IsDeleted:   false,
		Type:        config.Type,
	}
}
