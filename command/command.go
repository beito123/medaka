package command

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

//Reader is basic command reader interface
type Reader interface {
	Line() string
}

//Sender implements to sender of command
type Sender interface {
	Name() string
	SendMessage(msg string)
}
