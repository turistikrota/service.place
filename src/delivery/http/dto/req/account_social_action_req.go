package req

import "api.turistikrota.com/place/src/app/command"

type AccountSocialActionRequest struct {
	UserName string
	Platform string
	Value    string `json:"value" validate:"required"`
}

func (r *AccountSocialActionRequest) LoadSocial(social *AccountSocialRequest) {
	r.UserName = social.UserName
	r.Platform = social.Platform
}

func (r *AccountSocialActionRequest) ToAddCommand(userUUID string) command.AccountSocialAddCommand {
	return command.AccountSocialAddCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		Platform:    r.Platform,
		Value:       r.Value,
	}
}

func (r *AccountSocialActionRequest) ToUpdateCommand(userUUID string) command.AccountSocialUpdateCommand {
	return command.AccountSocialUpdateCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		Platform:    r.Platform,
		Value:       r.Value,
	}
}
