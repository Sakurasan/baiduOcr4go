package actions

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"shit/rspinfo"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type AToken struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

var (
	url_token = "https://aip.baidubce.com/oauth/2.0/token"
)

func GetToken() *AToken {
	path := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", beego.AppConfig.String("APIKey"), beego.AppConfig.String("SecretKey"))
	fmt.Println("path", path)
	result := Getman(path)
	fmt.Println("返回的Token Struct->", result, "\n------------------\n")

	AT, err := fmtJson(result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return AT
}

func Getman(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("Do Get Err")
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioread Err")
		return ""
	}
	return string(body)
}

func HttpsERV(path string) {
	resp, err := http.Post(path,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// func HttpDo(path string) string {
// 	client := &http.Client{}

// 	req, err := http.NewRequest("POST", path, strings.NewReader("name=cjb"))
// 	if err != nil {
// 		return ""
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	resp, err := client.Do(req)

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		// handle error
// 	}
// 	fmt.Println(string(body))
// 	return string(body)
// }

func PostOrderJ(path string, msg []byte) {

	/*通信发送包*/
	tr := &http.Transport{
		Dial: (&net.Dialer{
			//LocalAddr: localAddr,
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	Client := &http.Client{Transport: tr, Timeout: time.Second * 60}

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println("发送POST错误", err)
	}
	req.Header.Set("Content-Type", "application/json") //设置传输类型

	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println("POST: Client.Do error:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http 应答包出错resp.Status=", resp.Status)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Http 请求成功：", resp.Status)
	}

	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取应答出错：", err)
	}
	fmt.Println("打印返回信息:", string(Body))

	Rspstruct, err := rspinfo.GetRspDataStruct(Body)
	if err != nil {
		log.Fatal("获取返回结构失败")
	}
	fmt.Println("===============================")
	fmt.Println(Rspstruct)

}

//拆解返回的Token 报文为结构体
func fmtJson(str string) (*AToken, error) {
	Token := &AToken{}
	err := json.Unmarshal([]byte(str), Token)
	if err != nil {
		fmt.Println("jSON UM ERR")
		return nil, errors.New("jSON UM ERR")
	}
	return Token, nil
}
