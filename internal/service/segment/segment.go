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

func NewMockSegmentService(deps ...interface{}) *Service {
	is := Service{}
	for _, val := range deps {
		switch s := val.(type) {
		case segment.Repository:
			is.segmentRepository = s
		}
	}
	return &is
}
