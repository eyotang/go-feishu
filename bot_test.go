package feishu

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBotService_SendBotCardMessage(t *testing.T) {
	Convey("test BotService_SendBotCardMessage", t, func() {
		mux, server, client := setup(t)
		defer teardown(server)

		mux.HandleFunc("/open-apis/bot/v2/hook/891105b7-1234-4567-7890-c4c372235090", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"StatusCode": 0,
				"StatusMessage": "success"
			}`)
		})

		botKey := "891105b7-1234-4567-7890-c4c372235090"
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
		rsp, _, err := client.Bot.SendBotCardMessage(botKey, "", opt)
		So(err, ShouldBeNil)
		want := &BotResponse{
			StatusCode: 0, StatusMessage: "success",
		}
		So(rsp, ShouldResemble, want)
	})
}

func TestBotService_SendBotCardMessageInvalid(t *testing.T) {
	Convey("test BotService_SendBotCardMessageInvalid", t, func() {
		client, _ := NewClient()
		botKey := "891105b7-1234-4567-7890-c4c372235090"
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
		rsp, _, err := client.Bot.SendBotCardMessage(botKey, "", opt)
		So(err, ShouldBeNil)
		want := &ErrorMessage{
			CodeMsg: CodeMsg{
				Code:    19001,
				Message: "param invalid: incoming webhook access token invalid",
			},
			Data: map[string]interface{}{},
		}
		So(rsp, ShouldResemble, want)
	})
}

func TestBotService_SendBotCardMessageReport(t *testing.T) {
	Convey("test BotService_SendBotCardMessageReport", t, func() {
		client, _ := NewClient()
		// TODO: correct the bot key and secret, then run the test
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
		So(err, ShouldBeNil)
		want := &ErrorMessage{
			CodeMsg: CodeMsg{
				Code:    0,
				Message: "",
			},
			//Data: map[string]interface{},
		}
		So(rsp, ShouldResemble, want)
	})
}
