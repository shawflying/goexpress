package session

import (
	"fmt"
	"golang-study/utils/comutil"
	"testing"
)

//func TestGetRedisCache(t *testing.T) {
//	cookie := "15806111230"
//	data := GetRedisCache(cookie, "name")
//	fmt.Println(comutil.TransInterfaceToString(data))
//}

func TestSetRedisCache(t *testing.T) {
	cookie := "15806111230"
	value := make(map[string]interface{})
	value["name"] = "admin123"
	value["age"] = 19
	value["sex"] = 0
	value["address"] = "shanghai"
	temp := make(map[string]interface{})
	temp["love"] = "ss"
	temp["city"] = "nj"
	value["data"] = temp

	fmt.Println(comutil.TransInterfaceToString(value))
	err := SetRedisCache(cookie, value)
	fmt.Println(err)
}
