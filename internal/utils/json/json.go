package json

import (
	"encoding/json"
	"fmt"
)

func ToColorJson(obj interface{}) string {
	if obj == nil {
		return ""
	}

	str, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}
