package req

import (
	"github.com/turistikrota/service.place/src/app/command"
	"github.com/turistikrota/service.place/src/app/query"
)

type FeatureDetailRequest struct {
	UUID string `param:"uuid" validate:"required,object_id"`
}

func (r *FeatureDetailRequest) ToDisableCommand() command.FeatureDisableCommand {
	return command.FeatureDisableCommand{
		UUID: r.UUID,
	}
}

func (r *FeatureDetailRequest) ToEnableCommand() command.FeatureEnableCommand {
	return command.FeatureEnableCommand{
		UUID: r.UUID,
	}
}

func (r *FeatureDetailRequest) ToDeleteCommand() command.FeatureDeleteCommand {
	return command.FeatureDeleteCommand{
		UUID: r.UUID,
	}
}

func (r *FeatureDetailRequest) ToQuery() query.AdminFeatureDetailQuery {
	return query.AdminFeatureDetailQuery{
		UUID: r.UUID,
	}
}
