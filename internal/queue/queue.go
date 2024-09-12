package queue

import "github.com/vadskev/banners-rotation/internal/models"

type Queue interface {
	SendMessage(msg models.Message)
	Close()
}
