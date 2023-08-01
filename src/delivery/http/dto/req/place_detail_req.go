package req

import (
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/app/query"
)

type PlaceDetailRequest struct {
	UUID string `param:"uuid" validate:"required,object_id"`
}

func (r *PlaceDetailRequest) ToDisableCommand() command.PlaceDisableCommand {
	return command.PlaceDisableCommand{
		UUID: r.UUID,
	}
}

func (r *PlaceDetailRequest) ToEnableCommand() command.PlaceEnableCommand {
	return command.PlaceEnableCommand{
		UUID: r.UUID,
	}
}

func (r *PlaceDetailRequest) ToDeleteCommand() command.PlaceDeleteCommand {
	return command.PlaceDeleteCommand{
		UUID: r.UUID,
	}
}

func (r *PlaceDetailRequest) ToQuery() query.AdminPlaceViewQuery {
	return query.AdminPlaceViewQuery{
		UUID: r.UUID,
	}
}
