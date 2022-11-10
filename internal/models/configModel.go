package models

type AppConfigModel struct {
	Id                      int64  `json:"id" db:"id"`
	Id_language             int64  `json:"id_language" db:"id_language"`
	Id_currency             int64  `json:"id_currency" db:"id_currency"`
	Theme                   string `json:"theme" db:"theme"`
	Biometry_activated      bool   `json:"biometry_activated" db:"biometry_activated"`
	Receive_alert_on_email  bool   `json:"receive_alert_on_email" db:"receive_alert_on_email"`
	Receive_alert_on_mobile bool   `json:"receive_alert_on_mobile" db:"receive_alert_on_mobile"`
}
