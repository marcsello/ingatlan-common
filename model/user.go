package model

import "strings"

type User struct {
	// All these data are received from Telegram
	ID        int64  `json:"id" gorm:"primarykey"`               // This must be a signed int, because telegram assign negative id to groups
	FirstName string `json:"first_name" gorm:"type:varchar(64)"` // https://limits.tginfo.me/en
	LastName  string `json:"last_name" gorm:"type:varchar(64)"`
	Username  string `json:"username" gorm:"type:varchar(32)"` //https://core.telegram.org/method/account.checkUsername
	PhotoUrl  string `json:"photo_url" gorm:"type:varchar(128)"`
	// Those are needs to be *boolean, otherwise gorm won't update them: https://stackoverflow.com/questions/56653423/gorm-doesnt-update-boolean-field-to-false
	Active *bool `json:"active" gorm:"default:false"`
	Admin  *bool `json:"admin" gorm:"default:false"`
}

func (u *User) IsActive() bool {
	return u.Active != nil && *u.Active
}

func (u *User) IsAdmin() bool {
	return u.Admin != nil && *u.Admin
}

// Greet returns a proper name compiled from the FirstName, LastName and Username fields
func (u *User) Greet() string {
	name := u.FirstName // first name must be always present
	if u.LastName != "" {
		name += " " + u.LastName
	}

	if strings.TrimSpace(name) == "" {
		return "@" + u.Username
	} else {
		return name
	}

}

func (u *User) HavePhoto() bool {
	return u.PhotoUrl != ""
}
