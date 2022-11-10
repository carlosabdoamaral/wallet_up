package models

import "time"

type AccountModel struct {
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

type AccountDetails struct {
	Account struct {
		Id              int64  `json:"id,omitempty"                 db:"id"`
		Firstname       string `json:"first_name,omitempty"          db:"first_name"`
		Lastname        string `json:"last_name,omitempty"          db:"last_name"`
		Email           string `json:"email,omitempty"              db:"email"`
		Password        string `json:"password,omitempty"           db:"password"`
		PhonePrefix     int64  `json:"phone_prefix,omitempty"         db:"phone_prefix"`
		Ddd             string `json:"ddd,omitempty"                db:"ddd"`
		Phone           string `json:"phone,omitempty"              db:"phone"`
		Deleted         bool   `json:"deleted,omitempty"            db:"deleted"`
		IdNationality   int64  `json:"id_nationality,omitempty"     db:"id_nationality"`
		NationalityName string `json:"nationality,omitempty"        db:"nationality"`
		NationalityKey  string `json:"nationality_key,omitempty"    db:"nationality_key"`
	} `json:"account"`

	Config AppConfigDetails `json:"config"`
}

type NewAccountRequest struct {
	IdNationality int64  `json:"id_nationality,omitempty"     db:"id_nationality"`
	IdConfig      int64  `json:"id_config,omitempty"            db:"id_config"`
	Firstname     string `json:"first_name,omitempty"          db:"first_name"`
	Lastname      string `json:"last_name,omitempty"          db:"last_name"`
	Email         string `json:"email,omitempty"              db:"email"`
	Password      string `json:"password,omitempty"           db:"password"`
	PhonePrefix   string `json:"phone_prefix,omitempty"         db:"phone_prefix"`
	Ddd           string `json:"ddd,omitempty"                db:"ddd"`
	Phone         string `json:"phone,omitempty"              db:"phone"`
}

type AccountId struct {
	AccountId int64 `json:"account_id,omitempty"     db:"id"`
}

type EditAccountRequest struct {
	Account EditAccountInfos `json:"account"`
	Config  AppConfigModel   `json:"config"`
}

type EditAccountInfos struct {
	Id            int64  `json:"id,omitempty"            db:"id"`
	IdNationality int64  `json:"idNationality,omitempty" db:"id_nationality"`
	Firstname     string `json:"firstname,omitempty"      db:"first_name"`
	Lastname      string `json:"lastname,omitempty"      db:"last_name"`
	Email         string `json:"email,omitempty"         db:"email"`
	Password      string `json:"password,omitempty"      db:"password"`
	PhonePrefix   string `json:"phonePrefix,omitempty"     db:"phone_prefix"`
	Ddd           string `json:"ddd,omitempty"           db:"ddd"`
	Phone         string `json:"phone,omitempty"         db:"phone"`
	Deleted       bool   `json:"deleted,omitempty"       db:"deleted"`
}
