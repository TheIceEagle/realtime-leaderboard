package main

import (
	"context"
	"fmt"
	_ "log"
	// "time"

	"github.com/redis/go-redis/v9"
)

type RedisConnection struct{
	client *redis.Client
}

func NewRedisConnections() (*RedisConnection) {
	client:= redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	
	
	return  &RedisConnection{
		client: client,
	}

}

func main() {
	ctx := context.Background()

	rdb := NewRedisConnections()
	
	
	err := rdb.client.Publish(ctx, "Ddakji", "payload").Err()
		if err != nil {
		panic(err)
		}
	
	
	pubsub := rdb.client.Subscribe(ctx, "Ddakji")
	defer pubsub.Close()
	
	
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
		}

	
}
