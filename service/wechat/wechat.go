package wechat

import (
	"fmt"
	"github.com/mizuki1412/go-core-kit/library/httpkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"jd-mining-server/service/config"
	"net/http"
)

func Push2Wechat(to, msg string) {
	httpkit.Request(httpkit.Req{
		Method: http.MethodPost,
		Url:    configkit.GetStringD(config.WechatApi),
		Header: map[string]string{
			"Cookie": fmt.Sprintf("session=%s", configkit.GetStringD(config.WechatToken)),
		},
		FormData: map[string]string{
			"to":      to,
			"content": msg,
		},
	})
}
