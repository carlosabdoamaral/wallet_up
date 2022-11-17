package grpc

import (
	"context"

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
		Status: "Good to go myman",
	}, nil
}

func (s *AccountServer) Details(c context.Context, req *pb.Id) (*pb.AccountDetailsResponse, error) {
	common.PrintStartMethod("[GRPC] AccountDetails!")

	dbRes, dbErr := db.AccountDetails(req)
	if dbErr != nil {
		common.PrintError(dbErr.Error())
		return nil, dbErr
	}

	return dbRes, nil
}
