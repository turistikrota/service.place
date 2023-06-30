package req

type Request interface {
	FeatureDetail() *FeatureDetailRequest
	Pagination() *PaginationRequest
	FeatureCreate() *FeatureCreateRequest
	FeatureUpdate() *FeatureUpdateRequest
	PlaceDetail() *PlaceDetailRequest
	PlaceFilter() *PlaceFilterRequest
	PlaceCreate() *PlaceCreateRequest
	PlaceUpdate() *PlaceUpdateRequest
	PlaceView() *PlaceViewRequest
}

type request struct{}

func New() Request {
	return &request{}
}

func (r *request) FeatureDetail() *FeatureDetailRequest {
	return &FeatureDetailRequest{}
}

func (r *request) Pagination() *PaginationRequest {
	return &PaginationRequest{}
}

func (r *request) FeatureCreate() *FeatureCreateRequest {
	return &FeatureCreateRequest{}
}

func (r *request) FeatureUpdate() *FeatureUpdateRequest {
	return &FeatureUpdateRequest{}
}

func (r *request) PlaceDetail() *PlaceDetailRequest {
	return &PlaceDetailRequest{}
}

func (r *request) PlaceFilter() *PlaceFilterRequest {
	return &PlaceFilterRequest{}
}

func (r *request) PlaceCreate() *PlaceCreateRequest {
	return &PlaceCreateRequest{}
}

func (r *request) PlaceUpdate() *PlaceUpdateRequest {
	return &PlaceUpdateRequest{}
}

func (r *request) PlaceView() *PlaceViewRequest {
	return &PlaceViewRequest{}
}
