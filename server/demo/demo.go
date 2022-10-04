package register

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quangghaa/grpc-demo/db"
	pingPb "github.com/quangghaa/grpc-demo/proto/ping"
	pb "github.com/quangghaa/grpc-demo/proto/register"
	"github.com/quangghaa/grpc-demo/service/demo"
	connectionService "github.com/quangghaa/grpc-demo/service/demo"
	"github.com/quangghaa/grpc-demo/service/demo/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CONTEXT                   = context.Background()
	PING_SERVICE_ENDPOINT     = "localhost:8001"
	REGISTER_SERVICE_ENDPOINT = "localhost:8002"
)

type register struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func firstLoadEndpoint(ctx context.Context, router *runtime.ServeMux, h *handler.ApiHandler, cs *connectionService.ConnectionService) error {
	res, err := h.Connection_Get_Alive()
	if err != nil {
		return err
	}

	m := new(pb.Connection)
	if err := res.UnmarshalTo(m); err != nil {
		fmt.Println("Error while unmarshal: ", err)
		return err
	}

	opts := []grpc.DialOption{
		grpc.WithBlock(), // Block when calling Dial until the connection is really established
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(m.Endpoint, opts...)
	if err != nil {
		fmt.Println("Cannot Dial to endpoint: ", err.Error())
	}

	err = pingPb.RegisterPingHandler(ctx, router, conn)
	if err != nil {
		log.Fatalln("failed to dial server: ", err)
	}

	// Save to array, to remove later
	connInfo := &connectionService.ConnectionInfo{
		Endpoint: m.Endpoint,
		Conn:     conn,
	}
	cs.Add(connInfo)

	return nil
}

func httpHandlers(listener net.Listener) error {
	// load db
	db, err := db.Init()
	if err != nil {
		fmt.Println("Cannot load database")
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw_router := runtime.NewServeMux()

	cs := demo.NewConnectionService("1")

	handler := handler.NewApiHandler(db)
	ps := demo.NewPingService()
	rs := demo.NewRegisterService(handler)

	go func() {
		rs.Start(8001)
	}()

	// First load endpoint
	fmt.Println("First load endpoint after 2 seconds")
	time.AfterFunc(2*time.Second, func() {
		err := firstLoadEndpoint(ctx, gw_router, handler, cs)
		if err != nil {
			fmt.Println("Error first load: ", err)

		}
	})
	//

	demoService := demo.NewDemoService(db, ps, rs, cs)
	demoService.Register(ctx, gw_router, cs)

	server := &http.Server{
		Handler: gw_router,
	}

	return server.Serve(listener)
}

func GatewayServer(port int) {
	fmt.Println("Start gateway")

	defer fmt.Println("Gateway started")

	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("net.Listen error: %s \n", err)
		return
	}

	defer listener.Close()

	if err := httpHandlers(listener); err != nil {
		log.Fatal(err)
	}

}
