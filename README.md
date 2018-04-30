# BiHuLike
币乎点赞，懂的人会懂
# 使用方法
命令行执行：./BiHuLike -c config.json > bihu.log 2>&1 &
config.json
```json
{
  //followNameList 关注的用户名，与点赞限制条件，如果此人的点赞已经高于给定的值则放弃
  "followNameList" : [
    {
      "name" : "南宫远",
      "limitLikeCount" : 300
    },{
      "name" : "圊呓语",
      "limitLikeCount" : 300
    },{
      "name" : "区块链生存指南",
      "limitLikeCount" : 300
    },{
      "name" : "数字货币趋势狂人",
      "limitLikeCount" : 300
    },{
      "name" : "区块佣兵",
      "limitLikeCount" : 300
    },{
      "name" : "湘乡的大树",
      "limitLikeCount" : 300
    },{
      "name" : "吴庆英",
      "limitLikeCount" : 300
    },{
      "name" : "JIMI",
      "limitLikeCount" : 300
    },{
      "name" : "wdctll",
      "limitLikeCount" : 300
    },{
      "name" : "区块链研究院",
      "limitLikeCount" : 300
    },{
      "name" : "玩火的猴子",
      "limitLikeCount" : 300
    },{
      "name" : "孤独的异客",
      "limitLikeCount" : 300
    },{
      "name" : "币圈少年",
      "limitLikeCount" : 300
    },{
      "name" : "EOSCannon",
      "limitLikeCount" : 300
    },{
      "name" : "Bean",
      "limitLikeCount" : 300
    },{
      "name" : "柚子",
      "limitLikeCount" : 300
    },{
      "name" : "1212秋刀鱼",
      "limitLikeCount" : 300
    },{
      "name" : "珞珈山神",
      "limitLikeCount" : 300
    },{
      "name" : "鱼夫子说",
      "limitLikeCount" : 300
    },{
      "name" : "郭立芳37258",
      "limitLikeCount" : 300
    },{
      "name" : "庖丁解币",
      "limitLikeCount" : 300
    },{
      "name" : "钱串串",
      "limitLikeCount" : 300
    },{
      "name" : "胖哥",
      "limitLikeCount" : 300
    }
  ],
  //总点赞数限制
  "limitLikeCount" : 150,
  //是否需要开启间隔刷新
  "needRefreshInterval" : true,
  //间隔刷新的时间，单位秒
  "refreshInterval" : 20,
  //快速刷新的时间，单位秒，快速刷新出现在固定时间到点了触发
  "fastRefreshInterval" : 2,
  //快速刷新的次数限制
  "fastRefreshCount" : 40,
  //固定时间点列表
  "fixedRefreshTimeList" : [
    "08:00",
    "05:55",
    "06:30",
    "06:46",
    "06:59",
    "07:07",
    "07:21",
    "08:08",
    "09:09",
    "10:00",
    "10:10",
    "11:29",
    "16:16",
    "20:00",
    "20:05",
    "20:35",
    "21:35",
    "21:45",
    "21:50",
    "22:00"
  ],
  //登录状态刷新时间，单位小时，12小时清除用户token
  "loginRefreshTime" : 12,
  //以此用户来获取用户
  "phone" : "18051158888",
  "password" : "自行抓包获取，因为加密了",
  //同时绑定其他用户账号，获取到可点赞的列表的同时一次性处理其他账号
  "userList" : [
    {
      "phone" : "18051188888",
      "password" : "自行抓包获取，因为加密了"
    }
  ]
}

```
