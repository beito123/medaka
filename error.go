package medaka

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

type ErrorType int

const (
	ErrorMedaka ErrorType = iota
	ErrorServer
	ErrorNetwork
	ErrorPermission
	ErrorCommand
	ErrorPlugin
)

//Error is application error on Medaka
type Error interface {
	error
	Type() ErrorType
}

func IsMedakaError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorMedaka
}

func IsServerError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorServer
}

func IsNetworkError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorNetwork
}

func IsPermissionError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorPermission
}

func IsCommandError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorCommand
}

func IsPluginError(err error) bool {
	e, ok := err.(Error)
	return ok && e.Type() == ErrorPlugin
}
