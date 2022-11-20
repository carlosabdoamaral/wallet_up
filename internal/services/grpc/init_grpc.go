package grpc

import (
	"log"
	"net"

	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/grpc"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

type AppConfigServer struct {
	pb.UnimplementedAppConfigServiceServer
}

type WalletServer struct {
	pb.UnimplementedWalletServiceServer
}

func InitServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return
	}

	common.GrpcServer = grpc.NewServer()
	pb.RegisterAccountServiceServer(common.GrpcServer, &AccountServer{})
	pb.RegisterAppConfigServiceServer(common.GrpcServer, &AppConfigServer{})
	pb.RegisterWalletServiceServer(common.GrpcServer, &WalletServer{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := common.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
