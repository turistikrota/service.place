package http

type successMessages struct {
	FeatureCreate  string
	FeatureUpdate  string
	FeatureDelete  string
	FeatureDisable string
	FeatureEnable  string
	FeatureList    string
	PlaceCreate    string
	PlaceUpdate    string
	PlaceDelete    string
	PlaceDisable   string
	PlaceEnable    string
	PlaceFilter    string
	PlaceView      string
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
		FeatureCreate:  "http_success_feature_create",
		FeatureUpdate:  "http_success_feature_update",
		FeatureDelete:  "http_success_feature_delete",
		FeatureDisable: "http_success_feature_disable",
		FeatureEnable:  "http_success_feature_enable",
		FeatureList:    "http_success_feature_list",
		PlaceCreate:    "http_success_place_create",
		PlaceUpdate:    "http_success_place_update",
		PlaceDelete:    "http_success_place_delete",
		PlaceDisable:   "http_success_place_disable",
		PlaceEnable:    "http_success_place_enable",
		PlaceFilter:    "http_success_place_filter",
		PlaceView:      "http_success_place_view",
	},
	Error: errorMessages{
		RequiredAuth:      "http_error_required_auth",
		CurrentUserAccess: "http_error_current_user_access",
		AdminRoute:        "http_error_admin_route",
	},
}
