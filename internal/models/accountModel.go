package models

import "time"

// REFATORAR COM BASE NO .proto

type AccountDetails struct {
	Id              int64  `json:"id,omitempty"`
	Firstname       string `json:"first_name,omitempty"`
	Lastname        string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	PhonePrefix     string `json:"phone_prefix,omitempty"`
	Ddd             string `json:"ddd,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	IdNationality   int64  `json:"id_nationality,omitempty"`
	NationalityName string `json:"nationality,omitempty"`
	NationalityKey  string `json:"nationality_key,omitempty"`
}

type AccountDetailsResponse struct {
	Account AccountDetails   `json:"account"`
	Config  AppConfigDetails `json:"config"`
}

type AccountModel struct {
	Id            int64     `json:"id,omitempty"`
	IdNationality int64     `json:"idNationality,omitempty"`
	IdConfig      int64     `json:"idConfig,omitempty"`
	Firstname     string    `json:"firstname,omitempty"`
	Lastname      string    `json:"lastname,omitempty"`
	Email         string    `json:"email,omitempty"`
	Password      string    `json:"password,omitempty"`
	PhonePrefix   string    `json:"phonePrefix,omitempty"`
	Ddd           string    `json:"ddd,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	Deleted       bool      `json:"deleted,omitempty"`
}
type NewAccountRequest struct {
	IdNationality int64  `json:"id_nationality,omitempty"`
	IdConfig      int64  `json:"id_config,omitempty"`
	Firstname     string `json:"first_name,omitempty"`
	Lastname      string `json:"last_name,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"password,omitempty"`
	PhonePrefix   string `json:"phone_prefix,omitempty"`
	Ddd           string `json:"ddd,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type AccountId struct {
	AccountId int64 `json:"account_id,omitempty"`
}

type EditAccountRequest struct {
	Account EditAccountInfos `json:"account"`
	Config  AppConfigModel   `json:"config"`
}

type EditAccountInfos struct {
	Id            int64  `json:"id,omitempty"`
	IdNationality int64  `json:"idNationality,omitempty"`
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"password,omitempty"`
	PhonePrefix   string `json:"phonePrefix,omitempty"`
	Ddd           string `json:"ddd,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Deleted       bool   `json:"deleted,omitempty"`
}
