// Run Graylog mock server.
//
// Usage
//   $ graylog-mock-server [--port <port number>] [--log-level debug|info|warn|error|fatal|panic] [--data <data-file-path>]
package main

import (
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver/exec/cmd"
)

func main() {
	cmd.Run()
}
