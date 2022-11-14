package utils

import (
	"encoding/json"
	"fmt"
	"go_http/tools"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

type userData struct {
	Openid []string `json:"openid"`
}

func GetUserList(token string) ([]string, error) {
	data := &userData{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=", token)
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := tools.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(result))
	res := gjson.ParseBytes(result).Get("data")
	if err := json.Unmarshal([]byte(res.String()), data); err != nil {
		return nil, err
	}
	return data.Openid, nil
}
