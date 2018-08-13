package main

import (
	"github.com/beito123/medaka"
	"github.com/beito123/medaka/lang"
	"github.com/beito123/medaka/log"
)

/*
	Medaka
	Copyright (c) 2018 beito
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

func NewSender(logger *log.Logger) *ConsoleSender {
	return &ConsoleSender{
		logger: logger,
	}
}

//ConsoleSender implements command sender
type ConsoleSender struct {
	logger *log.Logger
}

func (sender *ConsoleSender) Name() string {
	return "Console"
}

func (sender *ConsoleSender) server() *medaka.Server {
	return medaka.Instance()
}

func (sender *ConsoleSender) SendMessage(msg string) {
	sender.logger.Info(msg)
}

func (sender *ConsoleSender) SendMessageWithText(text *lang.Text) {
	sender.SendMessage(sender.server().Translate(text))
}

func (sender *ConsoleSender) Permission() int {
	return 0
}

func (sender *ConsoleSender) HasPermission(per string) bool {
	return true //TODO: check nodes via Server
}
