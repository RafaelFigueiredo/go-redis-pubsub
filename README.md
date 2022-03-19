# Redis Pup/Sub


Just a sample project to test pub/sub on Redis. It connect to a redis database, create a channel named "my-topic" and keep sending the current timestamp and printing the latency to terminal.

## How to run
1. Star a redis database using docker, or remotely, I suggest use Digital Ocean's managed redis.

2. Export the environment variable with the connection string
```
export REDIS_URL="redis-url:port"
```

3. Run the code
```sh
go run main.go
```