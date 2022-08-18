package feishu

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAppService_SendAppCardMessage(t *testing.T) {
	Convey("test AppService_SendAppCardMessage", t, func() {
		mux, server, client := setup(t)
		defer teardown(server)

		mux.HandleFunc("/open-apis/im/v1/messages", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"code": 0,
				"message": "ok"
			}`)
		})

		opt := &AppCardMessageOption{
			MsgType:   "interactive",
			ReceiveID: "ou_b46ad73aaaqer1231bd1daeb7d3a41e9f1a",
			Card: AppCardOption{
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
		rsp, _, err := client.App.SendAppCardMessage("open_id", opt)
		So(err, ShouldBeNil)
		want := &CodeMsg{
			Code: 0, Message: "ok",
		}
		So(rsp, ShouldResemble, want)
	})
}
