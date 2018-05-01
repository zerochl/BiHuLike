package Follow

import (
	"BiHu/network"
	"net/url"
	"encoding/json"
	"BiHu/user"
	"BiHu/api"
	"BiHu/config"
)

func GetFollowList(loginEntity User.LoginEntity) []ArtEntity {
	//log.Print("userId:",loginEntity.UserId,";token:",loginEntity.AccessToken)
	followListByte , err:= network.HttpPostForm(Api.API_GET_FOLLOW_ARTICAL, url.Values{"userId": {loginEntity.UserId}, "accessToken": {loginEntity.AccessToken}})
	//log.Print("followList:",string(followListByte))
	if err != nil {
		return nil
	}
	followEntity := &FollowEntity{}
	json.Unmarshal(followListByte, followEntity)
	return followEntity.Data.ArtList.ArtEntityList
}

func GetStarFollowList(loginEntity User.LoginEntity, configEntity *config.Config) []ArtEntity {
	artEntityList := GetFollowList(loginEntity)
	if (artEntityList == nil) {
		return nil
	}
	//log.Print("userId:",loginEntity.UserId,";token:",loginEntity.AccessToken)
	//log.Print("artEntityList:" , len(artEntityList))
	starArtList := make([]ArtEntity,0)
	for _,followEntity := range configEntity.FollowNameList {
		for _,artEntity := range artEntityList {
			if (artEntity.UserName == followEntity.Name && artEntity.UP == 0 && artEntity.UPS < followEntity.LimitLikeCount) {
				starArtList = append(starArtList, artEntity)
			}
		}
	}
	return starArtList
}

type FollowEntity struct {
	Data DataEntity `json:"data"`
}

type DataEntity struct {
	ArtList ArtListEntity `json:"artList"`
}

type ArtListEntity struct {
	ArtEntityList []ArtEntity `json:"list"`
}

type ArtEntity struct {
	ID int `json:"id"`
	Follow int `json:"follow"`
	UP int `json:"up"`
	UserName string `json:"userName"`
	UPS int `json:"ups"`
}

