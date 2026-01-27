package http

import (
	"fmt"
	"strconv"

	Container "ares/container"
	"ares/helper/response"
	"ares/helper/vo"
	"ares/pkg/session"

	"runtime"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func middlewareSetup(app *fiber.App, container *Container.Presenter) {
	app.Use(sessionMiddleware)
	app.Use(logMiddleware)
	app.Use(recoveryMiddleware)
	app.Use(cors.New())
}

func recoveryMiddleware(c fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			session := vo.Parse(c.Context()).Session

			// Capture the stack trace
			stackTrace := make([]byte, 1024)
			runtime.Stack(stackTrace, false)

			session.LogError("PANIC", fmt.Sprintf("Recovered from panic: %v\nStack Trace: %s", r, stackTrace))
			c.JSON(response.UNDEFINED(fmt.Sprintf("%v", r)))
		}
	}()
	return c.Next()
}

// func validateJwt(c *fiber.Ctx) (err error) {
// 	rawJwt := c.Get(Constant.AUTHORIZATION)
// 	splittedJwt := strings.Split(rawJwt, " ")
// 	if len(splittedJwt) < 2 {
// 		return c.JSON(response.UNAUTHORIZED)
// 	}
// 	claims, err := token.ValidateToken(splittedJwt[1])
// 	if err != nil {
// 		return c.JSON(response.UNAUTHORIZED)
// 	}

// 	c.Set(Constant.USERNAME, claims.Username)
// 	return c.Next()
// }

func sessionMiddleware(c fiber.Ctx) error {
	traceId := vo.GenerateTraceId()

	headers := c.GetReqHeaders()
	headersTraceId := headers[vo.TraceId]
	if len(headersTraceId) != 0 {
		traceId = headersTraceId[0]
	}

	// campact json requestBody
	bodyMap := map[string]any{}
	bodyByte := c.Body()
	err := json.Unmarshal(bodyByte, &bodyMap)
	if err == nil {
		bodyByte, _ = json.Marshal(bodyMap)
	}

	newSess := session.New().
		SetTraceId(traceId).
		SetRequest(bodyByte).
		SetURL(c.OriginalURL()).
		SetHeader(headers).
		SetQuery(c.Queries()).
		SetMethod(c.Method())

	c.Locals(vo.AppSession, *newSess)

	return c.Next()
}

func logMiddleware(c fiber.Ctx) error {
	start := time.Now()
	session := c.Locals(vo.AppSession).(session.Session)

	// log request
	session.LogRequest(nil)

	// Proceed with the request
	code := strconv.Itoa(c.Response().StatusCode())
	msg := ""
	err := c.Next()
	if err != nil {
		msg = err.Error()
	}

	// Log the response
	session.LogResponse(time.Since(start), code, string(c.Response().Body()), msg)

	return err
}

func ErrorHandler(c fiber.Ctx, err error) error {
	return c.JSON(err.Error())
}
