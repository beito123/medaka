package main

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/beito123/medaka"
	"github.com/beito123/medaka/log"
	"github.com/beito123/medaka/util"

	colorable "github.com/mattn/go-colorable"
)

func main() {
	InitFormat()

	out := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	if util.OS == util.Windows {
		out.Out = colorable.NewColorableStdout() // Thanks you, mattn
	}

	logger := log.NewLogger(out)

	logger.Info("Start server...")

	path, err := os.Executable()
	if err != nil {
		logger.Fatal("Couldn't gets a path of the executable file")
		return
	}

	//server := medaka.NewServer(logger, filepath.Dir(path))

	server := &medaka.Server{
		Path:          filepath.Dir(path),
		Logger:        logger,
		CommandReader: NewReader(),
		CommandSender: NewSender(logger),
	}

	server.Start()

	logger.Debug("The End...")
}
