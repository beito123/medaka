package main

/*
	Medaka
	Copyright (c) 2018 beito
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"bufio"
	"os"
	"strings"
)

func NewReader() *ConsoleReader {
	reader := &ConsoleReader{}

	reader.start()

	return reader
}

type ConsoleReader struct {
	reader *bufio.Reader
}

func (console *ConsoleReader) start() {
	console.reader = bufio.NewReader(os.Stdin)
}

func (console *ConsoleReader) Line() string {
	text, _ := console.reader.ReadString('\n')

	return convertNewline(text, "")
}

//convertNewline replaces line code to nlcode in str
//Thanks: https://qiita.com/spiegel-im-spiegel/items/f1cc014ecb233afaa8af
func convertNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}
