package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redis_url = os.Getenv("REDIS_URL")

func main() {
	log.Println("connecting to redis database")

	r := redis.NewClient(&redis.Options{
		Addr:        redis_url,
		Password:    "gM5CtJvri8HPq12n",
		DB:          0, // use default DB,
		IdleTimeout: 5 * time.Second,
		TLSConfig:   &tls.Config{},
	})
	log.Println("connect")

	go func() {
		for {
			payload := time.Now()
			r.Publish(ctx, "my-topic", payload)
			time.Sleep(1 * time.Second)
		}
	}()

	ch := r.Subscribe(ctx, "my-topic").Channel()
	for message := range ch {
		sent_at, _ := time.Parse("2006-01-02T15:04:05-07:00", message.Payload)
		latency := time.Since(sent_at)
		log.Printf("channel: %s\tlatency: %s", message.Channel, latency)
	}
}
