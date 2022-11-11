package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// `evendayseData` 是一个有 10 个字段的结构。
// @property {string} Sid - 报价单的ID
// @property {string} Tts - 内容的文本到语音。
// @property {string} Content - 报价单的内容
// @property {string} Note - 以下属性在响应中并非全部可用。
// @property {string} Love - 点赞数
// @property {string} Translation - 内容的翻译
// @property {string} Picture - 图片的网址
// @property {string} Picture2 - 当天的图片
// @property {string} Caption - 图片的标题。
// @property {string} Dateline - 报价日期。
// @property {string} Spv - 页面被浏览的次数
// @property {string} Sppv - 页面被分享的次数
// @property {string} FenxiangImg - 用于分享的图片
// @property {string} Picture3 - 当天的图片
// @property {string} Picture4 - 图片的网址
// @property {[]string} Tags - 报价单的标签。
type evendayseData struct {
	Sid         string   `json:"sid"`
	Tts         string   `json:"tts"`
	Content     string   `json:"content"`
	Note        string   `json:"note"`
	Love        string   `json:"love"`
	Translation string   `json:"translation"`
	Picture     string   `json:"picture"`
	Picture2    string   `json:"picture2"`
	Caption     string   `json:"caption"`
	Dateline    string   `json:"dateline"`
	Spv         string   `json:"s_pv"`
	Sppv        string   `json:"sp_pv"`
	FenxiangImg string   `json:"fenxiang_img"`
	Picture3    string   `json:"picture3"`
	Picture4    string   `json:"picture4"`
	Tags        []string `json:"tags"`
}

// 从网上获取当天的英语句子
func GetEvendaysen() *evendayseData {
	data := &evendayseData{}
	rsp, err := http.Get("http://open.iciba.com/dsapi/")
	if err != nil {
		log.Panicln("获取每日一句失败")
	}
	defer rsp.Body.Close()
	result, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Panicln("读取Body失败")
	}
	json.Unmarshal(result, data)
	return data
}
