package platform

type messages struct {
	PlatformNotFound                 string
	PlatformNameRequired             string
	PlatformRegexpRequired           string
	PlatformPrefixRequired           string
	PlatformInvalidValue             string
	PlatformTranslationAlreadyExists string
	PlatformTranslationNotExists     string
	PlatformFailed                   string
	PlatformAlreadyExists            string
}

var I18nMessages = messages{
	PlatformNotFound:                 "error_platform_not_found",
	PlatformNameRequired:             "error_platform_name_required",
	PlatformRegexpRequired:           "error_platform_regexp_required",
	PlatformPrefixRequired:           "error_platform_prefix_required",
	PlatformInvalidValue:             "error_platform_invalid_value",
	PlatformTranslationAlreadyExists: "error_platform_translation_already_exists",
	PlatformTranslationNotExists:     "error_platform_translation_not_exists",
	PlatformFailed:                   "error_platform_failed",
	PlatformAlreadyExists:            "error_platform_already_exists",
}
