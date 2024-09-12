package app

import (
	"context"
	"log"

	"github.com/vadskev/banners-rotation/internal/api/rotation"
	"github.com/vadskev/banners-rotation/internal/config"
	"github.com/vadskev/banners-rotation/internal/config/env"
	"github.com/vadskev/banners-rotation/internal/queue"
	"github.com/vadskev/banners-rotation/internal/queue/kafka"
	"github.com/vadskev/banners-rotation/internal/storage"
	"github.com/vadskev/banners-rotation/internal/storage/pg"
)

type serviceProvider struct {
	logConfig  config.LogConfig
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	kafkaProducerConfig config.KafkaProducerConfig

	dbStorage   storage.Storage
	queueClient queue.Queue

	rotationImpl *rotation.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) LogConfig() env.LogConfig {
	if s.logConfig == nil {
		cfg, err := env.NewLogConfig()
		if err != nil {
			log.Fatalf("failed to get log config: %s", err.Error())
		}
		s.logConfig = cfg
	}
	return s.logConfig
}

func (s *serviceProvider) KafkaProducerConfig() config.KafkaProducerConfig {
	if s.kafkaProducerConfig == nil {
		cfg, err := env.NewKafkaProducerConfig()
		if err != nil {
			log.Fatalf("failed to get kafka producer config: %s", err.Error())
		}

		s.kafkaProducerConfig = cfg
	}

	return s.kafkaProducerConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) storage.Storage {
	if s.dbStorage == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbStorage = cl
	}
	return s.dbStorage
}

func (s *serviceProvider) QueueClient(ctx context.Context) queue.Queue {
	if s.queueClient == nil {
		qu, err := kafka.New(ctx, s.KafkaProducerConfig().Brokers(), s.KafkaProducerConfig().Config())
		if err != nil {
			log.Fatalf("failed to create kafka client: %v", err)
		}
		s.queueClient = qu
	}
	return s.queueClient
}

func (s *serviceProvider) RotationImpl(ctx context.Context) *rotation.Implementation {
	if s.rotationImpl == nil {
		s.rotationImpl = rotation.NewImplementation(s.DBClient(ctx), s.QueueClient(ctx))
	}
	return s.rotationImpl
}
