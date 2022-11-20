package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *AccountServer) Create(c context.Context, req *pb.NewAccountRequest) (*pb.StatusResponse, error) {

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	producer.SendMessage(body, "NEWACCOUNT")

	return &pb.StatusResponse{
		Status: "Success!",
	}, nil
}

func (s *AccountServer) Details(c context.Context, req *pb.Id) (*pb.AccountDetailsResponse, error) {

	dbRes, dbErr := db.AccountDetails(req)
	if dbErr != nil {
		return nil, dbErr
	}

	if dbRes.Account.Id == 0 {
		return nil, errors.New("account was not found")
	}

	return dbRes, nil
}

func (s *AccountServer) Edit(c context.Context, req *pb.EditAccountRequest) (*pb.StatusResponse, error) {

	_, err := s.Details(c, &pb.Id{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	producer.SendMessage(body, "EDITACCOUNT")

	return &pb.StatusResponse{
		Status: "OK",
	}, nil
}

func (s *AccountServer) SoftDelete(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	if req.GetId() <= 0 {
		err := errors.New("id must be greater than zero")
		return nil, err
	}

	if body, err := protojson.Marshal(req); err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		producer.SendMessage(body, "SOFTDELETEACCOUNT")
		return &pb.StatusResponse{Status: "soft delete request appended to queue"}, nil
	}
}

func (s *AccountServer) Restore(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	if req.GetId() <= 0 {
		err := errors.New("id must be greater than zero")
		return nil, err
	}

	if body, err := protojson.Marshal(req); err != nil {
		common.PrintError(err.Error())
		return nil, err
	} else {
		producer.SendMessage(body, "RESTOREACCOUNT")
		return &pb.StatusResponse{Status: "restore account request appended to queue"}, nil
	}
}

func (s *AccountServer) Delete(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	if req.GetId() <= 0 {
		err := errors.New("id must be greater than zero")
		return nil, err
	}

	if body, err := protojson.Marshal(req); err != nil {
		common.PrintError(err.Error())
		return nil, err
	} else {
		producer.SendMessage(body, "DELETEACCOUNT")
		return &pb.StatusResponse{Status: "restore account request appended to queue"}, nil
	}
}
