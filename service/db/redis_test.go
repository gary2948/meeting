package db

import (
//"github.com/garyburd/redigo/redis"
//"commonPackage"
//"testing"
)

type Client struct {
	Id   int64
	Addr string
}

//func TestRedis(t *testing.T) {
//	RedisKeyValue("key11", "33333")
//	c := Client{}
//	c.Addr = "123"
//	c.Id = 345
//	_, err := RedisObjSet("key22", c)
//	if err != nil {
//		t.Error(err)
//	}

//	d := Client{}
//	_, err = RedisObjGet("key22", &d)
//	if err != nil {
//		t.Error(err)
//	}
//	commonPackage.Printf(d)
//	commonPackage.Println(d.Addr)
//}
