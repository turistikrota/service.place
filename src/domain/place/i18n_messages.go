package place

type messages struct {
	Failed      string
	InvalidUUID string
	NotFound    string
}

var I18nMessages = messages{
	Failed:      "place_failed",
	InvalidUUID: "place_invalid_uuid",
	NotFound:    "place_not_found",
}
