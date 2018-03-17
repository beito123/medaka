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
	"strconv"
	"strings"
	"time"

	"github.com/beito123/medaka"
)

const DefaultTimestampFormat = "15:04:05" //2006-01-02 15:04:05

//Level is logging level
type Level uint

const (
	LevelInfo Level = iota
	LevelNotice
	LevelWarn
	LevelFatal
	LevelErr
	LevelDebug
)

func (lvl Level) Color() string {
	switch lvl {
	case LevelInfo:
		return White
	case LevelNotice:
		return Aqua
	case LevelWarn:
		return Yellow
	case LevelFatal:
		return DarkRed
	case LevelErr:
		return Red
	case LevelDebug:
		return Gray
	}

	return White
}

func (lvl Level) String() string {
	switch lvl {
	case LevelInfo:
		return "info"
	case LevelNotice:
		return "notice"
	case LevelWarn:
		return "warning"
	case LevelFatal:
		return "fatal"
	case LevelErr:
		return "error"
	case LevelDebug:
		return "debug"
	}

	return "unknown"
}

func NewLogger(out medaka.StdLogger) *consoleLogger {
	return &consoleLogger{
		Out:             out,
		OutLevel:        LevelDebug,
		TimestampFormat: DefaultTimestampFormat,
	}
}

//consoleLogger is ...
type consoleLogger struct {
	Out             medaka.StdLogger
	OutLevel        Level
	TimestampFormat string
}

//Info logs the message with info level
func (log *consoleLogger) Info(msg string) {
	log.Log(LevelInfo, msg)
}

//Notice logs the message with notice level
func (log *consoleLogger) Notice(msg string) {
	log.Log(LevelNotice, msg)
}

//Warn logs the message with warn level
func (log *consoleLogger) Warn(msg string) {
	log.Log(LevelWarn, msg)
}

//Fatal logs the message with fatal level
func (log *consoleLogger) Fatal(msg string) {
	log.Log(LevelFatal, msg)
}

//Debug logs the message with debug level
func (log *consoleLogger) Debug(msg string) {
	log.Log(LevelDebug, msg)
}

//Err logs the error
func (log *consoleLogger) Err(err error, trace []*medaka.CallerInfo) {
	if trace == nil {
		trace = medaka.Dump(1, medaka.TraceLimit)
	}

	e := "Error: " + err.Error()

	log.Log(LevelErr, e)

	log.Trace(trace)
}

func (log *consoleLogger) Trace(trace []*medaka.CallerInfo) {
	if trace == nil {
		trace = medaka.Dump(1, medaka.TraceLimit)
	}

	stack := "Stacktrace:\n"
	for i := 0; i < len(trace); i++ {
		info := trace[i]

		num := strconv.Itoa(i)
		line := strconv.Itoa(info.FileLine)
		stack += "#" + num + " " + info.FileName + "->" + info.FunctionName + "():" + line + "\n"
	}

	log.Debug(stack)
}

func (log *consoleLogger) SetLogDebug(b bool) {
	if b {
		log.OutLevel = LevelDebug
	} else {
		log.OutLevel = LevelErr
	}
}

//Log logs the message with the level
func (log *consoleLogger) Log(level Level, msg string) {
	if level > log.OutLevel {
		return
	}

	color := level.Color()

	timeFormat := DefaultTimestampFormat
	if log.TimestampFormat != "" {
		timeFormat = log.TimestampFormat
	}

	time := time.Now().Format(timeFormat)
	lvl := strings.ToUpper(level.String())

	text := color + "[" + time + "] [" + lvl + "] " + msg + Reset

	log.Out.Print(text)
}
