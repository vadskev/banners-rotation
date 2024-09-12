package rotation

import (
	"github.com/vadskev/banners-rotation/internal/queue"
	"github.com/vadskev/banners-rotation/internal/storage"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

type Implementation struct {
	desc.UnimplementedRotationServer
	storageService storage.Storage
	kafkaProducer  queue.Queue
}

func NewImplementation(storageService storage.Storage, kafkaProducer queue.Queue) *Implementation {
	return &Implementation{
		storageService: storageService,
		kafkaProducer:  kafkaProducer,
	}
}
