package grpc

import (
	"context"
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/services/rabbit/producer"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *CategoryServer) CreateCategory(c context.Context, req *pb.CreateCategoryRequest) (*pb.StatusResponse, error) {
	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "CREATECATEGORY")
	return nil, nil
}

func (s *CategoryServer) EditCategory(c context.Context, req *pb.EditCategoryRequest) (*pb.StatusResponse, error) {
	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "EDITCATEGORY")
	return nil, nil
}

func (s *CategoryServer) CategoryList(c context.Context, req *pb.Id) (*pb.CategoryListResponse, error) {
	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	fmt.Println(body)
	return nil, nil
}

func (s *CategoryServer) DeleteCategory(c context.Context, req *pb.Id) (*pb.StatusResponse, error) {
	body, err := protojson.Marshal(req)
	if err != nil {
		common.PrintError(err.Error())
		return nil, err
	}

	producer.SendMessage(body, "DELETECATEGORY")
	return nil, nil
}
