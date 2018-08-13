package log

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"strings"
)

const (
	ESCAPE = "\xc2\xa7" //§
	EOL    = "\n"
)

const (
	Black       = ESCAPE + "0"
	DarkBlue    = ESCAPE + "1"
	DarkGreen   = ESCAPE + "2"
	DarkAqua    = ESCAPE + "3"
	DarkRed     = ESCAPE + "4"
	Purple      = ESCAPE + "5"
	Gold        = ESCAPE + "6"
	Gray        = ESCAPE + "7"
	DarkGray    = ESCAPE + "8"
	Blue        = ESCAPE + "9"
	Green       = ESCAPE + "a"
	Aqua        = ESCAPE + "b"
	Red         = ESCAPE + "c"
	LightPurple = ESCAPE + "d"
	Yellow      = ESCAPE + "e"
	White       = ESCAPE + "f"
)

const (
	Obfuscated    = ESCAPE + "k"
	Bold          = ESCAPE + "l"
	Strikethrough = ESCAPE + "m"
	Underline     = ESCAPE + "n"
	Italic        = ESCAPE + "o"
	Reset         = ESCAPE + "r"
)

func Tokenize(str string) []string {
	var res []string

	lenEscape := len(ESCAPE)

	off := 0

	for off < len(str) {
		s := str[off:]
		index := strings.Index(s, ESCAPE)
		if index < 0 {
			res = append(res, s)
			break
		}

		codeIndex := index + lenEscape
		if len(s[codeIndex:]) < 1 { //if left 0
			res = append(res, s[:codeIndex])
			break
		}

		code := string([]rune(s[index:])[1:2])

		last := index + lenEscape + len(code)

		if index > 0 {
			res = append(res, s[:index])
		}

		res = append(res, s[index:last])

		off += last
	}

	return res
}
