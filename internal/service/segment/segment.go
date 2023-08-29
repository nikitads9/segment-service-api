package segment

import (
	"github.com/nikitads9/segment-service-api/internal/repository/segment"
)

type Service struct {
	segmentRepository segment.Repository
}

func NewSegmentService(segmentRepository segment.Repository) *Service {
	return &Service{
		segmentRepository: segmentRepository,
	}
}
