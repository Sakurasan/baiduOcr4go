package controllers

import (
	"baiduOcr4go/actions"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

func (this *MainController) Get() {
	this.TplName = "index.html"
	imgurl := this.Input().Get("url")

	var Token = actions.GetToken()
	if Token.Error == "" {
		fmt.Println("Has Token")
		fmt.Println(Token)
	}

	path := fmt.Sprintf("https://aip.baidubce.com/rest/2.0/ocr/v1/general?access_token=%s", Token)

	OcrResult := actions.HttpDo(path, imgurl)
	this.Data["App"] = OcrResult

}

func (this *MainController) Post() {
	this.TplName = "index.html"
	imgurl := this.Input().Get("url")
	imgbase64 := this.Input().Get("img_upload_base")
	imgocr := actions.Ocr{}
	imgocr.Image = imgbase64[strings.Index(imgbase64, ",")+1:]
	imgocr.Url = imgurl

	fmt.Println("Base64->", imgbase64, "\n")

	// fmt.Println("BASE64->", imgbase64[22:], "\n")
	fmt.Println("BASE64->", imgbase64[strings.Index(imgbase64, ",")+1:], "\n")

	var Token = actions.GetToken()
	if Token.Error == "" {
		fmt.Println("Has Token")
		fmt.Println(Token)
	}

	path := fmt.Sprintf("https://aip.baidubce.com/rest/2.0/ocr/v1/general?access_token=%s", Token.AccessToken)
	fmt.Println("path->", path)
	var OcrResult string
	OcrResult = actions.PostMan(path, &imgocr)
	wordResult := &actions.OcrResult{}
	err := json.Unmarshal([]byte(OcrResult), wordResult)
	if err != nil {
		fmt.Println("JSON 转结构失败")
		return
	}
	fmt.Println("解包结果", wordResult, "\n")

	if imgbase64 != "" {
		this.Data["RemoteImg"] = template.HTML("<img src=" + imgbase64 + " /><br>")
		// this.Data["RemoteImg"] = template.HTML(imgbase64)
	} else {
		this.Data["RemoteImg"] = template.HTML("<img src=" + imgurl + " /><br>")
	}
	var sl []template.HTML
	for _, slc := range wordResult.WordsResult {
		fmt.Println("返回内容->", slc.Words)
		sl = append(sl, template.HTML(slc.Words+"</br>"))

	}
	this.Data["App"] = sl

	// fmt.Println("返回内容->", OcrResult)
	// this.Data["App"] = OcrResult

}
