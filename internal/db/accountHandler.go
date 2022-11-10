package db

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

var (
	NewAccountQuery = "INSERT INTO user_tb(id_nationality, first_name, last_name, email, password, phone_prefix, ddd, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
)

func NewAccount(m *models.NewAccountRequest) {
	db := common.Database

	insert, err := db.Prepare(NewAccountQuery)
	utils.CheckErr(err, false, "Error preparing query")

	_, err = insert.Exec(m.IdNationality, m.Firstname, m.Lastname, m.Email, m.Password, m.PhonePrefix, m.Ddd, m.Phone)
	utils.CheckErr(err, false, "Error inserting data")

	common.PrintSuccess("Account created!")
}
