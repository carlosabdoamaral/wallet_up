package models

import "time"

type WalletModel struct {
	Id          int64     `json:"id" db:"id"`
	IdUser      int64     `json:"id_user" db:"id_user"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" db:"UpdatedAt"`
	Deleted     bool      `json:"deleted" db:"deleted"`
}

type WalletTransactionDetails struct {
	Id          int64     `json:"user_id"      db:"user_id"`
	WalletId    int64     `json:"wallet_id"    db:"wallet_id"`
	Type        string    `json:"type"         db:"type"`
	Value       float64   `json:"value"         db:"value"`
	Currency    string    `json:"currency"     db:"currency"`
	Description string    `json:"description"  db:"description"`
	DepositedAt time.Time `json:"deposited_at" db:"deposited_at"`
}

type WalletShareModel struct {
	Id         int64 `json:"id" db:"id"`
	IdWallet   int64 `json:"id_wallet" db:"id_wallet"`
	SharedWith int64 `json:"shared_with" db:"typshared_with"`
}
