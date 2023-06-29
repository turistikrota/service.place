package req

import "api.turistikrota.com/place/src/app/command"

type PlatformDetailRequest struct {
	Slug string `param:"slug" validate:"required"`
}

func (r *PlatformDetailRequest) ToDeleteCommand(adminUUID string) command.PlatformDeleteCommand {
	return command.PlatformDeleteCommand{
		AdminUUID: adminUUID,
		Slug:      r.Slug,
	}
}

func (r *PlatformDetailRequest) ToDisableCommand(adminUUID string) command.PlatformDisableCommand {
	return command.PlatformDisableCommand{
		AdminUUID: adminUUID,
		Slug:      r.Slug,
	}
}

func (r *PlatformDetailRequest) ToEnableCommand(adminUUID string) command.PlatformEnableCommand {
	return command.PlatformEnableCommand{
		AdminUUID: adminUUID,
		Slug:      r.Slug,
	}
}
