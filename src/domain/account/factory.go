package account

import (
	"time"

	"github.com/mixarchitecture/i18np"
)

type Factory struct {
	Errors Errors
	minAge int
	maxAge int
}

func NewFactory() Factory {
	return Factory{
		Errors: newAccountErrors(),
		minAge: 13,
		maxAge: 100,
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) NewAccount(userUUID string, username string) *Entity {
	t := time.Now()
	e := &Entity{
		UserUUID:      userUUID,
		UserName:      username,
		IsActive:      false,
		CompletedRate: 0,
		IsDeleted:     false,
		IsVerified:    false,
		CreatedAt:     &t,
		UpdatedAt:     &t,
	}
	e.CompletedRate = f.CalcCompletedRate(e)
	return e
}

func (f Factory) CalcCompletedRate(e *Entity) int {
	var rate int
	denominatorCount := 5 // 5 field
	list := []string{e.UserName, e.FullName, e.Description}
	if e.BirthDate != nil && e.BirthDate.Year() > 0 {
		rate += 1
	}
	for _, v := range list {
		if v != "" {
			rate++
		}
	}
	if len(e.Social) > 0 {
		rate++
	}
	return rate * 100 / denominatorCount
}

func (f Factory) Validate(e *Entity) *i18np.Error {
	if err := f.validateUserName(e.UserName); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateUserName(username string) *i18np.Error {
	if username == "" {
		return f.Errors.UserNameRequired()
	}
	return nil
}

func (f Factory) ValidateMinAge(birthDate *time.Time) *i18np.Error {
	if birthDate == nil {
		return nil
	}
	userAge := time.Now().Year() - birthDate.Year()
	if userAge < f.minAge {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge == f.minAge && time.Now().Month() < birthDate.Month() {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge == f.minAge && time.Now().Month() == birthDate.Month() && time.Now().Day() < birthDate.Day() {
		return f.Errors.MinAge(f.minAge)
	}
	if userAge > 150 {
		return f.Errors.MaxAge(f.maxAge)
	}
	return nil
}
