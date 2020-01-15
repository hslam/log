# codec
Package log implements logging of multi level.

## Feature
* debug
* trace
* all
* info
* warn
* error
* panic
* fatal
* off
* no log

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
	log.Warnf("%d %s %t", 1024, "HelloWorld", true)
	log.Errorf("%d %s %t", 1024, "HelloWorld", true)
	log.Panicf("%d %s %t", 1024, "HelloWorld", true)
	log.Fatalf("%d %s %t", 1024, "HelloWorld", true)
	log.Offf("%d %s %t", 1024, "HelloWorld", true)
}
```

### Output
```
[LogPrefix][D] 2020/01/15 22:26:50.944903 1024 HelloWorld true
[LogPrefix][T] 2020/01/15 22:26:50.944912 1024 HelloWorld true
[LogPrefix][A] 2020/01/15 22:26:50.944914 1024 HelloWorld true
[LogPrefix][I] 2020/01/15 22:26:50.944916 1024 HelloWorld true
[LogPrefix][W] 2020/01/15 22:26:50.944917 1024 HelloWorld true
[LogPrefix][E] 2020/01/15 22:26:50.944919 1024 HelloWorld true
[LogPrefix][P] 2020/01/15 22:26:50.944920 1024 HelloWorld true
[LogPrefix][F] 2020/01/15 22:26:50.944922 1024 HelloWorld true
[LogPrefix][O] 2020/01/15 22:26:50.944923 1024 HelloWorld true
```

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Authors
log was written by Meng Huang.
