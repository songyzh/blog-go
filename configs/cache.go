package configs

import "github.com/go-redis/redis/v7"

var Cache *redis.Client

func init(){
	Cache = redis.NewClient(&redis.Options{
		Addr:     "122.51.74.53:6379",
		Password: "uwXBT/U7wdThBYhjproRC%Kf;BQ{Z", // no password set
		DB:       0,  // use default DB
	})
}
