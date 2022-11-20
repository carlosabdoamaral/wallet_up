-- INIT
-- name & language: BRASIL/PT-BR , USA/EN-US
CREATE TABLE IF NOT EXISTS country_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    language_key VARCHAR(10) NOT NULL
);

-- key: BRL, ...
CREATE TABLE IF NOT EXISTS currency_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    key VARCHAR(255)
);

-- theme: DARK
CREATE TABLE IF NOT EXISTS app_config_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    id_language INT REFERENCES country_tb(id) NOT NULL,
    id_currency INT REFERENCES currency_tb(id) NOT NULL,
    theme VARCHAR(255),
    biometry_activated bool,
    receive_alert_on_email bool,
    receive_alert_on_mobile bool
);


CREATE TABLE IF NOT EXISTS user_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    id_nationality INT REFERENCES country_tb(id) NOT NULL DEFAULT 1,
    id_config INT REFERENCES app_config_tb(id) NOT NULL DEFAULT 1,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_prefix VARCHAR(10) NOT NULL,
    ddd VARCHAR(10) NOT NULL,
    phone VARCHAR(10) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    deleted bool DEFAULT false
);

CREATE TABLE IF NOT EXISTS wallet_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    id_user INT REFERENCES user_tb(id),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    deleted bool DEFAULT false
);

-- type : DEPOSIT, ...
CREATE TABLE IF NOT EXISTS transaction_type_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    type VARCHAR(255)
);

-- shared_with : some@email.com
CREATE TABLE IF NOT EXISTS wallet_share_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    id_wallet INT REFERENCES wallet_tb(id),
    shared_with VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS transaction_tb (
    id SERIAL PRIMARY KEY NOT NULL,
    id_user INT REFERENCES user_tb(id) NOT NULL,
    id_wallet INT REFERENCES wallet_tb(id),
    id_currency INT REFERENCES currency_tb(id),
    id_type INT REFERENCES transaction_type_tb(id),
    value float NOT NULL,
    description VARCHAR(255),
    deposited_at timestamp DEFAULT now() NOT NULL
)