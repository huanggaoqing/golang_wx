package utils

import (
	"encoding/json"
	"log"
)

type msgStruct struct {
	Touser     string   `json:"touser"`
	TemplateID string   `json:"template_id"`
	URL        string   `json:"url"`
	Topcolor   string   `json:"topcolor"`
	Data       *msgData `json:"data"`
}

type msgData struct {
	NewDate     *msgItem `json:"newDate"`
	City        *msgItem `json:"city"`
	Info        *msgItem `json:"info"`
	Low         *msgItem `json:"low"`
	Temperature *msgItem `json:"temperature"`
	Direct      *msgItem `json:"direct"`
	Content     *msgItem `json:"content"`
	Note        *msgItem `json:"note"`
}

type msgItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateParams struct {
	Openid      string
	Date        string
	City        string
	Info        string
	Temperature string
	Content     string
	Note        string
	Direct      string
}

func GetTemplate(params *TemplateParams) (string, error) {
	msg := &msgData{
		NewDate: &msgItem{
			Value: params.Date,
			Color: "#8beae5",
		},
		City: &msgItem{
			Value: params.City,
			Color: "#9c989c",
		},
		Info: &msgItem{
			Value: params.Info,
			Color: "#d8d596",
		},
		Temperature: &msgItem{
			Value: params.Temperature,
			Color: "#aea082",
		},
		Direct: &msgItem{
			Value: params.Direct,
			Color: "#8ab882",
		},
		Content: &msgItem{
			Value: params.Content,
			Color: "#67696d",
		},
		Note: &msgItem{
			Value: params.Note,
			Color: "#67696d",
		},
	}
	data := &msgStruct{
		Touser:     params.Openid,
		TemplateID: "Y4AiCznOOW0GDVtkTmKVTyDLIIulufZEhAfjFPzz8EM",
		URL:        "http://weixin.qq.com/download",
		Topcolor:   "#FF0000",
		Data:       msg,
	}
	res, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	log.Println(string(res))
	return string(res), nil
}
