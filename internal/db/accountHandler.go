package db

import (
	"database/sql"
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
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
		ct_b.id as config_language_id,
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

	EditAccountQuery = `
	UPDATE user_tb
	SET id_nationality = $1, first_name = $2, last_name  = $3, email = $4, password = $5, phone_prefix = $6, ddd = $7, phone = $8, updated_at = now(), deleted = $9
	WHERE user_tb.id = $10;`

	EditConfigQuery = `
	UPDATE app_config_tb
	SET id_language = $1, id_currency = $2, theme = $3, biometry_activated = $4, receive_alert_on_email = $5, receive_alert_on_mobile = $6
	WHERE app_config_tb.id = $7;
	`

	SoftDeleteAccountQuery = `UPDATE user_tb SET deleted = true WHERE id = $1;`
)

func NewAccount(m *pb.NewAccountRequest) {
	db := common.Database

	_, err := db.Exec(NewAccountQuery, m.IdNationality, m.Firstname, m.Lastname, m.Email, m.Password, m.PhonePrefix, m.Ddd, m.Phone)
	utils.CheckErr(err, false, "Error inserting data")

	common.PrintSuccess("Account created!")
}

func AccountDetails(m *models.AccountId) (*models.AccountDetails, error) {
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
			&account.Account.Id,
			&account.Account.Firstname,
			&account.Account.Lastname,
			&account.Account.Email,
			&account.Account.Password,
			&account.Account.PhonePrefix,
			&account.Account.Ddd,
			&account.Account.Phone,
			&account.Account.Deleted,
			&account.Account.IdNationality,
			&account.Account.NationalityName,
			&account.Account.NationalityKey,
			&account.Config.IdConfig,
			&account.Config.AppTheme,
			&account.Config.BiometryActivated,
			&account.Config.AlertOnEmail,
			&account.Config.AlertOnMobile,
			&account.Config.AppLanguageId,
			&account.Config.AppLanguage,
			&account.Config.AppLanguageKey,
		)

		result = account
	}

	return result, nil
}

func EditAccount(m *models.EditAccountRequest) {
	db := common.Database
	account := m.Account
	config := m.Config

	_, err := db.Exec(
		EditAccountQuery,
		account.IdNationality,
		account.Firstname,
		account.Lastname,
		account.Email,
		account.Password,
		account.PhonePrefix,
		account.Ddd,
		account.Phone,
		account.Deleted,
		account.Id,
	)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	_, err = db.Exec(
		EditConfigQuery,
		config.Id_language,
		config.Id_currency,
		config.Theme,
		config.BiometryActivated,
		config.AlertOnEmail,
		config.AlertOnMobile,
		config.Id,
	)
	if err != nil {
		common.PrintError(err.Error())
		return
	}
}

func SoftDeleteAccount(m *models.AccountId) {
	common.PrintInfo("[DB] SoftDeleteAccount")
	db := common.Database

	_, err := db.Exec(SoftDeleteAccountQuery, m.AccountId)
	if err != nil {
		common.PrintError(err.Error())
	}
}
