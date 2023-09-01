package segment_v1

import (
	"github.com/nikitads9/segment-service-api/internal/service/segment"
	desc "github.com/nikitads9/segment-service-api/pkg/segment_api"
)

type Implementation struct {
	desc.UnimplementedSegmentV1ServiceServer
	segmentService *segment.Service
}

func NewImplementation(segmentService *segment.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedSegmentV1ServiceServer{},
		segmentService,
	}
}

func newMockImplementation(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedSegmentV1ServiceServer{},
		i.segmentService,
	}
}
