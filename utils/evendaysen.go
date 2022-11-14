package utils

import (
	"encoding/json"
	"go_http/tools"
	"net/http"
)

// @property {string} Sid - 句子的ID
// @property {string} Tts - 内容的文本到语音。
// @property {string} Content - 句子的内容（英文）
// @property {string} Note - 句子的内容（中文）
// @property {string} Love - 点赞数
// @property {string} Translation - 内容的翻译
// @property {string} Picture - 图片的网址
// @property {string} Picture2 - 当天的图片
// @property {string} Caption - 图片的标题。
// @property {string} Dateline - 句子日期。
// @property {string} Spv - 页面被浏览的次数
// @property {string} Sppv - 页面被分享的次数
// @property {string} FenxiangImg - 用于分享的图片
// @property {string} Picture3 - 当天的图片
// @property {string} Picture4 - 图片的网址
// @property {[]string} Tags - 句子的标签。
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
func GetEvendaysen() (*evendayseData, error) {
	data := &evendayseData{}
	rsp, err := http.Get("http://open.iciba.com/dsapi/")
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	result, err := tools.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(result, data); err != nil {
		return nil, err
	}
	return data, nil
}
