package models

import "time"

type UserModel struct {
	Id            int64     `json:"id,omitempty"            db:"id"`
	IdNationality int64     `json:"idNationality,omitempty" db:"id_nationality"`
	IdConfig      int64     `json:"idConfig,omitempty"        db:"id_config"`
	Firstname     string    `json:"firstname,omitempty"      db:"first_name"`
	Lastname      string    `json:"lastname,omitempty"      db:"last_name"`
	Email         string    `json:"email,omitempty"         db:"email"`
	Password      string    `json:"password,omitempty"      db:"password"`
	PhonePrefix   string    `json:"phonePrefix,omitempty"     db:"phone_prefix"`
	Ddd           string    `json:"ddd,omitempty"           db:"ddd"`
	Phone         string    `json:"phone,omitempty"         db:"phone"`
	CreatedAt     time.Time `json:"createdAt,omitempty"     db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"     db:"updated_at"`
	Deleted       bool      `json:"deleted,omitempty"       db:"deleted"`
}

type UserDetails struct {
	Id                int64  `json:"id,omitempty"                 db:"id"`
	Firstname         string `json:"first_name,omitempty"          db:"first_name"`
	Lastname          string `json:"last_name,omitempty"          db:"last_name"`
	Email             string `json:"email,omitempty"              db:"email"`
	Password          string `json:"password,omitempty"           db:"password"`
	PhonePrefix       string `json:"phone_prefix,omitempty"         db:"phone_prefix"`
	Ddd               string `json:"ddd,omitempty"                db:"ddd"`
	Phone             string `json:"phone,omitempty"              db:"phone"`
	Deleted           string `json:"deleted,omitempty"            db:"deleted"`
	IdNationality     string `json:"id_nationality,omitempty"     db:"id_nationality"`
	NationalityName   string `json:"nationality,omitempty"        db:"nationality"`
	NationalityKey    string `json:"nationality_key,omitempty"    db:"nationality_key"`
	IdConfig          string `json:"id_config,omitempty"            db:"id_config"`
	AppTheme          string `json:"app_theme,omitempty"          db:"config_theme"`
	BiometryActivated string `json:"biometry_activated,omitempty" db:"config_biometry_activated"`
	AlertOnEmail      bool   `json:"alert_on_email,omitempty"     db:"config_alert_on_email"`
	AlertOnMobile     bool   `json:"alert_on_mobile,omitempty"    db:"config_alert_on_mobile"`
	AppLanguage       string `json:"app_language,omitempty"       db:"config_language"`
	AppLanguageKey    string `json:"app_language_key,omitempty"   db:"config_language_key"`
}
