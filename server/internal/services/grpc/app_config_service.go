package grpc

import (
	"context"
	"errors"
	"strings"

	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/services/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *AppConfigServer) Create(c context.Context, req *pb.NewAppConfigRequest) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdLanguage() <= 0:
		return nil, errors.New("language id must be greater than one")
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than one")
	case strings.EqualFold(req.GetTheme(), ""):
		return nil, errors.New("theme can not be null")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = producer.SendMessage(body, "CREATEAPPCONFIG")
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		Status: "Sended to queue",
	}, nil
}

func (s *AppConfigServer) Update(c context.Context, req *pb.AppConfigDetails) (*pb.StatusResponse, error) {
	switch {
	case req.GetIdConfig() <= 0:
		return nil, errors.New("config id must be greater than zero")
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than zero")
	case req.GetIdLanguage() <= 0:
		return nil, errors.New("language id must be greater than zero")
	case req.GetIdCurrency() <= 0:
		return nil, errors.New("currency id must be greater than zero")
	case strings.EqualFold(req.GetTheme(), ""):
		return nil, errors.New("theme can not be null")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = producer.SendMessage(body, "UPDATEAPPCONFIG")
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		Status: "Sended to queue",
	}, nil
}

func (s *AppConfigServer) Details(c context.Context, req *pb.Id) (*pb.AppConfigDetails, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	}

	res, err := db.AppConfigDetails(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *AppConfigServer) Delete(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	switch {
	case req.GetId() <= 0:
		return nil, errors.New("id must be greater than zero")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = producer.SendMessage(body, "DELETEAPPCONFIG")
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		Status: "Sended to queue",
	}, nil
}
