package res

import (
	"time"

	"api.turistikrota.com/account/src/app/query"
	"api.turistikrota.com/account/src/domain/account"
	"github.com/turistikrota/service.shared/helper"
)

type AccountGetResponse struct {
	UserName      string                     `json:"userName"`
	FullName      string                     `json:"fullName"`
	AvatarURL     string                     `json:"avatarUrl"`
	Description   string                     `json:"description"`
	Social        []AccountGetResponseSocial `json:"social"`
	IsActive      bool                       `json:"isActive"`
	CompletedRate int                        `json:"completedRate"`
	IsVerified    bool                       `json:"isVerified"`
	BirthDate     *time.Time                 `json:"birthDate"`
	CreatedAt     *time.Time                 `json:"createdAt"`
	UpdatedAt     *time.Time                 `json:"updatedAt"`
}

type AccountGetResponseSocial struct {
	Platform string `json:"platform"`
	Url      string `json:"url"`
}

func (r *response) AccountGet(res *query.AccountGetResult) *AccountGetResponse {
	return &AccountGetResponse{
		UserName:      res.Entity.UserName,
		FullName:      res.Entity.FullName,
		Description:   res.Entity.Description,
		Social:        r.AccountGetResponseSocial(res.Entity.Social),
		AvatarURL:     helper.CDN.DressAvatar(res.Entity.UserName),
		IsActive:      res.Entity.IsActive,
		CompletedRate: res.Entity.CompletedRate,
		IsVerified:    res.Entity.IsVerified,
		BirthDate:     res.Entity.BirthDate,
		CreatedAt:     res.Entity.CreatedAt,
		UpdatedAt:     res.Entity.UpdatedAt,
	}
}

func (r *response) AccountGetResponseSocial(social []account.EntitySocial) []AccountGetResponseSocial {
	res := make([]AccountGetResponseSocial, 0)
	for _, item := range social {
		res = append(res, AccountGetResponseSocial{
			Platform: item.Platform,
			Url:      item.FixedValue,
		})
	}
	return res
}
