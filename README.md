# graylog-mock-server

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/suzuki-shunsuke/graylog-mock-server/mockserver)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/graylog-mock-server.svg)](https://github.com/suzuki-shunsuke/graylog-mock-server)
[![GitHub tag](https://img.shields.io/github/tag/suzuki-shunsuke/graylog-mock-server.svg)](https://github.com/suzuki-shunsuke/graylog-mock-server/releases)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/graylog-mock-server/master/LICENSE)

[Graylog](https://www.graylog.org/) v2.5 API mock server (deprecated)

## Install

Download a binary from [the release page](https://github.com/suzuki-shunsuke/go-graylog/releases).

## Usage

```console
$ graylog-mock-server --help
graylog-mock-server - Run Graylog mock server.

USAGE:
   graylog-mock-server [options]

VERSION:
   0.1.0

OPTIONS:
   --port value       port number. If you don't set this option, a free port is assigned and the assigned port number is outputed to the console when the mock server runs.
   --log-level value  the log level of logrus which the mock server uses internally. (default: "info")
   --data value       data file path. When the server runs data of the file is loaded and when data of the server is changed data is saved at the file. If this option is not set, no data is loaded and saved.
   --help, -h         show help
   --version, -v      print the version
```

## License

[MIT](LICENSE)
