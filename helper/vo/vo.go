package vo

import (
	"ares/pkg/session"
	"context"
)

const (
	TraceId    = "traceid"
	Auth       = "auth"
	AppSession = "App_Session"
)

type ApplicationContext struct {
	context.Context
	Session *session.Session
}

func Parse(c context.Context) *ApplicationContext {
	data := c.Value(AppSession)
	session := data.(session.Session)
	return &ApplicationContext{Context: c, Session: &session}
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
