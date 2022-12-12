package db

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
)

func Deposit(m *pb.TransactionRequest) error {
	db := common.Database
	query := `INSERT INTO transaction_tb(id_user, id_wallet, id_currency, id_type, value, description) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := db.Exec(query, m.GetIdUser(), m.GetIdWallet(), m.GetIdCurrency(), m.GetIdType(), m.GetValue(), m.GetDescription())
	if err != nil {
		common.PrintFatal(err.Error())
		return err
	}

	return nil
}

func Withdraw(m *pb.TransactionRequest) error {
	db := common.Database
	query := `INSERT INTO transaction_tb(id_user, id_wallet, id_currency, id_type, value, description) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := db.Exec(query, m.GetIdUser(), m.GetIdWallet(), m.GetIdCurrency(), m.GetIdType(), m.GetValue(), m.GetDescription())
	if err != nil {
		common.PrintFatal(err.Error())
		return err
	}

	return nil

}

func DeleteTransaction(m *pb.Id) error {
	db := common.Database
	query := `DELETE FROM transaction_tb WHERE id = $1;`
	_, err := db.Exec(query, m.GetId())
	if err != nil {
		common.PrintFatal(err.Error())
		return err
	}

	return nil
}

func EditTransaction(m *pb.EditTransactionRequest) error {
	db := common.Database
	query := `UPDATE transaction_tb SET id_currency = $2,  id_type = $3, value = $4, description = $5 WHERE id = $6;`
	_, err := db.Exec(query, m.GetIdCurrency(), m.GetIdType(), m.GetValue(), m.GetDescription(), m.GetIdTransaction())
	if err != nil {
		common.PrintFatal(err.Error())
		return err
	}

	return nil
}
