package prototest

import (
	"GoNotes/src/prototest/person"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestPersonSerial(t *testing.T) {

	var p = person.Person{Age: 12, Name: "小代", Hobbies: []string{"dai", "测试"}}

	marshal, err := proto.Marshal(&p)
	if err != nil {
		logrus.Errorf("person marshal error %v", err)
		return
	}

	logrus.Infof("marshal %v", marshal)

	p2 := &person.Person{}

	err = proto.Unmarshal(marshal, p2)
	if err != nil {
		logrus.Errorf("person unmarshal error %v", err)
		return
	}

	logrus.Infof("p2 %v", p2)

	if p.String() != p2.String() {
		logrus.Infof("p %v , p2 %v", p, p2)
	}
}
