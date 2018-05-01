package User

import (
	"BiHu/network"
	"net/url"
	"BiHu/api"
	"encoding/json"
	"log"
)
//var loginEntity = &RootEntity{};
var loginEntityMap = make(map[string] *RootEntity)

func GetLoginEntity(userName,password string) LoginEntity {
	//if (loginEntity.Data.AccessToken != "") {
	//	return loginEntity.Data
	//}
	//log.Print("开始重新获取token")
	//loginResult := network.HttpPostForm(Api.API_LOGIN, url.Values{"phone": {userName}, "password": {password}})
	//json.Unmarshal(loginResult, loginEntity)
	//return loginEntity.Data
	_,ok := loginEntityMap[userName]
	if (ok) {
		return loginEntityMap[userName].Data
	}
	log.Print("开始重新获取token")
	loginResult, err := network.HttpPostForm(Api.API_LOGIN, url.Values{"phone": {userName}, "password": {password}})
	if (err != nil) {
		log.Print("获取token失败，开始循环获取")
		return GetLoginEntity(userName,password)
	}
	//log.Print("token:",string(loginResult))
	loginEntity := &RootEntity{}
	json.Unmarshal(loginResult, loginEntity)
	loginEntityMap[userName] = loginEntity
	//log.Print("token:",loginEntity.Data.AccessToken)
	return loginEntity.Data
}

func ClearLoginInfo()  {
	loginEntityMap = make(map[string] *RootEntity)
}

type RootEntity struct {
	Data LoginEntity `json:"data"`
}

type LoginEntity struct {
	AccessToken string `json:"accessToken"`
	UserId string `json:"userId"`
}
