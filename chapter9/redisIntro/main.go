package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0,})
	pong, _ := client.Ping(context.TODO()).Result()
	fmt.Println(pong)
}
