package res

import "github.com/turistikrota/service.place/src/app/command"

type PlaceCreateResponse struct {
	UUID string `json:"uuid"`
}

func (r *response) PlaceCreate(res *command.PlaceCreateResult) *PlaceCreateResponse {
	return &PlaceCreateResponse{
		UUID: res.UUID,
	}
}
