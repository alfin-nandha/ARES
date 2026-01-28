package session

import (
	"ares/pkg/logger"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"reflect"
	"strings"
	"time"

	JsonIter "github.com/json-iterator/go"
	Map "github.com/orcaman/concurrent-map"
)

type Session struct {
	Map                              Map.ConcurrentMap
	logger                           *logger.Logger
	ClientId                         int64
	RequestTime                      time.Time
	Method                           string
	Url                              string
	TraceId                          string
	AppName                          string
	Header, Query, Request, MetaData any
	Ctx                              context.Context
}

func New() *Session {
	return &Session{
		RequestTime: time.Now(),
		logger:      logger.Log,
		Map:         Map.New(),
		Ctx:         context.Background(),
	}
}

func (session *Session) SetClientId(clientId int64) *Session {
	session.ClientId = clientId
	return session
}
func (session *Session) SetTraceId(traceId string) *Session {
	session.TraceId = traceId
	return session
}

func (session *Session) SetMethod(method string) *Session {
	session.Method = method
	return session
}

func (session *Session) SetAppName(appName string) *Session {
	session.AppName = appName
	return session
}

func (session *Session) SetURL(url string) *Session {
	session.Url = url
	return session
}

func (session *Session) SetHeader(header any) *Session {
	session.Header = header
	return session
}
func (session *Session) SetQuery(query any) *Session {
	session.Query = query
	return session
}

func (session *Session) SetMetaData(md any) *Session {
	session.MetaData = md
	return session
}

func (session *Session) SetRequest(request any) *Session {
	session.Request = request
	return session
}

func (session *Session) Get(key string) (data any, err error) {
	data, ok := session.Map.Get(key)
	if !ok {
		err = errors.New("NotFound")
	}
	return
}

func (session *Session) Put(key string, data any) {
	session.Map.Set(key, data)
}

func (session *Session) String() string {
	b, _ := json.Marshal(session)
	return string(b)
}

func (session *Session) LogDatabase(sql string, rows int64, error any, timeDuration time.Duration) {
	session.logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.String("SQL", sql),
		zap.Any("Rows", rows),
		zap.Any("Error", error),
		zap.Any("ProcessTime", fmt.Sprintf("%d ms", timeDuration/time.Millisecond)),
	)
}

func (session *Session) LogRequest(message ...any) {
	fields := []zap.Field{
		zap.String("TraceId", session.TraceId),
		zap.String("Method", session.Method),
		zap.String("URI", session.Url),
		zap.Any("Request", session.logger.MaskData(session.Request)),
		zap.Any("Message", message),
	}
	if session.Header != nil {
		fields = append(fields, zap.Any("Header", session.Header))
	}
	if session.MetaData != nil {
		fields = append(fields, zap.Any("MetaData", session.MetaData))
	}
	if session.Query != nil {
		fields = append(fields, zap.Any("Query", session.Query))
	}
	session.logger.Info("INFO", fields...)
}

func (session *Session) LogResponse(timeDuration time.Duration, code string, response any, msg string) {
	session.logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.String("Method", session.Method),
		zap.String("URI", session.Url),
		zap.String("Code", code),
		zap.Any("Response", response),
		zap.Any("Message", msg),
		zap.String("ProcessTime", fmt.Sprintf("%d ms", timeDuration/time.Millisecond)),
	)
}

func (session *Session) LogInfo(message any, data ...any) {
	session.logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.Any("Message", message),
		zap.Any("Data", data),
	)
}

func (session *Session) LogError(message any, data any) {
	session.logger.Info("Error",
		zap.String("TraceId", session.TraceId),
		zap.Any("Message", message),
		zap.Any("Data", data),
	)
}

var json = JsonIter.ConfigCompatibleWithStandardLibrary

func formatResponse(message ...any) string {
	sb := strings.Builder{}

	for _, msg := range message {
		var m []byte
		if reflect.ValueOf(msg).Kind().String() == "string" {
			m = []byte(msg.(string))
		} else {
			m, _ = json.Marshal(msg)
		}

		sb.Write(m)
	}

	return sb.String()
}
