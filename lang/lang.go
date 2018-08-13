package lang

/*
	Medaka

	Copyright (c) 2018 beito
	
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/beito123/medaka/util"
)

type Lang struct {
	Prefix string

	config  *util.Config
	correct bool
}

func (lang *Lang) Load(path string) error {
	if lang.correct {
		return errors.New("Already loaded")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return lang.LoadReader(file)
}

func (lang *Lang) LoadReader(reader io.Reader) error {
	if lang.correct {
		return errors.New("Already loaded")
	}

	lang.config = util.NewConfig()
	err := lang.config.LoadReader(reader, util.Properties)
	if err != nil {
		return err
	}

	lang.correct = true

	return nil
}

func (lang *Lang) Translate(text *Text) string {
	msg := lang.config.GetString(text.Key, "")
	if len(msg) <= 0 {
		return "%" + text.Key
	}

	for i, v := range text.Args {
		msg = strings.Replace(msg, "{%"+strconv.Itoa(i)+"}", v, -1)
	}

	msg = text.Prefix + msg

	if len(lang.Prefix) > 0 {
		msg = "[" + lang.Prefix + "] " + msg
	}

	return msg
}

func (lang *Lang) TranslateWithString(key string, args ...string) string {
	return lang.Translate(&Text{
		Key:  key,
		Args: args,
	})
}
