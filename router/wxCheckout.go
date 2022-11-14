package router

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

const token string = "huanggaoqing"

// @property {string} Signature - 签名串是三个参数按照时间戳、nonce、token的顺序加密的结果。
// @property {string} Echostr - 一个随机字符串，用于验证请求是否来自微信。
// @property {string} Timestamp - 请求的时间戳。
// @property {string} Nonce - 一个随机字符串，用于防止重放攻击。
type wxCheckoutParams struct {
	Signature string `form:"signature"`
	Echostr   string `form:"echostr"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
}

// 用于验证微信服务器的签名。
func checkout(ctx *gin.Context) {
	wxParams := &wxCheckoutParams{}
	if err := ctx.BindQuery(wxParams); err != nil {
		log.Println("绑定失败", err.Error())
	}
	// 将 echostr， timestamp，nonce进行排序
	tempArray := []string{token, wxParams.Nonce, wxParams.Timestamp}
	sort.Strings(tempArray)
	// 将排序完的字符串切片拼接成一个字符串
	tempString := strings.Join(tempArray, "")
	// 将拼接完的字符串进行sha1加密
	s := sha1.New()
	s.Write([]byte(tempString))
	sha1String := hex.EncodeToString(s.Sum(nil))
	// 加密完的字符串和Signature进行对比
	if sha1String == wxParams.Signature {
		log.Println("验证成功")
		ctx.String(http.StatusOK, wxParams.Echostr)
		return
	}
	log.Println("验证失败")
}
