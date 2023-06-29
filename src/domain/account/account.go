package account

import "time"

type Entity struct {
	UUID          string         `json:"uuid"`
	UserUUID      string         `json:"user_uuid"`
	UserName      string         `json:"user_name"`
	FullName      string         `json:"full_name"`
	Description   string         `json:"description"`
	Social        []EntitySocial `json:"social"`
	IsActive      bool           `json:"is_active"`
	CompletedRate int            `json:"completed_rate"`
	IsDeleted     bool           `json:"is_deleted"`
	IsVerified    bool           `json:"is_verified"`
	BirthDate     *time.Time     `json:"birth_date"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
}

type EntitySocial struct {
	Platform   string `json:"platform"`
	Value      string `json:"value"`
	FixedValue string `json:"fixed_value"`
}
