package http

type successMessages struct {
	AccountDelete             string
	AccountCreate             string
	AccountUpdate             string
	AccountDisable            string
	AccountEnable             string
	AccountSocialAdd          string
	AccountSocialRemove       string
	AccountSocialUpdate       string
	AccountGet                string
	AccountListMy             string
	AccountProfileView        string
	PlatformCreate            string
	PlatformUpdate            string
	PlatformEnable            string
	PlatformDisable           string
	PlatformDelete            string
	PlatformTranslationCreate string
	PlatformTranslationUpdate string
	PlatformTranslationDelete string
	PlatformGet               string
	PlatformList              string
}

type errorMessages struct {
	RequiredAuth      string
	CurrentUserAccess string
	AdminRoute        string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		AccountDelete:             "http_success_account_deleted",
		AccountCreate:             "http_success_account_created",
		AccountUpdate:             "http_success_account_updated",
		AccountDisable:            "http_success_account_disabled",
		AccountEnable:             "http_success_account_enabled",
		AccountSocialUpdate:       "http_success_account_social_updated",
		AccountSocialAdd:          "http_success_account_social_added",
		AccountSocialRemove:       "http_success_account_social_removed",
		AccountGet:                "http_success_account_get",
		AccountListMy:             "http_success_account_list_my",
		AccountProfileView:        "http_success_account_profile_viewed",
		PlatformCreate:            "http_success_platform_created",
		PlatformUpdate:            "http_success_platform_updated",
		PlatformEnable:            "http_success_platform_enabled",
		PlatformDisable:           "http_success_platform_disabled",
		PlatformDelete:            "http_success_platform_deleted",
		PlatformTranslationCreate: "http_success_platform_translation_created",
		PlatformTranslationUpdate: "http_success_platform_translation_updated",
		PlatformTranslationDelete: "http_success_platform_translation_deleted",
		PlatformGet:               "http_success_platform_get",
		PlatformList:              "http_success_platform_list",
	},
	Error: errorMessages{
		RequiredAuth:      "http_error_required_auth",
		CurrentUserAccess: "http_error_current_user_access",
		AdminRoute:        "http_error_admin_route",
	},
}
