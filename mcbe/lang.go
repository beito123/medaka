package mcbe

import (
	"io"
)

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

func LoadLangFile(language string) io.Reader {
	langFile, err := OpenResource("/static/lang/" + language + ".ini")
	if err != nil {
		return nil
	}

	return langFile
}
