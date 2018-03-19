package permission

import "github.com/beito123/medaka"

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

func IsUnauthorisedError(err error) bool {
	_, ok := err.(UnauthorisedError)
	return ok
}

//NewError generates a Error
func NewError(msg string) Error {
	return Error{
		Message: msg,
	}
}

//Error is permission error
type Error struct {
	Message   string
	errorType medaka.ErrorType
}

func (e Error) Type() medaka.ErrorType {
	return e.errorType
}

func (e Error) Error() string {
	return e.Message
}

type UnauthorisedError struct {
	Message   string
	errorType medaka.ErrorType
}

func (e UnauthorisedError) Type() medaka.ErrorType {
	return e.errorType
}

func (e UnauthorisedError) Error() string {
	return e.Message
}
