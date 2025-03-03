package main

import (
    "context"
    "github.com/redis/go-redis/v9"
    "fmt"
)

func main() {
    // Initialize a new client to connect to Redis
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",  // Redis address
        Password: "",               // no password by default
        DB:       0,                // use default DB 0
    })
    ctx := context.Background()    // context for Redis calls
    // Ping the Redis server to test connection
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }
    println(pong)  // should print "PONG"
    
    
    //set vaue to redis
    rdb.Set(ctx, "username:1001", "alice", 0).Err()
    if err != nil {
        fmt.Println("Error setting value:", err)
    }
    
    
    //get value from redis
    val, err := rdb.Get(ctx, "username:1001").Result()
    if err == redis.Nil {
        fmt.Println("Key not found")
    } else if err != nil {
        fmt.Println("Error getting value:", err)
    } else {
        fmt.Println("Got value:", val)  // expect "alice"
    }
    
    //delete value from redus
    res, err := rdb.Del(ctx, "username:1001").Result()
    if err != nil {
        fmt.Println("Error deleting key:", err)
    } else {
        fmt.Println("Deleted keys:", res)  // res is 1 if the key was present
    }
}