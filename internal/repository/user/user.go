package user

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/user_mocks/user_service_repository.go -package=user_mocks . Repository

import (
	"bytes"
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = status.Error(codes.NotFound, "there is no instance with this name or id")
var ErrFailed = status.Error(codes.InvalidArgument, "the operation failed")

type Repository interface {
	GetSegmentId(ctx context.Context, slug string) (int64, error)
	AddToSegment(ctx context.Context, slug string, userId int64) error
	RemoveFromSegment(ctx context.Context, slug string, userId int64) error
	GetSegments(ctx context.Context, userId int64) ([]string, error)
	SetExpireTime(ctx context.Context, mod *model.SetExpireTimeInfo) error
	AddUser(ctx context.Context, userName string) (int64, error)
	GetUser(ctx context.Context, userId int64) (string, error)
	RemoveUser(ctx context.Context, userId int64) error
	GetUserHistoryCsv(ctx context.Context, userId int64) (*bytes.Buffer, error)
}

type repository struct {
	client db.Client
}

func NewUserRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}
