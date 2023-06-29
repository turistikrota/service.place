package req

import "api.turistikrota.com/account/src/app/command"

type PlatformCreateRequest struct {
	NewName string `json:"name" validate:"required"`
	Regexp  string `json:"regexp" validate:"required"`
	Prefix  string `json:"prefix" validate:"required"`
}

func (r *PlatformCreateRequest) ToCommand(adminUUID string) command.PlatformCreateCommand {
	return command.PlatformCreateCommand{
		AdminUUID: adminUUID,
		Name:      r.NewName,
		Regexp:    r.Regexp,
		Prefix:    r.Prefix,
	}
}

type PlatformActionRequest struct {
	Slug    string
	NewName string `json:"name"   validate:"omitempty"`
	Regexp  string `json:"regexp" validate:"required"`
	Prefix  string `json:"prefix" validate:"required"`
}

func (r *PlatformActionRequest) LoadDetail(detail *PlatformDetailRequest) {
	r.Slug = detail.Slug
}

func (r *PlatformActionRequest) ToUpdateCommand(adminUUID string) command.PlatformUpdateCommand {
	return command.PlatformUpdateCommand{
		AdminUUID: adminUUID,
		Name:      r.NewName,
		Regexp:    r.Regexp,
		Prefix:    r.Prefix,
		Slug:      r.Slug,
	}
}
