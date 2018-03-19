package cmd

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import "github.com/beito123/medaka"

func IsAlreadyRegisteredError(err error) bool {
	_, ok := err.(AlreadyRegisteredError)
	return ok
}

func IsAlreadyExistAliasError(err error) bool {
	_, ok := err.(AlreadyExistAliasError)
	return ok
}

func IsNotExistError(err error) bool {
	_, ok := err.(NotExistError)
	return ok
}

type CommandError struct {
}

func (e *CommandError) Type() medaka.ErrorType {
	return medaka.ErrorCommand
}

type AlreadyRegisteredError struct {
	CommandError
}

func (e AlreadyRegisteredError) Error() string {
	return "it has already registered the same name."
}

type AlreadyExistAliasError struct {
	CommandError
}

func (e AlreadyExistAliasError) Error() string {
	return "The alias is existing already."
}

type NotExistError struct {
	CommandError
}

func (e NotExistError) Type() medaka.ErrorType {
	return medaka.ErrorCommand
}

func (e NotExistError) Error() string {
	return "The command doesn't exist."
}
