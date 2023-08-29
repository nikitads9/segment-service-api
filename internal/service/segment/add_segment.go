package segment

import (
	"context"
)

func (s *Service) AddSegment(ctx context.Context, slug string) (int64, error) {
	return s.segmentRepository.AddSegment(ctx, slug)
}
