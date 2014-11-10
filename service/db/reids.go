package db

import (
	"commonPackage"
)

func RedisKeyValue(key, val string) (reply interface{}, err error) {
	rc := GetRedisConn()
	return rc.Do("SET", key, val)
}

func RedisObjSet(key string, obj interface{}) (reply interface{}, err error) {
	rc := GetRedisConn()
	b := commonPackage.JSONMarshal(obj)
	return rc.Do("SET", key, b)
}

func RedisBytesSet(key string, b []byte) (reply interface{}, err error) {
	rc := GetRedisConn()
	return rc.Do("SET", key, b)
}

func RedisObjGet(key string, v interface{}) (reply interface{}, err error) {
	rc := GetRedisConn()
	r, err := rc.Do("GET", key)
	if err != nil {
		return nil, err
	}

	commonPackage.Printf(r)

	if str, ok := r.([]byte); ok {
		commonPackage.JSONUnmarshal(str, v)
		return r, nil
	} else {
		return nil, commonPackage.NewErr("not string")
	}

}
