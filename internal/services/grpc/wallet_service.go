package grpc

import (
	"context"
	"errors"
	"strings"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/services/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *WalletServer) Create(c context.Context, req *pb.CreateWalletRequest) (*pb.StatusResponse, error) {
	switch {
	case strings.EqualFold(req.Name, ""):
		return nil, errors.New("wallet name can not be empty")
	case strings.EqualFold(req.Description, ""):
		return nil, errors.New("wallet description can not be empty")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "CREATEWALLET")
	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *WalletServer) Details(c context.Context, req *pb.Id) (*pb.WalletDetailsResponse, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	}

	return db.WalletDetails(req)
}

func (s *WalletServer) Edit(c context.Context, req *pb.EditWalletRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	case strings.EqualFold(req.GetName(), ""):
		return nil, errors.New("wallet name can not be empty")
	case strings.EqualFold(req.GetDescription(), ""):
		return nil, errors.New("wallet description can not be empty")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "EDITWALLET")
	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *WalletServer) Delete(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "DELETEWALLET")
	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *WalletServer) Share(c context.Context, req *pb.ShareWalletRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdWallet() <= 0:
		return nil, errors.New("wallet id must be greater than zero")
	case strings.EqualFold(req.GetEmail(), ""):
		return nil, errors.New("email can not be empty")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "SHAREWALLET")
	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *WalletServer) UnShare(c context.Context, req *pb.UnShareWalletRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdWallet() <= 0:
		return nil, errors.New("wallet id must be greater than zero")
	case strings.EqualFold(req.GetEmail(), ""):
		return nil, errors.New("email can not be empty")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "UNSHAREWALLET")
	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}
