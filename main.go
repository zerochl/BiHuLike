package main

import (
	"BiHu/config"
	"log"
	"time"
	"BiHu/user"
	"BiHu/follow"
	"BiHu/like"
)

func main()  {
	//读取配置
	configEntity := config.GetConfig()
	log.Print("config:",configEntity.RefreshInterval)
	ch := make(chan bool)
	//开始间隔任务
	go doIntervalTask(configEntity)
	//开始定时任务
	go doTimingTask(configEntity)
	//清楚token
	go doClearToken(configEntity)
	end := <- ch
	log.Print("end:",end)
}

func doIntervalTask(configEntity *config.Config)  {
	if (configEntity.RefreshInterval == 0 || !configEntity.NeedRefreshInterval) {
		return
	}
	for i:= 0;;i++{
		doTask(configEntity);
		time.Sleep(time.Second * time.Duration(configEntity.RefreshInterval))
	}
}

func doTimingTask(configEntity *config.Config)  {
	for i := 0;;i++ {
		time.Sleep(time.Second * 1)
		for _,customTime := range configEntity.FixedRefreshTimeList {
			//if (customTime == time.Date(2014, 1, 7, 6, 30, 4, 0, time.Local).Format("15:04")) {
			//	log.Print("符合条件，开始fast refresh")
			//	fastRequest(configEntity)
			//}
			if (customTime == time.Now().Format("15:04")) {
				log.Print("符合条件，开始fast refresh")
				fastRequest(configEntity)
			}
		}
	}
}

func doClearToken(configEntity *config.Config)  {
	//每天清除一次token
	for i := 0;;i++ {
		time.Sleep(time.Hour * time.Duration(configEntity.LoginRefreshTime))
		//time.Sleep(time.Second * time.Duration(4))
		User.ClearLoginInfo()
	}
}

func doTask(configEntity *config.Config)  {
	//network.HttpPostForm(API_LOGIN, url.Values{"phone": {"18051156285"}, "password": {"076742567778be885ad66804ec9facb21a4296aa41ac29c4f7d5afe7a206c699"}})
	loginEntity := User.GetLoginEntity(configEntity.Phone, configEntity.Password)
	//log.Print("userId:",loginEntity.UserId,";token:",loginEntity.AccessToken)
	//log.Print(loginEntity.AccessToken,";userId:",loginEntity.UserId)
	followArtList := Follow.GetStarFollowList(loginEntity,configEntity)
	log.Print("size:" , len(followArtList))
	if (followArtList == nil || len(followArtList) == 0) {
		return;
	}
	Like.DoLike(followArtList, loginEntity)
	if (configEntity.UserList == nil || len(configEntity.UserList) == 0) {
		return
	}
	for _,userInfo := range configEntity.UserList {
		loginEntity = User.GetLoginEntity(userInfo.Phone, userInfo.Password)
		Like.DoLike(followArtList, loginEntity)
	}
}

func fastRequest(configEntity *config.Config)  {
	if (configEntity.FastRefreshInterval == 0) {
		return
	}
	for i:= 0;i <= configEntity.FastRefreshCount;i++{
		doTask(configEntity);
		time.Sleep(time.Second * time.Duration(configEntity.FastRefreshInterval))
	}
}



