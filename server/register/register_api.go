package register

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	registerPb "github.com/quangghaa/grpc-demo/proto/register"
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

func GatewayServer(port int) error {

	ctx := context.Background()

	connR, err := grpc.DialContext(
		ctx,
		REGISTER_SERVICE_ENDPOINT,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial REGISTER server: ", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Ping service
	// err = pingPb.RegisterPingHandler(ctx, gwmux, connP)
	// if err != nil {
	// 	log.Fatalln("Failed to register PING gateway: ", err)
	// }

	// Register Register service
	err = registerPb.RegisterRegisterHandler(ctx, gwmux, connR)
	if err != nil {
		log.Fatalln("Failed to register REGISTER gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    "localhost:" + fmt.Sprint(port),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://localhost:" + fmt.Sprint(port))
	// go func() {
	// 	log.Fatalln(gwServer.ListenAndServe())
	// }()
	return gwServer.ListenAndServe()
}
