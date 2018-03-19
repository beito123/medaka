package lang

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

func NewText(key string, args ...string) *Text {
	return &Text{
		Key:  key,
		Args: args,
	}
}

func NewTextWithPrefix(key string, prefix string, args ...string) *Text {
	return &Text{
		Key:    key,
		Args:   args,
		Prefix: prefix,
	}
}

//Text is text container for Lang
type Text struct {
	Key    string
	Args   []string
	Prefix string
}
