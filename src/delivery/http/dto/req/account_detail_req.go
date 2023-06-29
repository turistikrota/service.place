package req

import (
	"api.turistikrota.com/account/src/app/command"
	"api.turistikrota.com/account/src/app/query"
)

type AccountDetailRequest struct {
	UserName string `param:"userName" validate:"required,username"`
}

func (r *AccountDetailRequest) ToDeleteCommand(userUUID string) command.AccountDeleteCommand {
	return command.AccountDeleteCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
	}
}

func (r *AccountDetailRequest) ToDisableCommand(userUUID string) command.AccountDisableCommand {
	return command.AccountDisableCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
	}
}

func (r *AccountDetailRequest) ToEnableCommand(userUUID string) command.AccountEnableCommand {
	return command.AccountEnableCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
	}
}

func (r *AccountDetailRequest) ToGetQuery(userUUID string) query.AccountGetQuery {
	return query.AccountGetQuery{
		UserUUID: userUUID,
		Name:     r.UserName,
	}
}

func (r *AccountDetailRequest) ToProfileQuery() query.AccountProfileViewQuery {
	return query.AccountProfileViewQuery{
		Name: r.UserName,
	}
}
