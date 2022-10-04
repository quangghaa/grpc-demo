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
		return nil, fmt.Errorf("existed endpoint")
	}
	conn = mcon.Connection{
		ServiceName: req.ServiceName,
		Endpoint:    req.Endpoint,
		Status:      req.Status,
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

func (s *ApiHandler) Connection_Edit(req *pb.Connection) (*anypb.Any, error) {
	conn := mcon.Connection{}
	if s.db.Model(mcon.Connection{}).Where("endpoint", req.Endpoint).First(&conn).RowsAffected != 1 {
		return nil, fmt.Errorf("record does not exist")
	}
	conn = mcon.Connection{
		ServiceName: req.ServiceName,
		Endpoint:    req.Endpoint,
		Status:      req.Status,
	}
	res := s.db.Model(&mcon.Connection{}).Where("endpoint = ?", req.Endpoint).Update("status", 0)
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

func (s *ApiHandler) Connection_Get_Alive() (*anypb.Any, error) {
	conn := mcon.Connection{}
	res := s.db.Model(mcon.Connection{}).Where("status", 1).Last(&conn)
	if res.Error != nil {
		return nil, res.Error
	}
	obj := &pb.Connection{
		Id:          conn.ID,
		ServiceName: conn.ServiceName,
		Endpoint:    conn.Endpoint,
	}
	v, err := anypb.New(obj)
	if err != nil {
		return nil, err
	}
	return v, nil
}
