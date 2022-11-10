package models

type AppConfigModel struct {
	Id                int64  `json:"id" db:"id"`
	Id_language       int64  `json:"id_language" db:"id_language"`
	Id_currency       int64  `json:"id_currency" db:"id_currency"`
	Theme             string `json:"theme" db:"theme"`
	BiometryActivated bool   `json:"biometry_activated" db:"biometry_activated"`
	AlertOnEmail      bool   `json:"receive_alert_on_email" db:"receive_alert_on_email"`
	AlertOnMobile     bool   `json:"receive_alert_on_mobile" db:"receive_alert_on_mobile"`
}

type AppConfigDetails struct {
	IdConfig          int64  `json:"id_config,omitempty"            db:"id_config"`
	AppTheme          string `json:"app_theme,omitempty"          db:"config_theme"`
	BiometryActivated bool   `json:"biometry_activated,omitempty" db:"config_biometry_activated"`
	AlertOnEmail      bool   `json:"alert_on_email,omitempty"     db:"config_alert_on_email"`
	AlertOnMobile     string `json:"alert_on_mobile,omitempty"    db:"config_alert_on_mobile"`
	AppLanguageId     int64  `json:"app_language_id,omitempty"       db:"config_language_id"`
	AppLanguage       string `json:"app_language,omitempty"       db:"config_language"`
	AppLanguageKey    string `json:"app_language_key,omitempty"   db:"config_language_key"`
}
