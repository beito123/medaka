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
	"net/http"
	"os"

	"github.com/beito123/medaka/data"
)

func OpenResource(path string) (http.File, error) {
	return data.Assets.Open(path)
}

func ExistResource(path string) bool {
	_, err := OpenResource(path)
	if err != nil {
		return false
	}

	return true
}

func CopyResource(path string, to string) error {
	src, err := OpenResource(path)
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(to)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
