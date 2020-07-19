package server_b

import (
	"context"

	"github.com/phev8/go_commons/pkg/api/shared"
	"github.com/phev8/go_service_B/pkg/api"
)

func (s *serviceBServer) GetData(ctx context.Context, req *shared.RequestObject) (*api.DataObject, error) {
	s.counter += 1

	return &api.DataObject{
		Counter: s.counter,
		CommonObject: &shared.CommonObject{
			Id:   23,
			Name: "test",
		},
	}, nil
}
