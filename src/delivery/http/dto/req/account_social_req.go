package req

import "api.turistikrota.com/place/src/app/command"

type AccountSocialRequest struct {
	AccountDetailRequest
	Platform string `param:"platform" validate:"required"`
}

func (r *AccountSocialRequest) ToRemoveCommand(userUUID string) command.AccountSocialRemoveCommand {
	return command.AccountSocialRemoveCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		Platform:    r.Platform,
	}
}
