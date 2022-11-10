package db

import (
	"database/sql"
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

var (
	NewAccountQuery     = "INSERT INTO user_tb(id_nationality, first_name, last_name, email, password, phone_prefix, ddd, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	AccountDetailsQuery = `SELECT
		user_tb.id,
		user_tb.first_name,
		user_tb.last_name,
		user_tb.email,
		user_tb.password,
		user_tb.phone_prefix,
		user_tb.ddd,
		user_tb.phone,
		user_tb.deleted,
		user_tb.id_nationality,
		ct_a.name as nationality,
		ct_a.language_key as nationality_key,
		user_tb.id_config,
		cf.theme as config_theme,
		cf.biometry_activated as config_biometry_activated,
		cf.receive_alert_on_email as config_alert_on_email,
		cf.receive_alert_on_mobile as config_alert_on_mobile,
		ct_b.name as config_language,
		ct_b.language_key as config_language_key
	FROM
		user_tb
	
	INNER JOIN country_tb ct_a on ct_a.id = user_tb.id_nationality
	INNER JOIN app_config_tb cf on cf.id = user_tb.id_config
	INNER JOIN country_tb ct_b on ct_b.id = cf.id_language
	
	WHERE
		user_tb.id = $1
	`
)

func NewAccount(m *models.NewAccountRequest) {
	db := common.Database

	_, err := db.Exec(NewAccountQuery, m.IdNationality, m.Firstname, m.Lastname, m.Email, m.Password, m.PhonePrefix, m.Ddd, m.Phone)
	utils.CheckErr(err, false, "Error inserting data")

	common.PrintSuccess("Account created!")
}

func AccountDetails(m *models.AccountDetailsRequest) (*models.AccountDetails, error) {
	db := common.Database
	result := &models.AccountDetails{}

	rows, err := db.Query(AccountDetailsQuery, m.AccountId)
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
		return nil, err
	}

	for rows.Next() {
		account := &models.AccountDetails{}
		rows.Scan(
			&account.Id,
			&account.Firstname,
			&account.Lastname,
			&account.Email,
			&account.Password,
			&account.PhonePrefix,
			&account.Ddd,
			&account.Phone,
			&account.Deleted,
			&account.IdNationality,
			&account.NationalityName,
			&account.NationalityKey,
			&account.IdConfig,
			&account.AppTheme,
			&account.BiometryActivated,
			&account.AlertOnEmail,
			&account.AlertOnMobile,
			&account.AppLanguage,
			&account.AppLanguageKey,
		)

		result = account
	}

	return result, nil
}
