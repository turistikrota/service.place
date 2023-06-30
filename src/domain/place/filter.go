package place

type EntityFilter struct {
	Locale           string
	Query            string
	Coordinates      []float64
	FeatureUUIDs     []string
	AverageTimeSpent *TimeSpent
	Distance         float64
	IsPayed          *bool
	MinReview        *int16
	MaxReview        *int16
	Types            []Type
	MinAveragePoint  *float32
	MaxAveragePoint  *float32
}

func (e *EntityFilter) IsZero() bool {
	return e.Locale == "" &&
		e.Query == "" &&
		e.Coordinates == nil &&
		e.FeatureUUIDs == nil &&
		e.AverageTimeSpent == nil &&
		e.Distance == 0 &&
		e.IsPayed == nil &&
		e.MinReview == nil &&
		e.MaxReview == nil &&
		e.MinAveragePoint == nil &&
		e.MaxAveragePoint == nil
}
