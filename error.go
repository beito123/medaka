package medaka

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

//NewError generates a MedakaError
func NewError(msg string) Error {
	return Error{
		Message: msg,
	}
}

//Error is application error on Medaka
type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}
