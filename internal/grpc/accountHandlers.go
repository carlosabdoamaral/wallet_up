package grpc

import (
	"context"

	"github.com/carlosabdoamaral/wallet_up/common"
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
