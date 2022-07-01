package hello_handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	objBytes, err := json.Marshal(RequestBody{Message: "Hello"})
	if err != nil {
		t.Fatal("Cannot marshal", err.Error())
	}
	request := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewBuffer(objBytes))

	writer := httptest.NewRecorder()
	Handler(writer, request)
	result := writer.Result()
	defer result.Body.Close()

	data, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal("Cannot read", err.Error())
	}

	var messageMap map[string]string
	err = json.Unmarshal(data, &messageMap)
	if err != nil {
		t.Fatal("Error unmarshalling response body to map[string]string")
	}
	message := messageMap["message"]
	if message != "HELLO!!" {
		t.Fatalf("test failed. Expected %v, received %v", "HELLO!!", message)
	} else {
		t.Logf("test passed. Expected %v, received %v", "HELLO!!", message)
	}

}
