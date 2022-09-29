package register

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pingPb "github.com/quangghaa/grpc-demo/proto/ping"
	pb "github.com/quangghaa/grpc-demo/proto/register"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RegisterService struct {
	pb.UnimplementedRegisterServer
}

var (
	CONTEXT               = context.Background()
	PING_SERVICE_ENDPOINT = "localhost:8001"
)

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (*RegisterService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {

	fmt.Println("Start register")
	endpoint := in.Host + ":" + in.Port
	conn, err := grpc.Dial(
		// context.Background(),
		endpoint,
		// grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("Logg")
		log.Fatalln("Failed to dial server: ", err)
	}
	fmt.Println("Pass conn")

	gwmux := runtime.NewServeMux()
	// Register Ping service
	err = pingPb.RegisterPingHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	fmt.Println("Pass regis")

	return &pb.RegisterReply{Message: in.Host + ":" + in.Port}, nil
}

func Start(port int) error {
	// Create a listener on TCP port
	strPort := fmt.Sprint(port)
	lis, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		log.Fatalln("Fail to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Register service to the server
	pb.RegisterRegisterServer(s, &RegisterService{})

	// Serve gRPC server
	log.Printf("Serving gRPC on %s\n", strPort)

	return s.Serve(lis)
}
