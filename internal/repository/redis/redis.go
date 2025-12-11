package dataonredis

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

/* 	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/config" */
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() {
	if RedisClient != nil {
		return
	}
	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
	/* cfg, _ := config.LoadConfig("config/config.yaml") */
	DB := os.Getenv("REDIS_DB")
	Timeout := os.Getenv("REDIS_TIMEOUT")
	intDB , _ := strconv.Atoi(DB)
	intTimeout, _ := strconv.Atoi(Timeout)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:        os.Getenv("REDIS_HOST"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          intDB,
		DialTimeout: time.Duration(intTimeout) * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Panicf("Failed to connect to redis client, error: %+v", err)
	}
}

func SaveUserToRedis(user model.User) error {
	key := strconv.Itoa(user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return RedisClient.Set(ctx, key, data, 0).Err()
}

func LoadUserFromRedis(userID int) (*model.User, error) {
	key := strconv.Itoa(userID)

	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsersFromRedis(start int, end int) ([]model.User, error) {
	var (
		users     []model.User
		cursor    uint64
		collected int
	)

	for {
		keys, nextCursor, err := RedisClient.Scan(context.Background(), cursor, "*", 2000).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range keys {

			if collected > end {
				break
			}

			log.Println("collected: ", collected)

			if collected >= start {

				val, err := RedisClient.Get(context.Background(), key).Result()
				if err != nil {
					return nil, err
				}

				var user model.User
				if err := json.Unmarshal([]byte(val), &user); err != nil {
					return nil, err
				}

				users = append(users, user)
			}

			collected++
		}

		log.Printf("Scanned %d keys, next cursor: %d\n", len(keys), nextCursor)
		cursor = nextCursor
		if cursor == 0 || collected <= end {
			break
		}
	}

	return users, nil
}

func DeleteUserFromRedis(userID int) error {
	key := strconv.Itoa(userID)
	return RedisClient.Del(context.Background(), key).Err()
}
