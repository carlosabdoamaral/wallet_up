package grpc

import (
	"context"
	"errors"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/services/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *OperationServer) Deposit(c context.Context, req *pb.TransactionRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than zero")
	case req.GetIdType() <= 0:
		return nil, errors.New("type id must be greater than zero")
	case req.GetIdUser() <= 0:
		return nil, errors.New("user id must be greater than zero")
	case req.GetIdWallet() <= 0:
		return nil, errors.New("wallet id must be greater than zero")
	case req.GetValue() <= 0:
		return nil, errors.New("value must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "DEPOSITWALLET")
	return nil, nil
}

func (s *OperationServer) Withdraw(c context.Context, req *pb.TransactionRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than zero")
	case req.GetIdType() <= 0:
		return nil, errors.New("type id must be greater than zero")
	case req.GetIdUser() <= 0:
		return nil, errors.New("user id must be greater than zero")
	case req.GetIdWallet() <= 0:
		return nil, errors.New("wallet id must be greater than zero")
	case req.GetValue() <= 0:
		return nil, errors.New("value must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "WITHDRAWWALLET")
	return nil, nil
}

func (s *OperationServer) DeleteTransaction(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "DELETETRANSACTION")
	return nil, nil
}

func (s *OperationServer) EditTransaction(c context.Context, req *pb.EditTransactionRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdTransaction() <= 0:
		return nil, errors.New("transaction id must be greater than zero")
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than zero")
	case req.GetIdType() <= 0:
		return nil, errors.New("type id must be greater than zero")
	case req.GetValue() <= 0:
		return nil, errors.New("value must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "EDITTRANSACTION")
	return nil, nil
}
