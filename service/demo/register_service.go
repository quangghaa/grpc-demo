package demo

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
	Router  *runtime.ServeMux
	Context context.Context
}

var (
	CONTEXT               = context.Background()
	PING_SERVICE_ENDPOINT = "localhost:8001"
)

func NewRegisterService(router *runtime.ServeMux, ctx context.Context) *RegisterService {
	return &RegisterService{
		Router:  router,
		Context: ctx,
	}
}

func (r *RegisterService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {

	fmt.Println("IN REGISTER SERVICE >>>>>>>>.")
	endpoint := in.Host + ":" + in.Port
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	fmt.Println(endpoint)
	fmt.Printf("Third: %v\n", r)

	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}

	err = pingPb.RegisterPingHandler(r.Context, r.Router, conn)

	// err := registerPb.RegisterRegisterHandlerFromEndpoint(r.Context, r.Router, endpoint, opts)
	if err != nil {
		fmt.Printf("CHECK error: %s\n", err)
		// log.Fatalln("Failed to dial server: ", err)
	}

	return &pb.RegisterReply{Message: in.Host + ":" + in.Port}, nil
}

func (r *RegisterService) Start(port int) error {
	fmt.Printf("Fourd: %v\n", r)
	fmt.Println("Start REGISTER serivce ...")
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
