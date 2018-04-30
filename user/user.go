package User

import (
	"BiHu/network"
	"net/url"
	"BiHu/api"
	"encoding/json"
	"log"
)
//var loginEntity = &RootEntity{};
var loginEntityMap = make(map[string] RootEntity)

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
	loginResult := network.HttpPostForm(Api.API_LOGIN, url.Values{"phone": {userName}, "password": {password}})
	loginEntity := RootEntity{}
	json.Unmarshal(loginResult, loginEntity)
	loginEntityMap[userName] = loginEntity
	return loginEntity.Data
}

func ClearLoginInfo()  {
	loginEntityMap = make(map[string] RootEntity)
}

type RootEntity struct {
	Data LoginEntity `json:"data"`
}

type LoginEntity struct {
	AccessToken string `json:"accessToken"`
	UserId string `json:"userId"`
}
