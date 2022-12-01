package utils

import (
	"flag"
	"log"

	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func ConnectToGRPCServer() *grpc.ClientConn {
	flag.Parse()

	conn, err := grpc.Dial(*Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to %v: %v", Addr, err)
	}

	common.GrpcConn = conn
	common.AccountServiceClient = pb.NewAccountServiceClient(conn)
	common.AppConfigServiceClient = pb.NewAppConfigServiceClient(conn)
	common.WalletServiceClient = pb.NewWalletServiceClient(conn)

	return conn
}
