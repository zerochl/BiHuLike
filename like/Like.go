package Like

import (
	"BiHu/follow"
	"net/url"
	"strconv"
	"BiHu/user"
	"BiHu/api"
	"BiHu/network"
)

func DoLike(artEntityList []Follow.ArtEntity,loginEntity User.LoginEntity)  {
	if (artEntityList == nil || len(artEntityList) == 0) {
		return
	}
	for _,artEntity := range artEntityList {
		network.HttpPostForm(Api.API_LIKE,
			url.Values{"userId": {loginEntity.UserId}, "accessToken": {loginEntity.AccessToken}, "artId": {strconv.Itoa(artEntity.ID)}})
	}
}
