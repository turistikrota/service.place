package feature

type messages struct {
	Failed      string
	InvalidUUID string
	NotFound    string
}

var I18nMessages = messages{
	Failed:      "feature_failed",
	InvalidUUID: "feature_invalid_uuid",
	NotFound:    "feature_not_found",
}
