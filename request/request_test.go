package request

import (
	"testing"
	"fmt"
	"github.com/shawflying/beego-common-utils/utils/comutil"
)

func TestGet(t *testing.T) {
	body, terr := Get("http://m.sh.189.cn/service/node/crypto?data=abc123&key=express&type=0")
	if terr != nil {

	}
	fmt.Println("get 请求结果" + comutil.TransInterfaceToString(body))
}

func TestPost(t *testing.T) {
	var PayParams map[string]interface{}
	PayParams = make(map[string]interface{});
	PayParams["money"] = "5"
	PayParams["number"] = "17721021494"
	PayParams["bpnum"] = ""
	PayParams["payid"] = ""
	PayParams["openid"] = "oKXUCj1MOddnp-sCpGi1J1dg3TyM"
	PayParams["from"] = "disney"
	PayParams["channel"] = "0"
	PayParams["note"] = "迪士尼活动"
	body, terr := Post("http://127.0.0.1:3255/test/post", PayParams);
	if terr != nil {

	}
	fmt.Println("Post 请求结果" + comutil.TransInterfaceToString(body))
}

func TestPostForm(t *testing.T) {
	PayParams := make(map[string]interface{});
	PayParams["mobile"] = "17721021494"
	PayParams["channel"] = "2"
	PayParams["desKey"] = "dzqd-wt-flow"

	//body, terr := PostForm("http://172.16.50.138:8091/csb/1.0/Encrypt", PayParams);
	body, terr := PostForm("http://httpbin.org/post", PayParams);
	if terr != nil {

	}
	fmt.Println("PostForm 请求结果" + comutil.TransInterfaceToString(body))
}
