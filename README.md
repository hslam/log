# codec
[![GoDoc](https://godoc.org/github.com/hslam/log?status.svg)](https://godoc.org/github.com/hslam/log)
[![Build Status](https://travis-ci.org/hslam/log.svg?branch=master)](https://travis-ci.org/hslam/log)
[![codecov](https://codecov.io/gh/hslam/log/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/log)](https://goreportcard.com/report/github.com/hslam/log)
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
```
package main

import (
	"github.com/hslam/log"
)

func main() {
	log.SetPrefix("LogPrefix")
	log.SetLevel(log.DebugLevel)
	log.Debugf("%d %s %t", 1024, "HelloWorld", true)
	log.Tracef("%d %s %t", 1024, "HelloWorld", true)
	log.Allf("%d %s %t", 1024, "HelloWorld", true)
	log.Infof("%d %s %t", 1024, "HelloWorld", true)
	log.Noticef("%d %s %t", 1024, "HelloWorld", true)
	log.Warnf("%d %s %t", 1024, "HelloWorld", true)
	log.Errorf("%d %s %t", 1024, "HelloWorld", true)
	log.Panicf("%d %s %t", 1024, "HelloWorld", true)
	log.Fatalf("%d %s %t", 1024, "HelloWorld", true)
}
```

### Output
<img src="https://raw.githubusercontent.com/hslam/log/master/output.png" width = "500" height = "153" alt="output" align=center>


### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Authors
log was written by Meng Huang.
