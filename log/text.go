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

	r := []rune(str)

	off := 0
	for off <= len(r)-1 {
		index := strings.Index(string(r[off:]), ESCAPE)
		s := r[off:]
		if index >= 0 {
			switch string(s[index : index+2]) {
			case Black, DarkBlue, DarkGreen, DarkAqua, DarkRed, Purple, Gold, Gray:
				fallthrough
			case DarkGray, Blue, Green, Aqua, Red, LightPurple, Yellow, White:
				fallthrough
			case Obfuscated, Bold, Strikethrough, Underline, Italic, Reset:
				if index > 0 {
					res = append(res, string(s[:index]))
				}

				res = append(res, string(s[index:index+2]))
			default:
				res = append(res, string(s[:index+2]))
			}

			off += index + 2
		} else {
			res = append(res, string(s[:]))
			break
		}
	}

	return res
}
