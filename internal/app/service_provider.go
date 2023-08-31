package app

import (
	"context"
	"log"

	segmentV1 "github.com/nikitads9/segment-service-api/internal/api/segment_v1"
	userV1 "github.com/nikitads9/segment-service-api/internal/api/user_v1"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/client/db/transaction"
	"github.com/nikitads9/segment-service-api/internal/config"
	segmentRepository "github.com/nikitads9/segment-service-api/internal/repository/segment"
	userRepository "github.com/nikitads9/segment-service-api/internal/repository/user"
	"github.com/nikitads9/segment-service-api/internal/service/segment"
	"github.com/nikitads9/segment-service-api/internal/service/user"
)

type serviceProvider struct {
	db         db.Client
	txManager  db.TxManager
	configPath string
	config     *config.Config

	segmentRepository segmentRepository.Repository
	userRepository    userRepository.Repository
	segmentService    *segment.Service
	userService       *user.Service

	segmentImpl *segmentV1.Implementation
	userImpl    *userV1.Implementation
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("could not get config err: %s", err.Error())
		}
		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can`t connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.Read(s.configPath)
		if err != nil {
			log.Fatalf("could not get config err: %s", err)
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetSegmentRepository(ctx context.Context) segmentRepository.Repository {
	if s.segmentRepository == nil {
		s.segmentRepository = segmentRepository.NewSegmentRepository(s.GetDB(ctx))
		return s.segmentRepository
	}

	return s.segmentRepository
}

func (s *serviceProvider) GetUserRepository(ctx context.Context) userRepository.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.GetDB(ctx))
		return s.userRepository
	}

	return s.userRepository
}

func (s *serviceProvider) GetSegmentService(ctx context.Context) *segment.Service {
	if s.segmentService == nil {
		segmentRepository := s.GetSegmentRepository(ctx)
		s.segmentService = segment.NewSegmentService(segmentRepository)
	}

	return s.segmentService
}

func (s *serviceProvider) GetUserService(ctx context.Context) *user.Service {
	if s.userService == nil {
		userRepository := s.GetUserRepository(ctx)
		s.userService = user.NewUserService(userRepository, s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) GetSegmentImpl(ctx context.Context) *segmentV1.Implementation {
	if s.segmentImpl == nil {
		s.segmentImpl = segmentV1.NewImplementation(s.GetSegmentService(ctx))
	}

	return s.segmentImpl
}

func (s *serviceProvider) GetUserImpl(ctx context.Context) *userV1.Implementation {
	if s.userImpl == nil {
		s.userImpl = userV1.NewImplementation(s.GetUserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.GetDB(ctx).DB())
	}

	return s.txManager
}
