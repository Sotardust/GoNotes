package test

import (
	"GoNotes/src/prototest/person"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestClient(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {

		logrus.Errorf("dial error %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logrus.Errorf("conn close error %v", err)

		}
	}(conn)

	info, err := person.NewPersonServiceClient(conn).GetPersonInfo(context.Background(), &person.PersonRequest{Id: "12"})

	if err != nil {
		logrus.Errorf("getPersonInfo failed %v", err)
	}

	logrus.Infof("test info %v", info)
}
