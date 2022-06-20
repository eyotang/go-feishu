package feishu

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContactService_BatchGetId(t *testing.T) {
	Convey("test ContactService_BatchGetId", t, func() {
		mux, server, client := setup(t)
		defer teardown(server)

		mux.HandleFunc("/open-apis/auth/v3/tenant_access_token/internal", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"code": 0,
				"expire": 7200,
				"msg": "ok",
				"tenant_access_token": "t-caecc734c2e3328a62489fe0648c4b98779515d3"
			}`)
		})

		mux.HandleFunc("/open-apis/contact/v3/users/batch_get_id?user_id_type=open_id", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{
				"code": 0,
				"data": {
					"user_list": [
						{
							"mobile": "15921667242",
							"user_id": "on_65e9d53c02f8a9ddc4ea67dbffba4d43"
						},
						{
							"email": "2sdljfl3@hypergryph.com"
						},
						{
							"email": "tangyongqiang@hypergryph.com",
							"user_id": "on_65e9d53c02f8a9ddc4ea67dbffba4d43"
						},
						{
							"email": "chenzhida@hypergryph.com",
							"user_id": "on_74907cd57767d1ec0b4e60cfa96a2263"
						}
					]
				},
				"msg": "success"
			}`)
		})

		opt := &BatchGetIdOptions{
			Emails:  []string{"2sdljfl3@hypergryph.com", "tangyongqiang@hypergryph.com", "chenzhida@hypergryph.com"},
			Mobiles: []string{"15921667242"},
		}
		users, _, err := client.Contact.BatchGetId("open_id", opt)
		So(err, ShouldBeNil)
		want := &BatchUsers{
			CodeMsg: CodeMsg{Code: 0, Message: "success"},
		}
		userList := []User{
			{UserId: "on_65e9d53c02f8a9ddc4ea67dbffba4d43", Mobile: "15921667242"},
			{Email: "2sdljfl3@hypergryph.com"},
			{UserId: "on_65e9d53c02f8a9ddc4ea67dbffba4d43", Email: "tangyongqiang@hypergryph.com"},
			{UserId: "on_74907cd57767d1ec0b4e60cfa96a2263", Email: "chenzhida@hypergryph.com"},
		}
		want.Data.UserList = userList
		So(users, ShouldResemble, want)
	})
}
