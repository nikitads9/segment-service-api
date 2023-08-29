package user

import (
	"context"

	"github.com/nikitads9/segment-service-api/internal/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no segment with this id")

type Repository interface {
	ModifySegments(ctx context.Context, slugsToAdd []string, slugsToRemove []string, id int64) error
	GetSegments(ctx context.Context, id int64) ([]string, error)
}
type repository struct {
	client db.Client
}

func NewUserRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) ModifySegments(ctx context.Context, slugsToAdd []string, slugsToRemove []string, id int64) error {
	return nil
}

func (r *repository) GetSegments(ctx context.Context, id int64) ([]string, error) {
	return nil, nil
}
