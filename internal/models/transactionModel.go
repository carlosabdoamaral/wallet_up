package models

import "time"

type CurrencyModel struct {
	Id  int64 `json:"id" db:"id"`
	Key int64 `json:"key" db:"key"`
}

type TransactionTypeModel struct {
	Id   int64 `json:"id" db:"id"`
	Type int64 `json:"type" db:"type"`
}

type TransactionModel struct {
	Id           int64     `json:"id" db:"id"`
	Id_user      int64     `json:"id_user" db:"id_user"`
	Id_wallet    int64     `json:"id_wallet" db:"id_wallet"`
	Id_currency  int64     `json:"id_currency" db:"id_currency"`
	Id_type      int64     `json:"id_type" db:"id_type"`
	Value        float64   `json:"value" db:"value"`
	Deposited_at time.Time `json:"deposited_at" db:"deposited_at"`
	Description  string    `json:"description" db:"description"`
}
