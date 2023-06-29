package res

import "api.turistikrota.com/account/src/app/command"

type AccountCreateResponse struct {
	AccountName string `json:"accountName"`
}

func (r *response) AccountCreate(res *command.AccountCreateResult) *AccountCreateResponse {
	return &AccountCreateResponse{
		AccountName: res.AccountName,
	}
}
