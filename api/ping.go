package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"fmt"
	"net/url"
	"io/ioutil"
	"net/http"
	"to-do-list/pkg/util"
)

var (
	appid		string
	secret		string
	grant_type		string
)

func InitWeApp() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		util.LogrusObj.Info("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	
	appid = file.Section("weapp").Key("appid").String()
	secret = file.Section("weapp").Key("secret").String()
	grant_type = file.Section("weapp").Key("grant_type").String()
}

func Ping(c *gin.Context) {
	InitWeApp()

	prefix := "https://api.weixin.qq.com/sns/jscode2session"
	uri, err := url.Parse(prefix)
	if err != nil {
		return
	}

	params := url.Values{}
	params.Add("appid", appid)
	params.Add("secret", secret)
	params.Add("grant_type", grant_type)
	params.Add("js_code", c.Query("js_code"))

	uri.RawQuery = params.Encode()

	fmt.Println(c.Query("js_code"))

	resp, err := http.Get(uri.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	c.JSON(200, string(body))
}
