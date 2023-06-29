package place

type EntityFilter struct {
	Locale           string
	Query            string
	Coordinates      []float64
	FeatureUUIDs     []string
	AverageTimeSpent TimeSpent
	Distance         float64
	IsPayed          bool
	MinReview        int16
	MaxReview        int16
	MinAveragePoint  float32
	MaxAveragePoint  float32
}
