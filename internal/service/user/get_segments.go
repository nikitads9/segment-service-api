package segment

import "context"

func (s *Service) GetSegments(ctx context.Context, id int64) ([]string, error) {
	return s.segmentRepository.GetSegments(ctx, id)
}
