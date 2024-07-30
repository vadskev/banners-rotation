package rotation

import (
	"github.com/vadskev/banners-rotation/internal/storage"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

type Implementation struct {
	desc.UnimplementedRotationServer
	storageService storage.Storage
}

func NewImplementation(storageService storage.Storage) *Implementation {
	return &Implementation{
		storageService: storageService,
	}
}
