package constant

// header
const (
	TRACE_ID      = "X-TRACE-ID"
	AUTHORIZATION = "Authorization"
	USERNAME      = "Username"
)

// trx type
const (
	AYOCONNECT = "AYOCONNECT"
)

// Mask Log Sensitive Key
var SENSITIVE_KEY []string = []string{}

// log
const (
	LOG_MAX_LENGTH  = 2000
	LOG_TRUNCATED   = "-----truncated"
	LOG_HIDE        = "[hide]"
	LOG_DATE_FORMAT = "2006/01/02 15:04:05"
	LOG_STARTUP     = "START-UP"
)
