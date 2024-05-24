package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Создание клиента Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // адрес Redis сервера
		Password: "",               // пароль, если есть
		DB:       0,                // номер базы данных
	})

	ctx := context.Background()

	// Установка значения
	err := rdb.Set(ctx, "Костя", "норм челик, умеет играть в доту, однако много ноет (|)", 0).Err()
	if err != nil {
		panic(err)
	}

	// Получение значения
	val, err := rdb.Get(ctx, "Костя").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("значение Костяна:", val)
}
