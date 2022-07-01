package util

import (
	"encoding/json"
	"fmt"

	colors "github.com/TwinProduction/go-color"
)

func Trace(prefix string, obj interface{}) {
	bytes, err := (json.MarshalIndent(obj, "", "\t"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(colors.Green, prefix+":", string(bytes), colors.Reset)
}

// Convert objects to string
func Stringify(obj interface{}) string {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}

	return string(objBytes)
}
