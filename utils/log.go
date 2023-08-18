package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}
