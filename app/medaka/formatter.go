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
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/beito123/medaka/log"
)

//TextFormatter format the text for Logrus
type TextFormatter struct {
}

//Format format the text
func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return append([]byte(ToANIS(entry.Message)), '\n'), nil
}

var FormatBold = ""
var FormatObfuscated = ""
var FormatItalic = ""
var FormatUnderline = ""
var FormatStrikethrough = ""

var FormatReset = ""

var ColorBlack = ""
var ColorDarkBlue = ""
var ColorDarkGreen = ""
var ColorDarkAqua = ""
var ColorDarkRed = ""
var ColorPurple = ""
var ColorGold = ""
var ColorGray = ""
var ColorDarkGray = ""
var ColorBlue = ""
var ColorGreen = ""
var ColorAqua = ""
var ColorRed = ""
var ColorLightPurple = ""
var ColorYellow = ""
var ColorWhite = ""

func InitFormat() {
	switch runtime.GOOS {
	case "linux", "freebsd": //linux
		initFormatCodes()
	default: //windows, ios
		initFallbackFormatCodes()
	}
}

func initFormatCodes() {
	FormatBold = cmd("tput", "bold")
	FormatObfuscated = cmd("tput", "smacs")
	FormatItalic = cmd("tput", "sitm")
	FormatUnderline = cmd("tput", "smul")
	FormatStrikethrough = "\x1b[9m"

	FormatReset = cmd("tput", "sgr0")

	colors, err := strconv.Atoi(cmd("tput", "colors"))
	if err != nil {
		colors = 256
	}

	if colors > 8 {
		if colors >= 256 {
			ColorBlack = cmd("tput", "setaf", "16")
			ColorDarkBlue = cmd("tput", "setaf", "19")
			ColorDarkGreen = cmd("tput", "setaf", "34")
			ColorDarkAqua = cmd("tput", "setaf", "37")
			ColorDarkRed = cmd("tput", "setaf", "124")
			ColorPurple = cmd("tput", "setaf", "127")
			ColorGold = cmd("tput", "setaf", "214")
			ColorGray = cmd("tput", "setaf", "145")
			ColorDarkGray = cmd("tput", "setaf", "59")
			ColorBlue = cmd("tput", "setaf", "63")
			ColorGreen = cmd("tput", "setaf", "83")
			ColorAqua = cmd("tput", "setaf", "87")
			ColorRed = cmd("tput", "setaf", "203")
			ColorLightPurple = cmd("tput", "setaf", "207")
			ColorYellow = cmd("tput", "setaf", "227")
			ColorWhite = cmd("tput", "setaf", "231")
		} else {
			ColorBlack = cmd("tput", "setaf", "0")
			ColorDarkBlue = cmd("tput", "setaf", "4")
			ColorDarkGreen = cmd("tput", "setaf", "2")
			ColorDarkAqua = cmd("tput", "setaf", "6")
			ColorDarkRed = cmd("tput", "setaf", "1")
			ColorPurple = cmd("tput", "setaf", "5")
			ColorGold = cmd("tput", "setaf", "3")
			ColorGray = cmd("tput", "setaf", "7")
			ColorDarkGray = cmd("tput", "setaf", "8")
			ColorBlue = cmd("tput", "setaf", "12")
			ColorGreen = cmd("tput", "setaf", "10")
			ColorAqua = cmd("tput", "setaf", "14")
			ColorRed = cmd("tput", "setaf", "9")
			ColorLightPurple = cmd("tput", "setaf", "13")
			ColorYellow = cmd("tput", "setaf", "11")
			ColorWhite = cmd("tput", "setaf", "15")
		}
	} else {
		ColorBlack = cmd("tput", "setaf", "0")
		ColorRed = cmd("tput", "setaf", "1")
		ColorGreen = cmd("tput", "setaf", "2")
		ColorYellow = cmd("tput", "setaf", "3")
		ColorBlue = cmd("tput", "setaf", "4")
		ColorLightPurple = cmd("tput", "setaf", "5")
		ColorAqua = cmd("tput", "setaf", "6")
		ColorGray = cmd("tput", "setaf", "7")

		ColorDarkGray = ColorBlack
		ColorDarkRed = ColorRed
		ColorDarkGreen = ColorGreen
		ColorGold = ColorYellow
		ColorDarkBlue = ColorBlue
		ColorPurple = ColorLightPurple
		ColorDarkAqua = ColorAqua
		ColorWhite = ColorGray
	}
}

func initFallbackFormatCodes() {
	FormatBold = "\x1b[1m"
	FormatObfuscated = ""
	FormatItalic = "\x1b[3m"
	FormatUnderline = "\x1b[4m"
	FormatStrikethrough = "\x1b[9m"

	FormatReset = "\x1b[m"

	ColorBlack = "\x1b[38;5;16m"
	ColorDarkBlue = "\x1b[38;5;19m"
	ColorDarkGreen = "\x1b[38;5;34m"
	ColorDarkAqua = "\x1b[38;5;37m"
	ColorDarkRed = "\x1b[38;5;124m"
	ColorPurple = "\x1b[38;5;127m"
	ColorGold = "\x1b[38;5;214m"
	ColorGray = "\x1b[38;5;145m"
	ColorDarkGray = "\x1b[38;5;59m"
	ColorBlue = "\x1b[38;5;63m"
	ColorGreen = "\x1b[38;5;83m"
	ColorAqua = "\x1b[38;5;87m"
	ColorRed = "\x1b[38;5;203m"
	ColorLightPurple = "\x1b[38;5;207m"
	ColorYellow = "\x1b[38;5;227m"
	ColorWhite = "\x1b[38;5;231m"
}

func cmd(n string, args ...string) string {
	out, err := exec.Command(n, args...).Output()
	if err != nil {
		return ""
	}

	return string(out)
}

//ToANIS convert format codes to terminal format codes
func ToANIS(str string) string {
	token := log.Tokenize(str)
	//fmt.Printf("%#v", token)

	var result strings.Builder
	for s := range token {
		switch token[s] {
		case log.Black:
			result.WriteString(ColorBlack)
		case log.DarkBlue:
			result.WriteString(ColorDarkBlue)
		case log.DarkGreen:
			result.WriteString(ColorDarkGreen)
		case log.DarkAqua:
			result.WriteString(ColorDarkAqua)
		case log.DarkRed:
			result.WriteString(ColorDarkRed)
		case log.Purple:
			result.WriteString(ColorPurple)
		case log.Gold:
			result.WriteString(ColorGold)
		case log.Gray:
			result.WriteString(ColorGray)
		case log.DarkGray:
			result.WriteString(ColorDarkGray)
		case log.Blue:
			result.WriteString(ColorBlue)
		case log.Green:
			result.WriteString(ColorGreen)
		case log.Aqua:
			result.WriteString(ColorAqua)
		case log.Red:
			result.WriteString(ColorRed)
		case log.LightPurple:
			result.WriteString(ColorLightPurple)
		case log.Yellow:
			result.WriteString(ColorYellow)
		case log.White:
			result.WriteString(ColorWhite)
		case log.Bold:
			result.WriteString(FormatBold)
		case log.Obfuscated:
			result.WriteString(FormatObfuscated)
		case log.Italic:
			result.WriteString(FormatItalic)
		case log.Underline:
			result.WriteString(FormatUnderline)
		case log.Strikethrough:
			result.WriteString(FormatStrikethrough)
		case log.Reset:
			result.WriteString(FormatReset)
		default:
			result.WriteString(token[s])
		}
	}

	return result.String()
}
