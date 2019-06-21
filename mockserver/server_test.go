package mockserver_test

import (
	"testing"

	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver"
)

func TestNewServer(t *testing.T) {
	server, err := mockserver.NewServer("", nil)
	if err != nil {
		t.Fatal(err)
	}
	if server == nil {
		t.Fatal("server is nil")
	}
	server.Start()
	defer server.Close()
	if server.Endpoint() == "" {
		t.Fatal("endpoint is empty")
	}
}
