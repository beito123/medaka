package util

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

//OSType uses for specifying an OS.
type OSType int

const (
	//Windows is Windows OS
	Windows OSType = iota

	//Linux is Linux OS and Freebsd OS
	Linux

	//Mac is Darwin OS
	Mac

	//Unknown is unknown OS
	Unknown
)

//OS is a kind of Operating System
var OS = getOS()

func getOS() OSType {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "linux", "freebsd":
		return Linux
	case "darwin":
		return Mac
	}

	return Unknown
}

//KillProcess kills processes on OS with process id
func KillProcess(pid int) {
	switch runtime.GOOS {
	case "windows":
		exec.Command("taskkill", "/pid", strconv.Itoa(pid), "/F")
	default:
		exec.Command("kill", "-9", strconv.Itoa(pid), " > /dev/null 2>&1").Run()
	} //ummm... should support posix?

	//os.Exit(0)
}

//GetExtFromPath returns filename extension with file path
func GetExtFromPath(path string) string {
	i := strings.LastIndex(path, ".")
	if i >= 0 {
		return path[i+1:]
	}

	return path
}

//ExistFile
//Thanks: https://qiita.com/hnakamur/items/848097aad846d40ae84b
func ExistFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}

//ExistDir
//Thanks: https://qiita.com/hnakamur/items/848097aad846d40ae84b
func ExistDir(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
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
