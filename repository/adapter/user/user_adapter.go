package user

import (
	"context"
	"fmt"
	"net/http"

	appInit "github.com/fauzanmh/olp-admin/init"
	"github.com/fauzanmh/olp-admin/pkg/helper"
	"go.uber.org/zap"
)

type User struct {
	config *appInit.Config
}

// NewProviderUser :nodoc:
func NewProviderUser(config *appInit.Config) UserAdapter {
	return &User{
		config: config,
	}
}

// --- DeleteMember --- //
func (auth *User) DeleteMember(ctx context.Context, id int64) (err error) {
	url := fmt.Sprintf("%s%s/%d", auth.config.Microservice.User.BaseURL, auth.config.Microservice.User.DeleteUser, id)
	httpMethod := http.MethodDelete
	headers := map[string]string{}
	client := helper.APICall{
		URL:    url,
		Method: httpMethod,
		Header: headers,
	}

	// log request
	zap.S().Named("user.delete-member.request").Info(helper.ConstructRequestLog(url, httpMethod, headers, nil))

	response, err := client.CallWithJson(ctx)
	if err != nil {
		return
	}

	// log response
	zap.S().Named("user.delete-member.response").Info(response.Body)

	if response.StatusCode != 200 {
		err = fmt.Errorf("got unsucceful response from user api")
		return
	}

	return
}
