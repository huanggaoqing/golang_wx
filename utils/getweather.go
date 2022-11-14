package utils

import (
	"encoding/json"
	"fmt"
	"go_http/tools"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

const WeatherKey string = "251518e073ef6c3c9504dd286c3f6a86"

// "temperature": "5",
// "humidity": "29",
// "info": "晴",
// "wid": "00",
// "direct": "东北风",
// "power": "6级",
// "aqi": "18"
type weatherData struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	Info        string `json:"info"`
	Wid         string `json:"wid"`
	Direct      string `json:"direct"`
	Power       string `json:"power"`
	Api         string `json:"api"`
}

// 从网上获取当天的天气
func GetWeather(city string) (*weatherData, error) {
	url := fmt.Sprintf("http://apis.juhe.cn/simpleWeather/query?city=%s&key=%s", city, WeatherKey)
	data := &weatherData{}
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	result, err := tools.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(result))
	res := gjson.ParseBytes(result).Get("result").Get("realtime")

	if err := json.Unmarshal([]byte(res.String()), data); err != nil {
		return nil, err
	}
	return data, nil
}
