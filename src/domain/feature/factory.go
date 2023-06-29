package feature

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newFeatureErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) New(icon string, translations map[Locale]Translations) *Entity {
	return &Entity{
		Icon:         icon,
		Translations: translations,
		IsActive:     true,
		IsDeleted:    false,
	}
}
