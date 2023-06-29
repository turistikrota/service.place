package req

import (
	"time"

	"api.turistikrota.com/place/src/app/command"
	"github.com/turistikrota/service.shared/formats"
)

type AccountUpdateRequest struct {
	UserName    string
	FullName    string `json:"fullName" validate:"required,max=70,min=3"`
	Description string `json:"description" validate:"required,max=1000"`
	BirthDate   string `json:"birthDate" validate:"required,datetime=2006-01-02"`
}

func (r *AccountUpdateRequest) LoadDetail(detail *AccountDetailRequest) {
	r.UserName = detail.UserName
}

func (r *AccountUpdateRequest) ToCommand(userUUID string) command.AccountUpdateCommand {
	birth, _ := time.Parse(formats.DateYYYYMMDD, r.BirthDate)
	return command.AccountUpdateCommand{
		UserUUID:    userUUID,
		UserName:    r.UserName,
		FullName:    r.FullName,
		Description: r.Description,
		BirthDate:   &birth,
	}
}
