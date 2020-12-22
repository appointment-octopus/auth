package db

import (
	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return conn
}

func RedisSetValue(key string, value []byte, expiration ...interface{}) error {
	conn := RedisConnect()
	defer conn.Close()

	_, err := conn.Do("SET", key, []byte(value))
	HandleError(err)

	conn.Do("EXPIRE", key, expiration[0])

	return err
}

func RedisGetValue(key string) ([]byte, error) {
	conn := RedisConnect()
	defer conn.Close()

	var err error
	value, err := conn.Do("GET", key)
	HandleError(err)

	var data []byte
	data, err = redis.Bytes(value, err)

	return data, err
}

func Flush(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("DEL", key))
	HandleError(err)

	return data, err
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}