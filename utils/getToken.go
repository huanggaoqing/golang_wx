package utils

import (
	"encoding/json"
	"fmt"
	"go_http/tools"
	"log"
	"net/http"
)

const appId string = "wx528b18779bfa67e8"
const appSecret string = "c55fe5f550e2ebde4346a9de36ae7704"

type tokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetWxToken() (*tokenData, error) {
	data := &tokenData{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := tools.ReadAll(rsp.Body)
	log.Println(string(result))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, data); err != nil {
		return nil, err
	}
	return data, nil
}
