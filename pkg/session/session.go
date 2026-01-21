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
	Map             Map.ConcurrentMap
	Logger          *logger.Logger
	RequestTime     time.Time
	Method          string
	Url             string
	TraceId         string
	AppName         string
	Header, Request any
	Ctx             context.Context
}

func New() *Session {
	return &Session{
		RequestTime: time.Now(),
		Logger:      logger.Log,
		Map:         Map.New(),
		Ctx:         context.Background(),
	}
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

func (session *Session) LogDatabase(sql string, rows int64, error any, timeDuration time.Duration) {
	session.Logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.String("SQL", sql),
		zap.Any("Rows", rows),
		zap.Any("Error", error),
		zap.Any("ProcessTime", fmt.Sprintf("%d ms", timeDuration/time.Millisecond)),
	)
}

func (session *Session) LogRequest(message ...any) {
	session.Logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.String("Method", session.Method),
		zap.String("URI", session.Url),
		zap.Any("Request", session.Logger.MaskData(session.Request)),
		zap.Any("Header", session.Header),
		zap.String("Message", formatResponse(message...)),
	)
}

func (session *Session) LogResponse(timeDuration time.Duration, code int, url string, method string, response any) {
	session.Logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.String("Method", method),
		zap.String("Url", url),
		zap.Int("HttpStatus", code),
		zap.Any("Response", response),
		zap.String("ProcessTime", fmt.Sprintf("%d ms", timeDuration/time.Millisecond)),
	)
}

func (session *Session) LogInfo(message any, data ...any) {
	session.Logger.Info("INFO",
		zap.String("TraceId", session.TraceId),
		zap.Any("Message", message),
		zap.Any("Data", data),
	)
}

func (session *Session) LogError(message any, data any) {
	session.Logger.Info("Error",
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
