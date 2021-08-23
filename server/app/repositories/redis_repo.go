package repositories

import (
	"github.com/go-redis/redis"
)

type IRedisRepository interface {
	Search(word, key string, length int64) (result []string, err error)
	Insert(word, key string) (err error)
	Delete(key string) (err error)
}

type RedisRepository struct {
	rdb *redis.Client
}

func NewInstanceOfRedisRepository(
	address,
	password string,
	DB int) RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       DB,
	})
	return RedisRepository{
		rdb: rdb,
	}
}

func (a *RedisRepository) Insert(word, key string) error {
	_ = a.rdb.ZAdd(key, redis.Z{
		Score:  0,
		Member: word,
	})
	return nil
}

func (a *RedisRepository) Search(word, key string, searchLength int64) ([]string, error) {
	pipe := a.rdb.Pipeline()
	rnk := pipe.ZRank(key, word)
	result := pipe.ZRange(key, rnk.Val(), rnk.Val()+searchLength)
	_, err := pipe.Exec()
	if err != nil {
		return []string{}, err
	}
	searchResult, err := result.Result()
	return searchResult, err
}

func (a *RedisRepository) Delete(key string) error {
	a.rdb.Del(key)
	return nil
}
