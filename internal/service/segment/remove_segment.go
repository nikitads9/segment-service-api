package segment

import (
	"context"
)

func (s *Service) RemoveSegment(ctx context.Context, id int64) error {
	return s.segmentRepository.RemoveSegment(ctx, id)
}
