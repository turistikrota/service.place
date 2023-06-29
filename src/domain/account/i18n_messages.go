package account

type messages struct {
	AccountUserNameRequired string
	AccountUserCodeRequired string
	AccountUserCodeInvalid  string
	AccountAlreadyExist     string
	AccountMinAge           string
	AccountMaxAge           string
	AccountFailed           string
	AccountNotFound         string
}

var I18nMessages = messages{
	AccountUserNameRequired: "error_account_user_name_required",
	AccountUserCodeRequired: "error_account_user_code_required",
	AccountUserCodeInvalid:  "error_account_user_code_invalid",
	AccountAlreadyExist:     "error_account_already_exist",
	AccountMinAge:           "error_account_min_age",
	AccountMaxAge:           "error_account_max_age",
	AccountFailed:           "error_account_failed",
	AccountNotFound:         "error_account_not_found",
}
