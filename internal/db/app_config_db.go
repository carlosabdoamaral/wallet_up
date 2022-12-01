package db

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
)

var (
	createConfigQuery = `
	INSERT INTO app_config_tb(id_language, id_currency, theme, biometry_activated, receive_alert_on_email, receive_alert_on_mobile)
	VALUES ($1, $2, $3, $4, $5, $6);
	`

	updateConfigQuery = `
	UPDATE app_config_tb
	SET
		id_language = $1,
		id_currency = $2,
		theme = $3,
		biometry_activated = $4,
		receive_alert_on_email = $5,
		receive_alert_on_mobile = $6
	WHERE
		id = $7;
	`

	configDetailsQuery = `
	SELECT
		id,
		id_currency as currency,
		theme,
		biometry_activated,
		receive_alert_on_email as alert_on_email,
		receive_alert_on_mobile as alert_on_mobile
	FROM app_config_tb
	WHERE id = $1
	`

	deleteConfigQuery = `
		DELETE FROM app_config_tb WHERE id = $1
	`
)

func CreateAppConfig(m *pb.NewAppConfigRequest) error {
	db := common.Database
	_, err := db.Exec(
		createConfigQuery,
		m.GetIdLanguage(),
		m.GetIdCurrency(),
		m.GetTheme(),
		m.GetBiometryActivated(),
		m.GetAlertOnEmail(),
		m.GetAlertOnMobile(),
	)

	return err
}

func UpdateAppConfig(m *pb.AppConfigDetails) error {
	db := common.Database
	_, err := db.Exec(
		updateConfigQuery,
		m.GetIdLanguage(),
		m.GetIdCurrency(),
		m.GetTheme(),
		m.GetBiometryActivated(),
		m.GetAlertOnEmail(),
		m.GetAlertOnMobile(),
		m.GetIdConfig(),
	)

	return err
}

func AppConfigDetails(m *pb.Id) (*pb.AppConfigDetails, error) {
	db := common.Database
	res := &pb.AppConfigDetails{}

	row := db.QueryRow(
		configDetailsQuery,
		m.GetId(),
	)

	err := row.Scan(
		&res.IdConfig,
		&res.IdCurrency,
		&res.Theme,
		&res.BiometryActivated,
		&res.AlertOnEmail,
		&res.AlertOnMobile,
	)
	if err != nil {
		return nil, err
	}

	

	return res, nil
}

func DeleteAppConfig(m *pb.Id) error {
	db := common.Database
	_, err := db.Exec(
		deleteConfigQuery,
		m.GetId(),
	)

	return err
}
