package db

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
)

func CreateWallet(m *pb.CreateWalletRequest) error {
	db := common.Database
	query := `INSERT INTO wallet_tb(id_user, name, description) VALUES ($1, $2, $3);`
	_, err := db.Exec(query, m.GetIdUser(), m.GetName(), m.GetDescription())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func EditWallet(m *pb.EditWalletRequest) error {
	db := common.Database
	query := `UPDATE wallet_tb SET name = $1, description = $2 WHERE id = $3;`
	_, err := db.Exec(query, m.GetName(), m.GetDescription(), m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func DeleteWallet(m *pb.Id) error {
	db := common.Database
	query := `
	DELETE FROM wallet_share_tb WHERE id_wallet = $1;
	DELETE FROM transaction_tb WHERE id_wallet = $2;
	DELETE FROM wallet_tb WHERE id = $3;
	`
	_, err := db.Exec(query, m.GetId(), m.GetId(), m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func ShareWallet(m *pb.ShareWalletRequest) error {
	db := common.Database
	query := `INSERT INTO wallet_share_tb(id_wallet, shared_with) VALUES ($1, $2);`
	_, err := db.Exec(query, m.GetIdWallet(), m.GetEmail())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func UnShareWallet(m *pb.UnShareWalletRequest) error {
	db := common.Database
	query := `DELETE FROM wallet_share_tb WHERE upper(shared_with) = upper($1) and id_wallet = $2`
	_, err := db.Exec(query, m.GetEmail(), m.GetIdWallet())
	if err != nil {
		common.PrintError(err.Error())
		return err
	}

	return nil
}

func WalletDetails(m *pb.Id) (*pb.WalletDetailsResponse, error) {
	db := common.Database
	res := &pb.WalletDetailsResponse{}

	// MARK: TRANSACTIONS
	query := `
	SELECT tt.id, tt.id_user, ct.id as currency_id, ct.key as currency_key, ttt.type as transaction_type, tt.value, tt.description, tt.deposited_at
	FROM transaction_tb tt
	INNER JOIN currency_tb ct on ct.id = tt.id_currency
	INNER JOIN transaction_type_tb ttt on ttt.id = tt.id_type
	WHERE id_wallet = $1`

	rows, err := db.Query(query, m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	for rows.Next() {
		t := &pb.TransactionDetails{}
		rows.Scan(t.IdTransaction, t.IdUser, t.IdCurrency, t.CurrencyKey, t.TransactionType, t.Value, t.Description, t.DepositedAt)
		res.Transactions = append(res.Transactions, t)
	}

	// MARK: SHARED CONTACTS
	query = `SELECT id, shared_with FROM wallet_share_tb wst WHERE wst.id_wallet = $1;`
	rows, err = db.Query(query, m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	for rows.Next() {
		t := &pb.SharedWith{}
		rows.Scan(t.IdSharedWith, t.Email)
		res.SharedWithList = append(res.SharedWithList, t)
	}

	// MARK: WALLET DETAILS
	query = `SELECT id, id_user, name, description, created_at, updated_at, deleted FROM wallet_tb WHERE id = $1;`
	rows, err = db.Query(query, m.GetId())
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	for rows.Next() {
		t := &pb.WalletDetails{}
		rows.Scan(t.Id, t.IdUser, t.Name, t.Description, t.CreatedAt, t.UpdatedAt, t.Deleted)
		res.WalletDetails = t
	}

	return res, nil
}
