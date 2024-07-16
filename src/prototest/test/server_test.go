package test

import (
	"GoNotes/src/prototest/person"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestServer(t *testing.T) {

	lin, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))

	if err != nil {
		logrus.Errorf("listen tcp 8080 error %v", err)
	}

	s := grpc.NewServer()

	person.RegisterPersonServiceServer(s, &person.ServiceImpl{})

	logrus.Infof("server1 listening at %v", lin.Addr())

	if err := s.Serve(lin); err != nil {
		logrus.Fatalf("failed to server1 :%v", err)
	}
}
