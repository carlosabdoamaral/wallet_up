package db

import (
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
	RestoreAccountQuery    = `UPDATE user_tb SET deleted = false WHERE id = $1;`
)

func NewAccount(m *pb.NewAccountRequest) {
	db := common.Database

	_, err := db.Exec(NewAccountQuery, m.IdNationality, m.Firstname, m.Lastname, m.Email, m.Password, m.PhonePrefix, m.Ddd, m.Phone)
	utils.CheckErr(err, false, "Error inserting data")

}

func AccountDetails(m *pb.Id) (*pb.AccountDetailsResponse, error) {
	db := common.Database
	account := &models.AccountDetailsResponse{}

	rows := db.QueryRow(AccountDetailsQuery, m.GetId())
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

	return &pb.AccountDetailsResponse{
		Account: &pb.AccountDetails{
			Id:              account.Account.Id,
			Firstname:       account.Account.Firstname,
			Lastname:        account.Account.Lastname,
			Email:           account.Account.Email,
			Password:        account.Account.Password,
			PhonePrefix:     account.Account.PhonePrefix,
			Ddd:             account.Account.Ddd,
			Phone:           account.Account.Phone,
			Deleted:         account.Account.Deleted,
			IdNationality:   account.Account.IdNationality,
			NationalityName: account.Account.NationalityName,
			NationalityKey:  account.Account.NationalityKey,
		},
		Config: &pb.AppConfigDetails{
			IdConfig:          account.Config.IdConfig,
			AppTheme:          account.Config.AppTheme,
			BiometryActivated: account.Config.BiometryActivated,
			AlertOnEmail:      account.Config.AlertOnEmail,
			AlertOnMobile:     account.Config.AlertOnMobile,
			AppLanguageId:     account.Config.AppLanguageId,
			AppLanguage:       account.Config.AppLanguage,
			AppLanguageKey:    account.Config.AppLanguageKey,
		},
	}, nil
}

func EditAccount(m *pb.EditAccountRequest) {
	db := common.Database
	account := m.Account
	config := m.Config

	_, err := db.Exec(
		EditAccountQuery,
		account.GetIdNationality(),
		account.GetFirstname(),
		account.GetLastname(),
		account.GetEmail(),
		account.GetPassword(),
		account.GetPhonePrefix(),
		account.GetDdd(),
		account.GetPhone(),
		account.GetDeleted(),
		m.GetId(),
	)
	if err != nil {
		return
	}

	_, err = db.Exec(
		EditConfigQuery,
		config.GetIdLanguage(),
		config.GetIdCurrency(),
		config.GetTheme(),
		config.GetBiometryActivated(),
		config.GetAlertOnEmail(),
		config.GetAlertOnMobile(),
		m.GetId(),
	)
	if err != nil {
		return
	}
}

func SoftDeleteAccount(m *pb.Id) error {
	db := common.Database
	_, err := db.Exec(SoftDeleteAccountQuery, m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func RestoreAccount(m *pb.Id) error {
	db := common.Database
	_, err := db.Exec(RestoreAccountQuery, m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}
