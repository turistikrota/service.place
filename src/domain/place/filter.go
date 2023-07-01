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
	Sort             Sort
	Order            Order
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

type (
	Sort  string
	Order string
)

const (
	SortByMostPopular Sort = "most_popular"
	SortByMostLiked   Sort = "most_liked"
	SortByMostRecent  Sort = "most_recent"
	SortByNearest     Sort = "nearest"
)

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

func (s Sort) IsValid() bool {
	return s == SortByMostPopular ||
		s == SortByMostLiked ||
		s == SortByMostRecent ||
		s == SortByNearest
}

func (o Order) IsValid() bool {
	return o == OrderAsc ||
		o == OrderDesc
}
