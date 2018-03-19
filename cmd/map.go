package cmd

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

	"github.com/beito123/medaka"
)

type SimpleMap struct {
	CommandMap map[string]medaka.Command
	AliasMap   map[string]medaka.Command
}

func (m *SimpleMap) Add(cmd medaka.Command) error {
	if m.Exist(cmd.Name()) {
		return AlreadyRegisteredError{}
	}

	aliases := cmd.Aliases()

	for i := range aliases {
		if m.existAlias(aliases[i]) {
			return AlreadyExistAliasError{}
		}
	}

	m.CommandMap[strings.ToLower(cmd.Name())] = cmd

	for i := range aliases {
		m.AliasMap[strings.ToLower(aliases[i])] = cmd
	}

	return nil
}

func (m *SimpleMap) Remove(cmd string) error {
	if !m.Exist(cmd) {
		return NotExistError{}
	}

	delete(m.CommandMap, strings.ToLower(cmd))

	return nil
}

func (m *SimpleMap) Exist(cmd string) bool {
	_, ok := m.CommandMap[strings.ToLower(cmd)]

	return ok
}

func (m *SimpleMap) existAlias(name string) bool {
	_, ok := m.AliasMap[strings.ToLower(name)]

	return ok
}

func (m *SimpleMap) Command(name string) medaka.Command {
	name = strings.ToLower(name)

	if m.existAlias(name) {
		return m.AliasMap[name]
	}

	if m.Exist(name) {
		return m.CommandMap[name]
	}

	return nil
}

func (m *SimpleMap) Aliases() map[string]medaka.Command {
	return m.AliasMap
}

func (m *SimpleMap) List() []medaka.Command {
	list := make([]medaka.Command, 0, len(m.CommandMap))
	for _, value := range m.CommandMap {
		list = append(list, value)
	}

	return list
}

func (m *SimpleMap) Map() map[string]medaka.Command {
	return m.CommandMap
}
