package grpc

import (
	"context"
	"errors"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *AccountServer) Create(c context.Context, req *pb.NewAccountRequest) (*pb.StatusResponse, error) {
	common.PrintStartMethod("[GRPC] CreateAccount!")

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "NEWACCOUNT")

	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *AccountServer) Details(c context.Context, req *pb.Id) (*pb.AccountDetailsResponse, error) {
	common.PrintStartMethod("[GRPC] AccountDetails!")

	dbRes, dbErr := db.AccountDetails(req)
	if dbErr != nil {
		common.PrintError(dbErr.Error())
		return nil, dbErr
	}

	if dbRes.Account.Id == 0 {
		return nil, errors.New("account was not found")
	}

	return dbRes, nil
}

func (s *AccountServer) Edit(c context.Context, req *pb.EditAccountRequest) (*pb.StatusResponse, error) {
	common.PrintStartMethod("[GRPC] EditAccount")

	_, err := s.Details(c, &pb.Id{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "EDITACCOUNT")

	return &pb.StatusResponse{
		Status: "OK",
	}, nil
}
