# codec
[![GoDoc](https://godoc.org/github.com/hslam/log?status.svg)](https://godoc.org/github.com/hslam/log)
[![Build Status](https://github.com/hslam/log/workflows/build/badge.svg)](https://github.com/hslam/log/actions)
[![codecov](https://codecov.io/gh/hslam/log/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/log?v=7e100)](https://goreportcard.com/report/github.com/hslam/log)
[![GitHub release](https://img.shields.io/github/release/hslam/log.svg)](https://github.com/hslam/log/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/log.svg?style=flat-square)](https://github.com/hslam/log/blob/master/LICENSE)

Package log implements multilevel logging.

## Feature
* debug
* trace
* all
* info
* notice
* warn
* error
* panic
* fatal
* off

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
	logger.SetMicroseconds(false)
	logger.Debugf("%d %s %t", 1024, "HelloWorld", true)
	logger.Tracef("%d %s %t", 1024, "HelloWorld", true)
	logger.Allf("%d %s %t", 1024, "HelloWorld", true)
	logger.Infof("%d %s %t", 1024, "HelloWorld", true)
	logger.Noticef("%d %s %t", 1024, "HelloWorld", true)
	logger.Warnf("%d %s %t", 1024, "HelloWorld", true)
	logger.Errorf("%d %s %t", 1024, "HelloWorld", true)
	logger.Panicf("%d %s %t", 1024, "HelloWorld", true)
	logger.Fatalf("%d %s %t", 1024, "HelloWorld", true)
}
```

### Output
<img src="https://raw.githubusercontent.com/hslam/log/master/output.png" width = "382" height = "136" alt="output" align=center>


### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Author
log was written by Meng Huang.
