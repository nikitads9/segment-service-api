package segment

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no segment with this id")

type Repository interface {
	AddSegment(ctx context.Context, slug string) (int64, error)
	RemoveSegment(ctx context.Context, id int64) error
}
type repository struct {
	client db.Client
}

func NewSegmentRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddSegment(ctx context.Context, slug string) (int64, error) {
	return 0, nil
}

func (r *repository) RemoveSegment(ctx context.Context, id int64) error {
	return nil
}
