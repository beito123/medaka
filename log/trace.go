package log

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"regexp"
	"runtime"
)

//Thanks: http://sgykfjsm.github.io/blog/2016/01/20/golang-function-tracing/

var regStack = regexp.MustCompile(`^(\S.+)\.(\S.+)$`)

type CallerInfo struct {
	PackageName  string
	FunctionName string
	FileName     string
	FileLine     int
}

func Dump(skip int, count int) (callerInfo []*CallerInfo) {
	for i := 1; i <= count; i++ {
		pc, _, _, ok := runtime.Caller(skip + i)
		if !ok {
			break
		}

		fn := runtime.FuncForPC(pc)
		fileName, fileLine := fn.FileLine(pc)

		_fn := regStack.FindStringSubmatch(fn.Name())
		callerInfo = append(callerInfo, &CallerInfo{
			PackageName:  _fn[1],
			FunctionName: _fn[2],
			FileName:     fileName,
			FileLine:     fileLine,
		})
	}
	return
}
