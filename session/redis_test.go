package session

import (
	"fmt"
	"golang-study/utils/comutil"
	"testing"
)

func TestGetRedisCache(t *testing.T) {
	cookie := "15806111230"
	data := GetRedisCache(cookie, "name")
	fmt.Println(comutil.TransInterfaceToString(data))
}

//
//func TestSetRedisCache(t *testing.T) {
//	cookie := "15806111230"
//	value := make(map[string]interface{})
//	value["name"] = "admin123"
//	value["age"] = 19
//	value["sex"] = 0
//	value["address"] = "shanghai"
//	temp := make(map[string]interface{})
//	temp["love"] = "ss"
//	temp["city"] = "nj"
//	value["data"] = temp
//
//	fmt.Println("入参：", comutil.TransInterfaceToString(value))
//	err := SetRedisCache(cookie, "info", value)
//	fmt.Println(err)
//}

func TestGetSession(t *testing.T) {
	SessionId = "17721021494"
	GetSession("abc")
}

func TestSetSession(t *testing.T) {
	SessionId = "17721021494"
	SetSession("name", "root")
	SetSession("openid", "12346666666666666666666")
}
