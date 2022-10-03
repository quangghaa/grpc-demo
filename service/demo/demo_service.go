package demo

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	registerPb "github.com/quangghaa/grpc-demo/proto/register"
)

type DemoService struct {
	pingService     *PingService
	registerService *RegisterService
	connService     *ConnectionService
}

func NewDemoService(ps *PingService, rs *RegisterService, cs *ConnectionService) *DemoService {
	return &DemoService{
		pingService:     ps,
		registerService: rs,
		connService:     cs,
	}
}

func (d *DemoService) Register(ctx context.Context, router *runtime.ServeMux, cs *ConnectionService) error {
	// ctx := context.Background()
	d.registerService.Router = router
	d.registerService.Context = ctx
	d.registerService.ConnService = cs
	registerPb.RegisterRegisterHandlerServer(ctx, router, d.registerService)
	return nil
}
