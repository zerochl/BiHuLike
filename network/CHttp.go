package network

import (
	"io/ioutil"
	"net/http"
	"log"
	"strings"
	//"fmt"
	"net/url"
	//"fmt"
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

	//fmt.Println(string(body))
	return body,nil
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
