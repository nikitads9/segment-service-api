package user_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
)

func (i *Implementation) GetUserHistoryCsv(ctx context.Context, req *desc.GetUserHistoryCsvRequest) (*desc.GetUserHistoryCsvResponse, error) {
	buffer, err := i.userService.GetUserHistoryCsv(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetUserHistoryCsvResponse{
		Chunk: buffer.Bytes(),
	}, nil
}
