package medaka

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

var (
	//Version is the project's version.
	Version string

	//Revision is the project's revision
	Revision string
)

const (
	//Name is the project name.
	Name = "Medaka"

	//APIVersion is Medaka API version.
	//Notice: Medaka doesn't support plugins now.
	APIVersion = "1.0"

	//CodeName is the project's codename.
	CodeName = "Natsu (Summer)"

	//SupportMCBEVersion is supported version in the project
	SupportMCBEVersion = "1.2.x"
)

const (
	TraceLimit = 30 //For debug
)

//Logger is basic logger interface
type Logger interface {
	Info(msg ...interface{})
	Notice(msg ...interface{})
	Warn(msg ...interface{})
	Fatal(msg ...interface{})
	Debug(msg ...interface{})
	Infof(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Err(err error, trace []*CallerInfo)
	Trace(trace []*CallerInfo)
	SetLogDebug(bool)
}

type StdLogger interface {
	Print(...interface{})
	Printf(format string, args ...interface{})
}

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
