# go-feishu
A fast [feishu](https://open.feishu.cn/) development sdk written in Golang

```shell script
go get github.com/eyotang/go-feishu
```
```go
// 创建 飞书 客户端
client, _ := feishu.NewClient()

// 创建 飞书 客户端，本地缓存AccessToken
cacheClient, _ := feishu.NewLocalCacheClient(appId, appSecret)

// 创建 飞书 客户端，Redis缓存AccessToken
cacheClient, _ := NewRedisCacheClient(appId, appSecret, redisAddr, redisPasswd)

// 通过手机号或邮箱获取用户 ID
opt := &feishu.BatchGetIdOptions{
			Emails:  []string{"2sdljfl3@a.com", "tangyongqiang@b.com", "chenzhida@c.com"},
			Mobiles: []string{"15921667321"},
		}
users, _, err := cacheClient.Contact.BatchGetId(opt)

// 通过客户端机器人，发送消息
botKey := "2343205b7-1238-2345-8828-c4c3720098090"
secret := "lSkesfsfVdcCfmpdRLSvBb"
opt := &BotCardMessageOption{
	MsgType: "interactive",
	Card: BotCardOption{
		Config: CardConfigOption{
			WideScreenMode: true, EnableForward: true,
		},
		Header: HeadOption{
			Title: TitleOption{
				Tag:     "plain_text",
				Content: "test report",
			},
			Template: "green",
		},
		Elements: []interface{}{
			TitleOption{Tag: "markdown", Content: "abc"},
		},
	},
}
rsp, _, err := client.Bot.SendBotCardMessage(botKey, secret, opt)
```