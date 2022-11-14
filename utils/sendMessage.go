package utils

import (
	"fmt"
	"go_http/tools"
	"log"
	"net/http"
	"strings"
	"time"
)

func StartSend() {
	evendaysen, err := GetEvendaysen()
	if err != nil {
		log.Println("获取当天句子失败", err.Error())
	}
	token, err := GetWxToken()
	if err != nil {
		log.Println("获取token失败", err.Error())
	}
	users, err := GetUserList(token.AccessToken)
	for _, openId := range users {
		if openId == "ooQma5qKDtUMnhnz4RiesakAzGdk" {
			go SendMessage(openId, evendaysen, token.AccessToken, "上海")
		} else {
			go SendMessage(openId, evendaysen, token.AccessToken, "北京")
		}
	}
	if err != nil {
		log.Println("获取用户列表失败", err.Error())
	}

}

func SendMessage(openId string, evendaysen *evendayseData, token string, city string) {
	weather, err := GetWeather("北京")
	if err != nil {
		log.Println("获取当天天气失败", err.Error())
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)
	sendMsgData := &TemplateParams{
		Openid:      openId,
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		City:        city,
		Info:        weather.Info,
		Temperature: weather.Temperature,
		Direct:      weather.Direct,
		Content:     evendaysen.Content,
		Note:        evendaysen.Note,
	}
	temp, err := GetTemplate(sendMsgData)
	if err != nil {
		log.Println("获取模版失败", err.Error())
	}
	rsp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(temp))
	if err != nil {
		log.Println("发送消息失败", err.Error())
	}
	result, err := tools.ReadAll(rsp.Body)
	if err != nil {
		log.Println("序列化失败", err.Error())
	}
	log.Println("rsp", string(result))
}
