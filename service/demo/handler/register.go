package handler

import (
	"fmt"

	mcon "github.com/quangghaa/grpc-demo/models/connection"
	pb "github.com/quangghaa/grpc-demo/proto/register"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *ApiHandler) Connection_Create(req *pb.Connection) (*anypb.Any, error) {
	conn := mcon.Connection{}
	if s.db.Model(mcon.Connection{}).Where("endpoint", req.Endpoint).First(&conn).RowsAffected > 0 {
		return nil, fmt.Errorf("Existed endpoint")
	}
	conn = mcon.Connection{
		ServiceName: req.ServiceName,
		Endpoint:    req.Endpoint,
	}
	res := s.db.Create(&conn)
	if res.Error != nil {
		return nil, res.Error
	}
	req.Id = conn.ID
	item, err := anypb.New(req)
	if err != nil {
		return nil, err
	}
	return item, nil
}
