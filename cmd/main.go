package main

import (
	"context"
	"log"

	"github.com/vadskev/banners-rotation/internal/app"
)

func main() {

	//TODO
	//Добавить баннер
	//Удалить баннер
	//Засчитать переход
	//Выбрать баннер для показа
	//
	//
	//Выгрузка статистики kafka
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to run app: %s", err.Error())
	}

}
