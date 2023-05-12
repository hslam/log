# log
[![GoDoc](https://godoc.org/github.com/hslam/log?status.svg)](https://godoc.org/github.com/hslam/log)
[![Build Status](https://github.com/hslam/log/workflows/build/badge.svg)](https://github.com/hslam/log/actions)
[![codecov](https://codecov.io/gh/hslam/log/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/log?v=7e100)](https://goreportcard.com/report/github.com/hslam/log)
[![GitHub release](https://img.shields.io/github/release/hslam/log.svg)](https://github.com/hslam/log/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/log.svg?style=flat-square)](https://github.com/hslam/log/blob/master/LICENSE)

Package log implements multilevel logging.

## Feature
* Custom prefix
* Multiple levels
* Highlight color
* File line
* Call stack

## Level
* Debug
* Trace
* All
* Info
* Notice
* Warn
* Error
* Panic
* Fatal
* Off

## Get started

### Install
```
go get github.com/hslam/log
```
### Import
```
import "github.com/hslam/log"
```
### Usage
#### Example
```go
package main

import (
	"github.com/hslam/log"
)

func main() {
	logger := log.New()
	logger.SetPrefix("LogPrefix")
	logger.SetLevel(log.DebugLevel)
	logger.SetHighlight(true)
	logger.SetLine(true)

	logger.Assertf(1 == 1, "%d %s %t", 1024, "HelloWorld", true)

	logger.Debugf("%d %s %t", 1024, "HelloWorld", true)
	logger.Tracef("%d %s %t", 1024, "HelloWorld", true)
	logger.Allf("%d %s %t", 1024, "HelloWorld", true)
	logger.Infof("%d %s %t", 1024, "HelloWorld", true)
	logger.Noticef("%d %s %t", 1024, "HelloWorld", true)
	logger.Warnf("%d %s %t", 1024, "HelloWorld", true)
	logger.Errorf("%d %s %t", 1024, "HelloWorld", true)
	logger.Fatalf("%d %s %t", 1024, "HelloWorld", true)
}
```

### Output

```
[LogPrefix] [2023/05/12 20:51:06.249 +08:00] [DEBUG] [main.go:16] ["1024 HelloWorld true"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [TRACE] [main.go:17] ["1024 HelloWorld true"] [stack="main.main\n\t/Users/huangmeng/go/src/github.com/hslam/tmp/log/main.go:17\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [ALL] [main.go:18] ["1024 HelloWorld true"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [INFO] [main.go:19] ["1024 HelloWorld true"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [NOTICE] [main.go:20] ["1024 HelloWorld true"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [WARN] [main.go:21] ["1024 HelloWorld true"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [ERROR] [main.go:22] ["1024 HelloWorld true"] [stack="main.main\n\t/Users/huangmeng/go/src/github.com/hslam/tmp/log/main.go:22\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"]
[LogPrefix] [2023/05/12 20:51:06.250 +08:00] [FATAL] [main.go:23] ["1024 HelloWorld true"] [stack="main.main\n\t/Users/huangmeng/go/src/github.com/hslam/tmp/log/main.go:23\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:250"]
```

#### Highlight
<img src="https://raw.githubusercontent.com/hslam/log/master/example.png" width = "948" height = "113" alt="example" align=center>

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Author
log was written by Meng Huang.
