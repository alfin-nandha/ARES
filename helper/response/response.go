package response

import (
	"ares/proto"
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
)

const (
	NOT_FOUND_CODE    = 404
	BAD_REQUEST_CODE  = 400
	FORBIDDEN_CODE    = 403
	UNAUTHORIZED_CODE = 401
	SUCCESS_CODE      = 200
	UNDEFINED_CODE    = 500
)

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

const (
	NOT_FOUND_MSG      = "data not found"
	ALREADY_EXIST_MSG  = "data already exist"
	UNAUTHORIZED_MSG   = "unauthorized"
	OK_MSG             = "ok"
	SUCCESS_MSG        = "success"
	INTERNAL_ERROR_MSG = "internal server error"
)

var (
	NOT_FOUND      = new(FAILED, NOT_FOUND_MSG, NOT_FOUND_CODE)
	ALREADY_EXIST  = new(FAILED, ALREADY_EXIST_MSG, FORBIDDEN_CODE)
	UNAUTHORIZED   = new(FAILED, UNAUTHORIZED_MSG, UNAUTHORIZED_CODE)
	INTERNAL_ERROR = new(FAILED, INTERNAL_ERROR_MSG, UNDEFINED_CODE)
	OK             = new(SUCCESS, OK_MSG, SUCCESS_CODE)
)

func HttpResponse(protoRes *proto.Response) (resp map[string]any) {
	marshaller := protojson.MarshalOptions{
		EmitUnpopulated: false, // Shows empty fields like RuleSetVersion
		UseProtoNames:   true,  // Uses camelCase (or snake_case if defined in proto)
		Indent:          "  ",  // Makes it pretty
	}

	b, _ := marshaller.Marshal(protoRes)
	json.Unmarshal(b, &resp)
	return
}

func UNDEFINED(msg string) *proto.Response {
	return &proto.Response{
		Status: &proto.ResponseStatus{
			Code:         "FAILED",
			Description:  msg,
			InternalCode: UNDEFINED_CODE,
		},
	}
}

func BAD_REQUEST(msg string) *proto.Response {
	return &proto.Response{
		Status: &proto.ResponseStatus{
			Code:         "FAILED",
			Description:  msg,
			InternalCode: BAD_REQUEST_CODE,
		},
	}
}

func new(code, msg string, internalCode int) *proto.Response {
	return &proto.Response{
		Status: &proto.ResponseStatus{
			Code:         code,
			Description:  msg,
			InternalCode: int32(internalCode),
		},
	}
}
