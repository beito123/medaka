package main

import "github.com/beito123/medaka"

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

func NewSender(logger medaka.Logger) *ConsoleSender {
	return &ConsoleSender{
		logger: logger,
	}
}

//ConsoleSender implements command sender
type ConsoleSender struct {
	logger medaka.Logger
}

func (sender *ConsoleSender) Name() string {
	return "Console"
}

func (sender *ConsoleSender) SendMessage(msg string) {
	sender.logger.Info(msg)
}
