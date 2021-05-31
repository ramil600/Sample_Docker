package main

import (
	"fmt"

	redis "gopkg.in/redis.v4"
)

func main() {
	fmt.Println("Hello, world!")
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
