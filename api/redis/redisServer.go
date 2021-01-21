package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/samhj/AchmadGo/api/config"
	models "github.com/samhj/AchmadGo/api/models"
)

//Redis ...
func Redis() *redis.Client {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     config.Config("REDIS_HOST") + ":" + config.Config("REDIS_PORT"),
		Password: config.Config("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis Setup Error:", err.Error())
		return nil
	}

	fmt.Println(pong)

	return client

}

//SetServer ...
func SetServer(conn *redis.Client, item []byte) error {
	var ctx = context.Background()

	err := conn.Set(ctx, "server", item, 10*time.Minute).Err()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

//GetServer ...
func GetServer(conn *redis.Client) (models.Server, error) {
	var ctx = context.Background()

	val, err := conn.Get(ctx, "server").Result()
	if err != nil {
		fmt.Println(err)
	}
	server := models.Server{}
	err = json.Unmarshal([]byte(val), &server)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("VAL:",val)

	return server, err
}
