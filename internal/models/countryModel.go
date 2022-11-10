package models

type CountryModel struct {
	Id       int64 `json:"id" db:"id"`
	Name     int64 `json:"name" db:"name"`
	Language int64 `json:"language" db:"language"`
}
