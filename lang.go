package medaka

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"io"
	"strings"
)

var SupportedLang = []string{
	"eng",
	"jpn",
}

const (
	DefaultLang = "eng"
)

func IsSupportedLang(language string) bool {
	lang := strings.ToLower(language)

	supported := false
	for i := range SupportedLang {
		if SupportedLang[i] == lang {
			supported = true
			break
		}
	}

	return supported
}

func LoadLangFile(language string) io.Reader {
	langFile, err := OpenResource("/static/lang/" + strings.ToLower(language) + ".ini")
	if err != nil {
		return nil
	}

	return langFile
}
