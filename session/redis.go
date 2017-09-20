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
func GetRedisCache(cookie, key string) interface{} {
	val, err := client.Get(namespace + "session:" + cookie).Result()
	if err != nil {
		fmt.Println(err)
	}
	var s map[string]interface{}
	json.Unmarshal([]byte(val), &s)
	fmt.Println(comutil.TransInterfaceToString(s))
	fmt.Println("直接获取值：", s[key])
	return s[key]
}

//设置缓存
func SetRedisCache(key string, value map[string]interface{}) error {
	str := comutil.TransInterfaceToString(value)
	err := client.Set(namespace+"session:"+key, str, 10000*time.Second).Err()
	return err
}
