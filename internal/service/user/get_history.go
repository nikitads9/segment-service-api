package user

import (
	"bytes"
	"context"
)

func (s *Service) GetUserHistoryCsv(ctx context.Context, id int64) (*bytes.Buffer, error) {
	return s.userRepository.GetUserHistoryCsv(ctx, id)
}
