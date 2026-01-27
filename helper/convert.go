package helper

import "encoding/json"

func ObjectToObject(in any, out any) {
	dataByte, _ := json.Marshal(in)
	_ = json.Unmarshal(dataByte, &out)
}
