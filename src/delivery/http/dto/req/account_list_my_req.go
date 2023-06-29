package req

import "api.turistikrota.com/account/src/app/query"

type AccountListMyRequest struct{}

func (r *AccountListMyRequest) ToQuery(userUUID string) query.AccountListMyQuery {
	return query.AccountListMyQuery{
		UserUUID: userUUID,
	}
}
