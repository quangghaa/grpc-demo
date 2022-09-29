package demo

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	registerPb "github.com/quangghaa/grpc-demo/proto/register"
)

type DemoService struct {
	pingService     *PingService
	registerService *RegisterService
}

func NewDemoService(ps *PingService, rs *RegisterService) *DemoService {
	fmt.Printf("Second: %v\n", rs.Router)
	return &DemoService{
		pingService:     ps,
		registerService: rs,
	}
}

func (d *DemoService) Register(ctx context.Context, router *runtime.ServeMux) error {
	// ctx := context.Background()
	d.registerService.Router = router
	d.registerService.Context = ctx
	registerPb.RegisterRegisterHandlerServer(ctx, router, d.registerService)
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// err := registerPb.RegisterRegisterHandlerFromEndpoint(ctx, router, "localhost:8001", opts)
	// if err != nil {
	// 	fmt.Println("ERROR dialing ... ", err)
	// }
	return nil
}
