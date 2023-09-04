package segment

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/segment_mocks/segment_service_repository.go -package=segment_mocks . Repository

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrSegNotFound = status.Error(codes.NotFound, "there is no segment with this name or id")
var ErrFailed = status.Error(codes.InvalidArgument, "the operation failed")

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
