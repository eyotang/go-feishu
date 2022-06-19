package feishu

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccessTokenManagerService_GetAccessToken(t *testing.T) {
	Convey("test AccessTokenManagerService_GetAccessToken", t, func() {
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

		err, accessToken := client.accessTokenManager.GetAccessToken()
		So(err, ShouldBeNil)
		want := "t-caecc734c2e3328a62489fe0648c4b98779515d3"
		So(accessToken, ShouldEqual, want)

		// request again, debug with break points.
		time.Sleep(1 * time.Second)
		err, accessToken = client.accessTokenManager.GetAccessToken()
		So(err, ShouldBeNil)
		want = "t-caecc734c2e3328a62489fe0648c4b98779515d3"
		So(accessToken, ShouldEqual, want)
	})
}
