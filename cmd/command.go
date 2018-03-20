package cmd

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import "github.com/beito123/medaka/lang"

//Reader is basic command reader interface
type Reader interface {
	Line() string
}

//Sender implements to sender of command
type Sender interface {
	Name() string
	SendMessage(msg string)
	SendMessageWithText(text *lang.Text) //TODO: Improve
	Permission() int                     //TODO: rewrite
	HasPermission(per string) bool
}

//Command is basic command interface
type Command interface {
	Name() string
	Description() string
	Usage() string
	Permission() string
	Aliases() []string
	Execute(sender Sender, args []string) error
}

//Map is basic command map interface
type Map interface {
	Add(cmd Command) error
	Remove(cmd string) error
	Exist(cmd string) bool
	Command(cmd string) Command
	Aliases() map[string]Command
	List() []Command
	Map() map[string]Command
}
