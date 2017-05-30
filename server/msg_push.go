package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func parseRequestBody(r *http.Request) (*TextRequestBody, error) {
	res := &TextRequestBody{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("err:", err)
		return res, err
	}
	xml.Unmarshal(body, res)
	return res, nil
}

func value2CDATA(str string) CDATAText {
	return CDATAText{"<![CDATA[" + str + "]]>"}
}

func makeReponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	body := &TextResponseBody{}
	body.Content = value2CDATA(content)
	body.FromUserName = value2CDATA(fromUserName)
	body.ToUserName = value2CDATA(toUserName)
	body.MsgType = value2CDATA("text")
	body.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(body, " ", "  ")
}

func MessagePush(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("该请求非微信来源")
		return
	}
	if r.Method == "POST" {
		body, err := parseRequestBody(r)
		if err != nil {
			log.Println("解析消息出错:", err)
			return
		}
		//fmt.Printf("msg:%s \nuser:%s \ntouser:%s", body.Content, body.FromUserName, body.ToUserName)
		revmsg := GetTuLingMsg(body.Content)
		mergemsg := body.Content + "\t" + revmsg
		save2db(body.FromUserName, mergemsg)
		reponsebody, err := makeReponseBody(body.ToUserName, body.FromUserName, revmsg)
		if err != nil {
			log.Println("生成回复失败:", err)
			return
		}
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprintf(w, string(reponsebody))
	}
}
