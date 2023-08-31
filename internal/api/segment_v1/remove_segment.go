package segment_v1

import (
	"context"

	desc "github.com/nikitads9/segment-service-api/pkg/segment_api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) RemoveSegment(ctx context.Context, req *desc.RemoveSegmentRequest) (*emptypb.Empty, error) {
	err := i.segmentService.RemoveSegment(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
