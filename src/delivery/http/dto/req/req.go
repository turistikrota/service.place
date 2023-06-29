package req

type Request interface {
	AccountDetail() *AccountDetailRequest
	AccountCreate() *AccountCreateRequest
	AccountUpdate() *AccountUpdateRequest
	AccountSocial() *AccountSocialRequest
	AccountSocialAction() *AccountSocialActionRequest
	AccountListMy() *AccountListMyRequest
	PlatformCreate() *PlatformCreateRequest
	PlatformDetail() *PlatformDetailRequest
	PlatformAction() *PlatformActionRequest
	PlatformGet() *PlatformGetRequest
	PlatformList() *PlatformListRequest
	PlatformTranslation() *PlatformTranslationRequest
	PlatformTranslationAction() *PlatformTranslationActionRequest
}

type request struct{}

func New() Request {
	return &request{}
}

func (r *request) AccountDetail() *AccountDetailRequest {
	return &AccountDetailRequest{}
}

func (r *request) AccountCreate() *AccountCreateRequest {
	return &AccountCreateRequest{}
}

func (r *request) AccountUpdate() *AccountUpdateRequest {
	return &AccountUpdateRequest{}
}

func (r *request) AccountSocial() *AccountSocialRequest {
	return &AccountSocialRequest{}
}

func (r *request) AccountSocialAction() *AccountSocialActionRequest {
	return &AccountSocialActionRequest{}
}

func (r *request) AccountListMy() *AccountListMyRequest {
	return &AccountListMyRequest{}
}

func (r *request) PlatformCreate() *PlatformCreateRequest {
	return &PlatformCreateRequest{}
}

func (r *request) PlatformDetail() *PlatformDetailRequest {
	return &PlatformDetailRequest{}
}

func (r *request) PlatformAction() *PlatformActionRequest {
	return &PlatformActionRequest{}
}

func (r *request) PlatformGet() *PlatformGetRequest {
	return &PlatformGetRequest{}
}

func (r *request) PlatformList() *PlatformListRequest {
	return &PlatformListRequest{}
}

func (r *request) PlatformTranslation() *PlatformTranslationRequest {
	return &PlatformTranslationRequest{}
}

func (r *request) PlatformTranslationAction() *PlatformTranslationActionRequest {
	return &PlatformTranslationActionRequest{}
}
