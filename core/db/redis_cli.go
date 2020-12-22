package db

import (
	"github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func RedisConnect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		var err error
		
		instanceRedisCli.conn, err = redis.Dial("tcp", ":6379")
		HandleError(err)
		
		defer instanceRedisCli.conn.Close()
		_, err = instanceRedisCli.conn.Do("AUTH", "password")

		HandleError(err)
	}

	return instanceRedisCli
}

func (redisCli *RedisCli) RedisSetValue(key string, value []byte, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) RedisGetValue(key string) ([]byte, error) {
	return redis.Bytes(redisCli.conn.Do("GET", key))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}