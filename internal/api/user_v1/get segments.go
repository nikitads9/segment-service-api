package user_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
)

func (i *Implementation) GetSegments(ctx context.Context, req *desc.GetSegmentsRequest) (*desc.GetSegmentsResponse, error) {
	slugs, err := i.userService.GetSegments(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetSegmentsResponse{
		Slugs: slugs,
	}, nil
}
