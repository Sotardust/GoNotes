package person

import (
	"context"
)

type ServiceImpl struct {
}

func (_ *ServiceImpl) GetPersonInfo(ctx context.Context, request *PersonRequest) (*PersonResponse, error) {
	response := &PersonResponse{
		Person: &Person{
			Name:    "小代",
			Age:     23,
			Hobbies: []string{"你好", "我也好"},
		},
	}

	return response, nil
}
