package mcbe

import (
	"github.com/beito123/medaka"
	"github.com/beito123/medaka/lang"
	"github.com/beito123/medaka/log"
)

/*
	Medaka

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

type VanillaCommand interface {
	medaka.Command
	Run(sender medaka.CommandSender, args []string) bool
}

//DefaultCommand is basic command for vanilla commands
type DefaultCommand struct {
	medaka.Command
}

func (cmd *DefaultCommand) Aliases() []string {
	return []string{}
}

func (base *DefaultCommand) Execute(cmd VanillaCommand, sender medaka.CommandSender, args []string) error {
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

func (cmd *VersionCommand) Execute(sender medaka.CommandSender, args []string) error {
	return cmd.DefaultCommand.Execute(cmd, sender, args)
}

func (cmd *VersionCommand) Run(sender medaka.CommandSender, args []string) bool {
	sender.SendMessageWithText(lang.NewText("command.version.execute", medaka.Version, medaka.Revision, medaka.SupportMCBEVersion))

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

func (cmd *StopCommand) Execute(sender medaka.CommandSender, args []string) error {
	return cmd.DefaultCommand.Execute(cmd, sender, args)
}

func (cmd *StopCommand) Run(sender medaka.CommandSender, args []string) bool {
	sender.SendMessageWithText(lang.NewText("command.stop.execute"))

	ser := Instance()

	ser.Shutdown()

	return true
}
