package app

import (
	"context"
	"log"

	"github.com/nikitads9/segment-service-api/internal/config"
	"github.com/nikitads9/segment-service-api/internal/pkg/db"
	segmentRepository "github.com/nikitads9/segment-service-api/internal/repository/segment"
	"github.com/nikitads9/segment-service-api/internal/service/segment"
)

type serviceProvider struct {
	db                db.Client
	configPath        string
	config            *config.Config
	segmentRepository segmentRepository.Repository
	segmentService    *segment.Service
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

func (s *serviceProvider) GetSegmentService(ctx context.Context) *segment.Service {
	if s.segmentService == nil {
		segmentRepository := s.GetSegmentRepository(ctx)
		s.segmentService = segment.NewSegmentService(segmentRepository)
	}

	return s.segmentService
}
