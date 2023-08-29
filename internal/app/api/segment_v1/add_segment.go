package segment_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/segment_api"
)

func (i *Implementation) AddSegment(ctx context.Context, req *desc.AddSegmentRequest) (*desc.AddSegmentResponse, error) {
	id, err := i.segmentService.AddSegment(ctx, req.GetSlug())
	if err != nil {
		return nil, err
	}

	return &desc.AddSegmentResponse{
		Id: id,
	}, nil
}
