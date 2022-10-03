package demo

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	pb "github.com/quangghaa/grpc-demo/proto/ping"
	"google.golang.org/grpc"
)

type PingService struct {
	pb.UnimplementedPingServer
}

var ()

func NewPingService() *PingService {
	return &PingService{}
}

func (*PingService) PingMe(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	fmt.Println("Third ctx: ", ctx)
	return &pb.PingReply{Message: "PONG"}, nil
}

func (*PingService) SlowPing(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	duration, err := strconv.Atoi(in.Delay)
	if err != nil {
		log.Fatalln("Parameters not valid: ", err)
	}
	time.Sleep(time.Duration(duration) * time.Second)

	return &pb.PingReply{Message: "SLOW PONG " + in.Delay + " seconds"}, nil
}

func (*PingService) Start(port int) error {
	fmt.Println("Start PING serivce ...")
	// Create a listener on TCP port
	strPort := fmt.Sprint(port)
	lis, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		log.Fatalln("Fail to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterPingServer(s, &PingService{})

	// Serve gRPC server
	log.Printf("Serving gRPC on %s\n", strPort)
	return s.Serve(lis)
}
