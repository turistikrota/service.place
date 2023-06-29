package feature

type messages struct {
	Failed      string
	InvalidUUID string
}

var I18nMessages = messages{
	Failed:      "feature_failed",
	InvalidUUID: "feature_invalid_uuid",
}
