package actions

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	url_ocr = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
)

type Ocr struct {
	Image            string `json:"image,omitempty"`
	Url              string `json:"url,omitempty"`
	Langusge_type    string `json:"language_type,omitempty"`
	Detect_direction string `json:"detect_direction,omitempty"`
	Detect_langusge  string `json:"detect_language,omitempty"`
	Pronabilly       string `json:"probability,omitempty"`
	Type             string `json:"type,omitempty"`
}

type OcrResult struct {
	Direction      int32  `json:"direction,omitempty"`
	Probability    string `json:"probability,omitempty"`
	LogID          int64  `json:"log_id"`
	WordsResultNum int    `json:"words_result_num"`
	WordsResult    []struct {
		Location struct {
			Width  int `json:"width"`
			Top    int `json:"top"`
			Height int `json:"height"`
			Left   int `json:"left"`
		} `json:"location"`
		Words string `json:"words"`
	} `json:"words_result"`
}

type AutoGenerated struct {
	Location struct {
		Width  int `json:"width"`
		Top    int `json:"top"`
		Height int `json:"height"`
		Left   int `json:"left"`
	} `json:"location"`
	Words string `json:"words"`
}

// type Ocr_result struct {
// 	Direction      int32    `json:"direction,omitempty"`
// 	Probability    string   `json:"probability,omitempty"`
// 	LogID          int64    `json:"log_id,omitempty"`
// 	WordsResultNum int      `json:"words_result_num,omitempty"`
// 	WordsResult    []string `json:"words_result,omitempty"`
// 	// WordsResult    []struct {
// 	// 	Words string `json:"words,omitempty"`
// 	// } `json:"words_result,omitempty"`
// }

func DoOcr(path string, msg []byte) {
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
	// req.Header.Set("Content-Type", "application/json") //设置传输类型
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

	fmt.Println("===============================")
}

//图片转 BASE64
func DealImg() []byte {
	ff, _ := ioutil.ReadFile("output2.jpg") //我还是喜欢用这个快速读文件
	bufstore := make([]byte, 5000000)       //数据缓存
	base64.StdEncoding.Encode(bufstore, ff) // 文件转base64
	// _ = ioutil.WriteFile("./output2.jpg.txt", dist, 0666) //直接写入到文件就ok完活了。
	return bufstore
}

// //对OCR返回解析为struct
// func (this *Ocr_result) UnMarshaL(result string) []string {
// 	err := json.Unmarshal([]byte(result), this)
// 	if err != nil {
// 		fmt.Println("Err :Ocr_result,UnMarshaL")
// 	}

// 	return this.WordsResult
// }

func HttpDo(path, msg string) string {
	client := &http.Client{}

	req, err := http.NewRequest("POST", path, strings.NewReader(msg))
	if err != nil {
		return ""
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return string(body)
}

func PostMan(path string, ocrS *Ocr) string { //发送请求内容

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
	// Client := &http.Client{Transport: nil, Timeout: time.Second * 60}

	postValue := url.Values{}
	if ocrS.Image != "" {
		postValue.Add("image", ocrS.Image)
	} else {
		postValue.Add("url", ocrS.Url)
	}

	resp, err := Client.PostForm(path, postValue) //-----URL -----
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http 应答包出错resp.Status=", resp.Status)
	}
	var body []byte
	if resp.StatusCode == 200 {
		body, _ = ioutil.ReadAll(resp.Body)
		fmt.Println("Http 请求成功：", resp.Status)
	}

	return string(body)
}