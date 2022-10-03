package demo

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type ConnectionService struct {
	Id       string
	ConnPool []*grpc.ClientConn
}

func NewConnectionService(id string) *ConnectionService {
	return &ConnectionService{
		Id: id,
	}
}

func (c *ConnectionService) Add(conn *grpc.ClientConn) error {
	fmt.Println("Second pool id: ", c.Id)
	c.ConnPool = append(c.ConnPool, conn)
	fmt.Println("Add success: ", len(c.ConnPool))
	return nil
}

func (c *ConnectionService) Check() error {
	fmt.Println("Third pool id: ", c.Id)
	if len(c.ConnPool) == 0 {
		fmt.Println("No connection established")
		return nil
	}
	fmt.Println("Established connections: ")
	for _, c := range c.ConnPool {
		fmt.Println(fmt.Sprint(c))
	}
	return nil
}

func (c *ConnectionService) Scan() error {
	fmt.Println("Start scanning ...")
	i := 1
	for {
		fmt.Println("Scan >> %d", i)
		i++

		if i == 3 || i == 9 || i == 12 {
			fmt.Println("Going to remove after 5 seconds")
			time.AfterFunc(5*time.Second, func() {
				fmt.Println("REMOVED !!!")
			})
		}

		time.Sleep(5 * time.Second)

	}
}
