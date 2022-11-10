package db

import (
	"database/sql"

	"github.com/carlosabdoamaral/wallet_up/common"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(common.DB_DRIVER, common.DB_URL)
	if err != nil {
		return nil, err
	}

	common.Database = db
	return db, nil
}
