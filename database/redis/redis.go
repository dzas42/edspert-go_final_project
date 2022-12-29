package redis

import (
	"final-project/internal/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func InitConnection(conf config.Config) *redis.Client {
   fmt.Println(conf.RedisHost)
   rdb := redis.NewClient(&redis.Options{
      Addr:     conf.RedisHost,
      Password: conf.RedisUserPassword, // no password set
      DB:       0,                      // use default DB
   })
   if err := rdb.Ping().Err(); err != nil {
      log.Fatalf("Couldn't connect redis")
   }
   
   return rdb
}