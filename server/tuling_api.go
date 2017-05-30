package server

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	tulingurl string
	key       string
)

func InitTuLingApi(conf *viper.Viper) {
	tulingurl = conf.GetString("tulingurl")
	key = conf.GetString("tulingkey")
}

func GetTuLingMsg(content string) string {
	res := ""
	client := &http.Client{}
	info := url.Values{}
	info.Add("key", key)
	info.Add("info", content)
	mdata := info.Encode()
	request, err := http.NewRequest("POST", tulingurl, strings.NewReader(mdata))
	if err != nil {
		log.Println("err:", err)
		return res
	}
	request.Header.Set("content-Type", "application/x-www-form-urlencoded")
	reponse, _ := client.Do(request)
	result, err := ioutil.ReadAll(reponse.Body)

	request.Body.Close()
	if err != nil {
		log.Println("获取图灵消息失败:", err)
		return res
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		log.Println("json解析失败:", err)
		return res
	}
	return data["text"].(string)
}
