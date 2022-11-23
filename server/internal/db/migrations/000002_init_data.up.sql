INSERT INTO currency_tb(key) VALUES ('BRL'); -- Real
INSERT INTO currency_tb(key) VALUES ('USD'); -- Dolar
INSERT INTO currency_tb(key) VALUES ('EUR'); -- Euro
INSERT INTO currency_tb(key) VALUES ('XBT'); -- Bitcoin

INSERT INTO country_tb(name, language_key) VALUES ('BRAZIL', 'PT-BR');        -- Brasil
INSERT INTO country_tb(name, language_key) VALUES ('UNITED STATES', 'EN-US'); -- Estados Unidos

INSERT INTO transaction_type_tb(type) VALUES ('PROFIT');
INSERT INTO transaction_type_tb(type) VALUES ('EXPENSE');
INSERT INTO transaction_type_tb(type) VALUES ('KEPT');
INSERT INTO transaction_type_tb(type) VALUES ('INVESTED');

INSERT INTO app_config_tb(id_language, id_currency, theme, biometry_activated, receive_alert_on_email, receive_alert_on_mobile)
VALUES (1, 1, 'LIGHT', false, false, false);