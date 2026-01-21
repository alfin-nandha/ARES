package helper

import "encoding/json"

func ObjectToObject(in interface{}, out interface{}) {
	dataByte, _ := json.Marshal(in)
	_ = json.Unmarshal(dataByte, &out)
}
