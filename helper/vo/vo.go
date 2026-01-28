package vo

import (
	"ares/pkg/session"
	"context"
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

const (
	TraceId    = "traceid"
	Auth       = "Authorization"
	AppSession = "App_Session"
)

type BasicAuth struct {
	Username string
	Password string
	IsValid  bool
}

type ApplicationContext struct {
	context.Context
	Session *session.Session
}

func GenerateTraceId() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}

func Parse(c context.Context) *ApplicationContext {
	data := c.Value(AppSession)
	session := data.(session.Session)
	return &ApplicationContext{Context: c, Session: &session}
}

func AuthDecode(auth string) (basicAuth BasicAuth) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}

	payload, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return
	}

	return BasicAuth{
		Username: pair[0],
		Password: pair[1],
		IsValid:  true,
	}
}

// func (c *ApplicationContext) BuildResponse(data interface{}) error {
// 	c.Session.LogFull(data)
// 	return c.Context.JSON(http.StatusOK, data)
// }

// func (c *ApplicationContext) Response(code int, status string, message string, data interface{}) error {
// 	if data == nil {
// 		data = struct{}{}
// 	}

// 	response := Response.Response{
// 		Code:    code,
// 		Status:  status,
// 		Message: message,
// 		Data:    data,
// 	}

// 	c.Session.LogFull(response)
// 	return c.Context.JSON(http.StatusOK, response)
// }
