package network

import (
	"io/ioutil"
	"net/http"
	"log"
	"strings"
	//"fmt"
	"net/url"
	//"fmt"
	"fmt"
	"math/rand"
	"strconv"
)


func GetTextByUrl(url string) string {
	client := &http.Client{}

	var r http.Request
	r.ParseForm()
	r.Form.Add("phone", "18051156285")
	r.Form.Add("password", "076742567778be885ad66804ec9facb21a4296aa41ac29c4f7d5afe7a206c699")

	bodystr := strings.TrimSpace(r.Form.Encode())

	req, err := http.NewRequest("POST", url, strings.NewReader(bodystr))
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("User-Agent", userAgent)
	//req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var textByte []byte
	textByte, _ = ioutil.ReadAll(resp.Body)
	//log.Println("result:", string(textByte[:]))
	return string(textByte)
}

func HttpPostForm(urlStr string,data url.Values) ([]byte ,error) {
	resp, err := http.PostForm(urlStr, data)

	if err != nil {
		// handle error
		log.Print("request error:" , err.Error())
		return nil,err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return body,nil
}

func HttpPost(urlstr string, data url.Values, userToken string) string {
	client := &http.Client{}

	var r http.Request
	r.ParseForm()
	for key,value := range data {
		fmt.Println(key,":",value)
		log.Println("key:",key,";value:",value[0])
		r.Form.Add(key, value[0])
	}

	bodystr := strings.TrimSpace(r.Form.Encode())

	req, err := http.NewRequest("POST", urlstr, strings.NewReader(bodystr))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("device", "android")
	req.Header.Add("version", "1.2.0")
	req.Header.Add("wToken", "")
	req.Header.Add("Authorization", "Token " + userToken)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var textByte []byte
	textByte, _ = ioutil.ReadAll(resp.Body)
	//log.Println("result:", string(textByte[:]))
	return string(textByte)
}

func GetNonce() string {
	var nonce string
	for i := 0;i < 32;i++ {
		nonce = nonce + getX16(rand.Intn(16))
		if i == 7 || i == 11 || i == 15 || i == 19 {
			nonce = nonce + "-"
		}
	}
	return nonce

}

func getX16(num int) string {
	if num < 10 {
		return strconv.Itoa(num)
	}
	if num == 10 {
		return "a"
	}
	if num == 11 {
		return "b"
	}
	if num == 12 {
		return "c"
	}
	if num == 13 {
		return "d"
	}
	if num == 14 {
		return "e"
	}
	if num == 15 {
		return "f"
	}
	return "0"
}

//func GetTextByJson(jsonStr string) string {
//	msg := parser.Parser(jsonStr)
//	msg.HTML_TEXT = GetTextByUrl(msg.URL, msg.COOKIE, msg.USER_AGENT)
//	replyJsonStr, err := json.Marshal(msg)
//	if err != nil {
//		return "err" + err.Error()
//	}
//	return string(replyJsonStr)
//}
