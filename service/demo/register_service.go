package demo

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pingPb "github.com/quangghaa/grpc-demo/proto/ping"
	pb "github.com/quangghaa/grpc-demo/proto/register"
	"github.com/quangghaa/grpc-demo/service/demo/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RegisterService struct {
	pb.UnimplementedRegisterServer
	Router  *runtime.ServeMux
	Context context.Context

	ConnService *ConnectionService

	handler *handler.ApiHandler
}

var (
	CONTEXT               = context.Background()
	PING_SERVICE_ENDPOINT = "localhost:8001"
)

func NewRegisterService(h *handler.ApiHandler) *RegisterService {
	return &RegisterService{
		handler: h,
	}
}

func (r *RegisterService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	endpoint := in.Host + ":" + in.Port

	opts := []grpc.DialOption{
		grpc.WithBlock(), // Block when calling Dial until the connection is really established
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		fmt.Println("Cannot Dial to endpoint: %s ", err.Error())
		return &pb.RegisterReply{Message: "Cannot dial", Conns: []string{}}, nil
	}

	err = pingPb.RegisterPingHandler(r.Context, r.Router, conn)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	// Save to db
	connObj := &pb.Connection{
		ServiceName: "Ping",
		Endpoint:    endpoint,
	}

	_, err = r.handler.Connection_Create(connObj)
	if err != nil {
		fmt.Println("Cannot save to db: %s ", err.Error())
	}

	// Save to array, to remove later
	r.ConnService.Add(conn)

	return &pb.RegisterReply{Message: in.Host + ":" + in.Port, Conns: []string{}}, nil
}

func (r *RegisterService) CheckConnection(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	conns := []string{}
	for _, conn := range r.ConnService.ConnPool {
		conns = append(conns, fmt.Sprint(conn.GetState()))
	}

	if len(conns) == 0 {
		return &pb.RegisterReply{Message: "Empty", Conns: conns}, nil
	}

	return &pb.RegisterReply{Message: "OK", Conns: conns}, nil
}

func (c *RegisterService) ScanConnection(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	fmt.Println("Start scanning ...")
	go func() {
		count := 1
		for {
			count++
			if len(c.ConnService.ConnPool) > 1 {
				for i := 0; i < len(c.ConnService.ConnPool)-1; i++ {
					err := c.ConnService.ConnPool[i].Close()
					if err != nil {
						fmt.Println("Error while close connection: ", err)
					} else {
						fmt.Println("Close connection:")
						fmt.Println(c.ConnService.ConnPool[i])
					}
				}
			}
			l := len(c.ConnService.ConnPool) - 1
			c.ConnService.ConnPool = c.ConnService.ConnPool[l:]

			time.Sleep(10 * time.Second)

		}
	}()
	return &pb.RegisterReply{Message: "OK"}, nil
}

func (r *RegisterService) Start(port int) error {
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
	pb.RegisterRegisterServer(s, r)

	// Serve gRPC server
	log.Printf("Serving gRPC on %s\n", strPort)

	return s.Serve(lis)
}
