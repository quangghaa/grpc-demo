package register

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quangghaa/grpc-demo/db"
	"github.com/quangghaa/grpc-demo/service/demo"
	"github.com/quangghaa/grpc-demo/service/demo/handler"
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
