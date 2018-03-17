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
	"github.com/beito123/medaka/log"
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
	CodeName = "The egg"

	//SupportMCBEVersion is supported version in the project
	SupportMCBEVersion = "1.2.0"
)

//Logger is basic logger interface
type Logger interface {
	Info(msg string)
	Notice(msg string)
	Warn(msg string)
	Fatal(msg string)
	Debug(msg string)
	Err(err error, trace []*log.CallerInfo)
	Trace(trace []*log.CallerInfo)
	SetLogDebug(bool)
}
