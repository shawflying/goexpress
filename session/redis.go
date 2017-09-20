package session

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"golang-study/utils/comutil"
	"time"
)

var namespace = "telecom:"

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2,
		PoolSize: 10,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

var client = createClient()

//获取缓存信息
func GetRedisCache(SessionId, key string) interface{} {
	val, err := client.Get(namespace + "session:" + SessionId).Result()
	if err != nil {
		fmt.Println(err)
	}
	var s map[string]interface{}
	json.Unmarshal([]byte(val), &s)
	fmt.Println("获取session值：", comutil.TransInterfaceToString(s))
	fmt.Println("直接获取值：", s[key])
	return s[key]
}

//设置缓存 直接存储会重写掉
func SetRedisCache(cookie string, key string, value interface{}) error {
	var s map[string]interface{}
	val, err := client.Get(namespace + "session:" + cookie).Result()
	if err != nil {
		s = make(map[string]interface{})
		s[key] = value
	} else {
		json.Unmarshal([]byte(val), &s)
	}

	fmt.Println("存放session前：", comutil.TransInterfaceToString(s))
	s[key] = value
	fmt.Println("存放session后：", comutil.TransInterfaceToString(s))
	str := comutil.TransInterfaceToString(s)
	err = client.Set(namespace+"session:"+cookie, str, 10000*time.Second).Err()
	return err
}

var SessionId = ""

func GetSession(key string) interface{} {
	return GetRedisCache(SessionId, key)
}

func SetSession(key string, value interface{}) error {
	return SetRedisCache(SessionId, key, value)
}
