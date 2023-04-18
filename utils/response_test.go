package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetRespError(t *testing.T) {
	msg := "Something went wrong"
	data := map[string]interface{}{
		"key1": "value",
		"key2": 123,
	}

	want := Response{
		Status:  StatusError,
		Message: msg,
		Data:    data,
	}

	got := GetRespError(msg, data)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("GetRespError(%q, %v) = %v, want %v", msg, data, *got, want)
	}
}

func TestGetRespSuccess(t *testing.T) {
	msg := "Operation completed successfully"
	data := map[string]interface{}{
		"result": "success",
	}

	want := Response{
		Status:  StatusSuccess,
		Message: msg,
		Data:    data,
	}

	got := GetRespSuccess(msg, data)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("GetRespSuccess(%q, %v) = %v, want %v", msg, data, *got, want)
	}

	jsonGot, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("Error marshalling response to JSON: %v", err)
	}

	var jsonWant Response
	err = json.Unmarshal(jsonGot, &jsonWant)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON to response: %v", err)
	}

	if !reflect.DeepEqual(jsonWant, want) {
		t.Errorf("JSON unmarshalled from %s does not match expected response: got %v, want %v", jsonGot, jsonWant, want)
	}
}
