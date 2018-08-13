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
	"errors"
	"strings"

	"github.com/beito123/medaka/lang"
	"github.com/beito123/medaka/log"
)

var (
	errNotExist     = errors.New("not exist")
	errAlreadyExist = errors.New("already exist command")
	errExistAlias   = errors.New("already exist alias")
)

// CommandReader is basic command reader interface
type CommandReader interface {
	Line() string
}

// CommandSender implements to sender of command
type CommandSender interface {
	Name() string
	SendMessage(msg string)
	SendMessageWithText(text *lang.Text) //TODO: Improve
	Permission() int                     //TODO: rewrite
	HasPermission(per string) bool
}

// Command is basic command interface
type Command interface {
	Name() string
	Description() string
	Usage() string
	Permission() string
	Aliases() []string
	Execute(sender CommandSender, args []string) error
}

// CommandMap is basic command map interface
type CommandMap interface {
	Add(cmd Command) error
	Remove(cmd string) error
	Exist(cmd string) bool
	Command(cmd string) Command
	Aliases() map[string]Command
	List() []Command
	Map() map[string]Command
}

func NewSimpleCommandMap() *SimpleCommandMap {
	return &SimpleCommandMap{
		CommandMap: make(map[string]Command),
		AliasMap:   make(map[string]Command),
	}
}

type SimpleCommandMap struct {
	CommandMap map[string]Command
	AliasMap   map[string]Command
}

func (m *SimpleCommandMap) Add(cmd Command) error {
	if m.Exist(cmd.Name()) {
		return errAlreadyExist
	}

	aliases := cmd.Aliases()

	for i := range aliases {
		if m.existAlias(aliases[i]) {
			return errExistAlias
		}
	}

	m.CommandMap[strings.ToLower(cmd.Name())] = cmd

	for i := range aliases {
		m.AliasMap[strings.ToLower(aliases[i])] = cmd
	}

	return nil
}

func (m *SimpleCommandMap) Remove(cmd string) error {
	if !m.Exist(cmd) {
		return errNotExist
	}

	delete(m.CommandMap, strings.ToLower(cmd))

	return nil
}

func (m *SimpleCommandMap) Exist(cmd string) bool {
	_, ok := m.CommandMap[strings.ToLower(cmd)]

	return ok
}

func (m *SimpleCommandMap) existAlias(name string) bool {
	_, ok := m.AliasMap[strings.ToLower(name)]

	return ok
}

func (m *SimpleCommandMap) Command(name string) Command {
	name = strings.ToLower(name)

	if m.existAlias(name) {
		return m.AliasMap[name]
	}

	if m.Exist(name) {
		return m.CommandMap[name]
	}

	return nil
}

func (m *SimpleCommandMap) Aliases() map[string]Command {
	return m.AliasMap
}

func (m *SimpleCommandMap) List() []Command {
	list := make([]Command, 0, len(m.CommandMap))
	for _, value := range m.CommandMap {
		list = append(list, value)
	}

	return list
}

func (m *SimpleCommandMap) Map() map[string]Command {
	return m.CommandMap
}

type VanillaCommand interface {
	Command
	Run(sender CommandSender, args []string) bool
}

//DefaultCommand is a basic command interface for vanilla commands
type DefaultCommand struct {
	Command
}

func (cmd *DefaultCommand) Aliases() []string {
	return []string{}
}

func (base *DefaultCommand) Execute(cmd VanillaCommand, sender CommandSender, args []string) error {
	if !sender.HasPermission(cmd.Permission()) {
		sender.SendMessageWithText(lang.NewTextWithPrefix("command.noPermission", log.Red))
		return nil
	}

	result := cmd.Run(sender, args)
	if !result {
		sender.SendMessageWithText(lang.NewText("command.usage", cmd.Description()))
		return nil
	}

	return nil
}

//VersionCommand versions the server
type VersionCommand struct {
	DefaultCommand
}

func (cmd *VersionCommand) Name() string {
	return "version"
}

func (cmd *VersionCommand) Description() string {
	return "%command.version.description"
}

func (cmd *VersionCommand) Usage() string {
	return "%command.version.usage"
}

func (cmd *VersionCommand) Permission() string {
	return "minecraft.command.version"
}

func (cmd *VersionCommand) Execute(sender CommandSender, args []string) error {
	return cmd.DefaultCommand.Execute(cmd, sender, args)
}

func (cmd *VersionCommand) Run(sender CommandSender, args []string) bool {
	sender.SendMessageWithText(lang.NewText("command.version.execute", Version, Revision, SupportMCBEVersion))

	return true
}

//StopCommand stops the server
type StopCommand struct {
	DefaultCommand
}

func (cmd *StopCommand) Name() string {
	return "stop"
}

func (cmd *StopCommand) Description() string {
	return "%command.stop.description"
}

func (cmd *StopCommand) Usage() string {
	return "%command.stop.usage"
}

func (cmd *StopCommand) Permission() string {
	return "minecraft.command.stop"
}

func (cmd *StopCommand) Execute(sender CommandSender, args []string) error {
	return cmd.DefaultCommand.Execute(cmd, sender, args)
}

func (cmd *StopCommand) Run(sender CommandSender, args []string) bool {
	sender.SendMessageWithText(lang.NewText("command.stop.execute"))

	ser := Instance()

	ser.Shutdown()

	return true
}
