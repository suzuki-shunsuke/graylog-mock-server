package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver/handler"
)

func TestNewAPIError(t *testing.T) {
	e := handler.NewAPIError("test")
	if e.Type != "ApiError" {
		t.Fatalf(`e.Type = "%s", wanted "ApiError"`, e.Type)
	}
	if e.Message != "test" {
		t.Fatalf(`e.Message = "%s", wanted "test"`, e.Message)
	}
}
