package place

type EntityFilter struct {
	Locale           string
	Query            string
	Coordinates      []float64
	FeatureUUIDs     []string
	AverageTimeSpent *TimeSpent
	Distance         *float64
	IsPayed          *bool
	MinReview        *int16
	MaxReview        *int16
	Types            []Type
	MinAveragePoint  *float32
	MaxAveragePoint  *float32
	Sort             Sort
	Order            Order
}

func (e *EntityFilter) GetPerfectDistance() float64 {
	if e.Distance == nil {
		return 100
	}
	distances := map[float64]float64{
		7:  500,
		8:  300,
		9:  200,
		10: 100,
		11: 50,
		12: 20,
		13: 10,
		14: 5,
		15: 3,
	}
	if distance, ok := distances[*e.Distance]; ok {
		return distance
	}
	return 10
}

func (e *EntityFilter) IsZero() bool {
	return e.Locale == "" &&
		e.Query == "" &&
		len(e.Coordinates) == 0 &&
		len(e.FeatureUUIDs) == 0 &&
		e.AverageTimeSpent == nil &&
		e.Distance == nil &&
		e.IsPayed == nil &&
		e.MinReview == nil &&
		e.MaxReview == nil &&
		e.MinAveragePoint == nil &&
		e.MaxAveragePoint == nil &&
		len(e.Types) == 0 &&
		!e.Sort.IsValid() &&
		!e.Order.IsValid()
}

type (
	Sort  string
	Order string
)

const (
	SortByMostRecent Sort = "most_recent"
	SortByNearest    Sort = "nearest"
)

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

func (s Sort) IsValid() bool {
	return s == SortByMostRecent ||
		s == SortByNearest
}

func (o Order) IsValid() bool {
	return o == OrderAsc ||
		o == OrderDesc
}
