package register

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quangghaa/grpc-demo/service/demo"
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw_router := runtime.NewServeMux()

	ps := demo.NewPingService(gw_router)
	rs := demo.NewRegisterService(gw_router, ctx)
	fmt.Printf("First: %v\n", rs.Router)
	go func() {
		rs.Start(8001)
	}()

	demoService := demo.NewDemoService(ps, rs)
	demoService.Register(ctx, gw_router)

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

	// ctx := context.Background()

	// connR, err := grpc.DialContext(
	// 	ctx,
	// 	REGISTER_SERVICE_ENDPOINT,
	// 	grpc.WithBlock(),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalln("Failed to dial REGISTER server: ", err)
	// }

	// gwmux := runtime.NewServeMux()

	// err = registerPb.RegisterRegisterHandler(ctx, gwmux, connR)
	// if err != nil {
	// 	log.Fatalln("Failed to register REGISTER gateway: ", err)
	// }

	// gwServer := &http.Server{
	// 	Addr:    "localhost:" + fmt.Sprint(port),
	// 	Handler: gwmux,
	// }

	// log.Println("Serving gRPC-Gateway on http://localhost:" + fmt.Sprint(port))
	// // go func() {
	// // 	log.Fatalln(gwServer.ListenAndServe())
	// // }()
	// gwServer.ListenAndServe()
}
