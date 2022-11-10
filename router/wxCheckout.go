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

type wxCheckoutParams struct {
	Signature string `form:"signature"`
	Echostr   string `form:"echostr"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
}

func checkout(ctx *gin.Context) {
	wxParams := &wxCheckoutParams{}
	if err := ctx.BindQuery(wxParams); err != nil {
		log.Panicln("绑定失败")
	}
	// 将 echostr， timestamp，nonce进行排序
	tempArray := []string{token, wxParams.Nonce, wxParams.Timestamp}
	sort.Strings(tempArray)
	// 将排序完的字符串切片拼接成一个字符串
	tempString := strings.Join(tempArray, "")
	// 将拼接完的字符串进行sha1加密
	s := sha1.New()
	s.Write([]byte(tempString))
	hex.EncodeToString(s.Sum(nil))
	sha1String := hex.EncodeToString(s.Sum(nil))
	// 加密完的字符串和Signature进行对比
	if sha1String == wxParams.Signature {
		log.Println("验证成功")
		ctx.String(http.StatusOK, wxParams.Echostr)
		return
	}
	log.Panicln("验证失败")
}
