package req

import "api.turistikrota.com/place/src/app/command"

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
