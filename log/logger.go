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
	"log"
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

func NewLogger(out medaka.StdLogger) *Logger {
	return &Logger{
		Out:             out,
		OutLevel:        LevelDebug,
		TimestampFormat: DefaultTimestampFormat,
	}
}

func NewStdLogger() *log.Logger {
	return log.New(&LoggerWriter{
		Level: LevelInfo,
	}, "", 0)
}

// Logger is basic console logger
type Logger struct {
	Out             medaka.StdLogger
	OutLevel        Level
	TimestampFormat string
	Prefix          string
}

// Print logs the message
func (log *Logger) Print(msg ...interface{}) {
	log.Log(LevelInfo, msg...)
}

// Printf logs the message with format
func (log *Logger) Printf(format string, args ...interface{}) {
	log.Logf(LevelInfo, format, args...)
}

// Info logs the message as info level
func (log *Logger) Info(msg ...interface{}) {
	log.Log(LevelInfo, msg...)
}

// Infof logs the message as info level with format
func (log *Logger) Infof(format string, args ...interface{}) {
	log.Logf(LevelInfo, format, args...)
}

// Notice logs the message as notice level
func (log *Logger) Notice(msg ...interface{}) {
	log.Log(LevelNotice, msg...)
}

// Noticef logs the message as notice level with format
func (log *Logger) Noticef(format string, args ...interface{}) {
	log.Logf(LevelNotice, format, args...)
}

// Warn logs the message as warn level
func (log *Logger) Warn(msg ...interface{}) {
	log.Log(LevelWarn, msg...)
}

// Warnf logs the message as warn level with format
func (log *Logger) Warnf(format string, args ...interface{}) {
	log.Logf(LevelWarn, format, args...)
}

// Fatal logs the message as fatal level
func (log *Logger) Fatal(msg ...interface{}) {
	log.Log(LevelFatal, msg...)
}

// Fatalf logs the message as fatal level with format
func (log *Logger) Fatalf(format string, args ...interface{}) {
	log.Logf(LevelFatal, format, args...)
}

// Debug logs the message as debug level
func (log *Logger) Debug(msg ...interface{}) {
	log.Log(LevelDebug, msg...)
}

// Debugf logs the message as debug level with format
func (log *Logger) Debugf(format string, args ...interface{}) {
	log.Logf(LevelDebug, format, args...)
}

// Err logs the error
func (log *Logger) Err(err error, trace []*medaka.CallerInfo) {
	if trace == nil {
		trace = medaka.Dump(1, medaka.TraceLimit)
	}

	e := "Error: " + err.Error()

	log.Log(LevelErr, e)

	log.Trace(trace)
}

// Trace logs traces
func (log *Logger) Trace(trace []*medaka.CallerInfo) {
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

// SetLogDebug set log level to debug mode
func (log *Logger) SetLogDebug(b bool) {
	if b {
		log.OutLevel = LevelDebug
	} else {
		log.OutLevel = LevelErr
	}
}

// Log logs the message with the level
func (log *Logger) Log(level Level, msg ...interface{}) {
	log.Logf(level, strings.Repeat("%s", len(msg)), msg...)
}

// Logf logs the message with the level
func (log *Logger) Logf(level Level, format string, args ...interface{}) {
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

	prefix := color + "[" + time + "] [" + lvl + "] "

	if len(log.Prefix) > 0 {
		prefix += log.Prefix + " "
	}

	log.Out.Printf(prefix+format+Reset, args...)
}

func (log *Logger) New() *Logger {
	return log.NewWithPrefix(log.Prefix)
}

func (log *Logger) NewWithPrefix(prefix string) *Logger {
	return &Logger{
		Out:             log.Out,
		OutLevel:        log.OutLevel,
		TimestampFormat: log.TimestampFormat,
		Prefix:          prefix,
	}
}

// LoggerWriter implements io.Writer maybe...
type LoggerWriter struct {
	Logger
	Level Level
}

// Write writes the logger p
func (log *LoggerWriter) Write(p []byte) (n int, err error) {
	log.Log(log.Level, string(p))
	return len(p), nil
}
