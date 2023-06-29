package entity

type fields struct {
	UUID          string
	UserUUID      string
	UserName      string
	FullName      string
	Description   string
	Social        string
	IsActive      string
	CompletedRate string
	IsDeleted     string
	IsVerified    string
	BirthDate     string
	CreatedAt     string
	UpdatedAt     string
}

type socialFields struct {
	Platform   string
	Value      string
	FixedValue string
}

var Fields = fields{
	UUID:          "uuid",
	UserUUID:      "user_uuid",
	UserName:      "user_name",
	FullName:      "full_name",
	Description:   "description",
	Social:        "social",
	IsActive:      "is_active",
	CompletedRate: "completed_rate",
	IsDeleted:     "is_deleted",
	IsVerified:    "is_verified",
	BirthDate:     "birth_date",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var SocialFields = socialFields{
	Platform:   "platform",
	Value:      "value",
	FixedValue: "fixed_value",
}

func SocialField(path string) string {
	return Fields.Social + "." + path
}

func SocialFieldInArray(path string) string {
	return Fields.Social + ".$." + path
}
