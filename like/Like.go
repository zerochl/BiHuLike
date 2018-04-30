package Like

import (
	"bihu/follow"
	"bihu/network"
	"bihu/api"
	"net/url"
	"bihu/user"
	"strconv"
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
