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
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/beito123/medaka/lang"
	"github.com/beito123/medaka/log"
	"github.com/beito123/medaka/util"
)

var instance *Server

func Instance() *Server {
	return instance
}

type Server struct {
	Path string

	Logger        *log.Logger
	Lang          *lang.Lang
	CommandReader CommandReader
	CommandSender CommandSender

	settings *util.Config

	commandMap CommandMap

	running bool
	stopped bool
}

func (ser *Server) Start() {

	defer func() {
		err := recover()
		if err != nil {
			ser.handlePanic(err)
		}
	}()

	if ser.running {
		panic("The server is already running!")
	}

	instance = ser
	ser.running = true
	ser.stopped = false

	log.DefaultTraceLimit = 30 // For debug

	//generates dir//worlds, players, plugins
	if !util.ExistDir(ser.Path + "/worlds") {
		err := os.MkdirAll(ser.Path+"/worlds", 0777)
		if err != nil {
			ser.Logger.Fatal("Couldn't make worlds directory.")
			panic(err)
		}
	}

	if !util.ExistDir(ser.Path + "/players") {
		err := os.MkdirAll(ser.Path+"/players", 0777)
		if err != nil {
			ser.Logger.Fatal("Couldn't make players directory.")
			panic(err)
		}
	}

	//load medaka.yml
	ser.Logger.Info("Loading medaka.yml...")
	if !util.ExistFile(ser.Path + "/medaka.yml") {
		CopyResource("/static/medaka.yml", ser.Path+"/medaka.yml")
	}

	ser.settings = util.NewConfig()
	ser.settings.Load(ser.Path+"/medaka.yml", util.YAML, nil)

	//set log debug for logger from medaka.yml
	ser.Logger.SetLogDebug(ser.SettingBool("settings.debug", false))

	ser.Logger.Info("Loading server properties...")

	config := util.NewConfig()
	err := config.Load(ser.Path+"/server.properties", util.Properties, map[string]interface{}{
		"motd":                         "A Minecraft Server",
		"sub-motd":                     nil,
		"server-port":                  19132,
		"server-ip":                    nil,
		"view-distance":                10,
		"white-list":                   false,
		"achievements":                 true,
		"announce-player-achievements": true,
		"spawn-protection":             16,
		"max-players":                  20,
		"allow-flight":                 false,
		"spawn-animals":                true,
		"spawn-mobs":                   true,
		"gamemode":                     0,
		"force-gamemode":               false,
		"hardcore":                     false,
		"pvp":                          true,
		"difficulty":                   1,
		"level-name":                   "world",
		"level-seed":                   nil,
		"level-type":                   "DEFAULT",
		"enable-query":                 true,
		"auto-save":                    true,
		"force-resources":              false,
		"xbox-auth":                    false,
	})

	if err != nil {
		ser.Logger.Fatal("Couldn't load server.properties.")
		panic(err)
	}

	//Lang

	language := ser.SettingString("settings.language", DefaultLang)

	if !IsSupportedLang(language) {
		language = DefaultLang
	}

	langReader := LoadLangFile(language)
	if langReader == nil {
		panic("Couldn't get a language file.")
	}

	ser.Lang = &lang.Lang{}
	err = ser.Lang.LoadReader(langReader)
	if err != nil {
		ser.Logger.Fatal("Couldn't load language.")
		panic(err)
	}

	ser.Logger.Info(ser.TranslateWithString("medaka.lang.loaded"))

	//start message
	ser.Logger.Info(ser.TranslateWithString("medaka.server.start", SupportMCBEVersion, CodeName))

	//Scheduler

	//Ready Network
	//Set Threshold for batching packets
	//Set CompressionLevel
	//Set CompressionAsync

	//ops
	//whitelist
	//ban-player
	//ban-ips

	//max-player
	//autoSave

	//onlineMode
	//hardcore

	//network
	//set motd

	//log server info
	//license

	//timings

	//consoleSender
	//simpleCommandMap
	ser.initCommands()

	//init
	//Entity
	//Tile
	//BlockFactor
	//Enchantment
	//ItemFactory
	//ItemCreative
	//Biome

	//CraftingManager

	//ResourcePackManager

	//plugins manager

	//

	//query

	//LoadPlugins

	//enable statup plugins

	//register interface//raklibInterface

	//init LevelProvider

	//register Default Generators

	//worlds

	//default level

	if config.HasChanged() {
		config.Save()
	}

	//if default isn't level object, shutdown.

	//enable plugins

	ser.tickProcesser()
}

func (ser *Server) tickProcesser() {
	ticker := time.NewTicker(time.Duration(20) * time.Millisecond) //20 tick
	defer ticker.Stop()

	for tick := range ticker.C {
		if !ser.tick(tick) {
			break
		}
	}

	ser.forceShutdown()
}

func (ser *Server) tick(tick time.Time) bool {
	if !ser.running {
		return false
	}

	//update network

	//update plugins

	//scheduler

	//Update players and levels

	//Update Title
	//Update network name(player count etc...)
	//Update network statistics

	//Update Query

	//Autosave

	//Clear level cache

	//TPS
	//USE

	ser.checkConsole()

	return true
}

func (ser *Server) checkConsole() {
	//TODO: writes simple command system
	text := ser.CommandReader.Line()

	if len(text) == 0 {
		return
	}

	exp := strings.Split(text, " ")

	cmd := ser.commandMap.Command(exp[0])
	if cmd == nil {
		ser.Logger.Info(ser.TranslateWithString("command.unknown"))
		return
	}

	var args []string
	if len(exp) > 1 {
		args = exp[1:]
	}

	ser.SendCommand(ser.CommandSender, cmd, args)
}

func (ser *Server) SendCommand(sender CommandSender, command Command, args []string) {
	//event

	err := command.Execute(sender, args)
	if err != nil {
		ser.Logger.Fatal("Happened the error while executeing the command.")
		ser.Logger.Err(err, nil)
	}
}

func (ser *Server) Shutdown() {
	ser.running = false
}

func (ser *Server) forceShutdown() { //TODO: implements shutdown thread...
	if ser.stopped {
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			ser.Logger.Fatal("Crashed while crashing, kill this process.")

			util.KillProcess(os.Getpid()) //The End :P
		}
	}()

	ser.stopped = true

	ser.Shutdown()

	ser.Logger.Info(ser.TranslateWithString("medaka.server.stop"))

	//Save and close...

	//all kick
	//players
	//network
	//plugin
	//world
}

func (ser *Server) Settings() *util.Config {
	return ser.settings
}

func (ser *Server) SettingBool(key string, def bool) bool {
	return ser.settings.GetBool(key, def)
}

func (ser *Server) SettingString(key string, def string) string {
	return ser.settings.GetString(key, def)
}

func (ser *Server) SettingInt(key string, def int) int {
	return ser.settings.GetInt(key, def)
}

func (ser *Server) Translate(text *lang.Text) string {
	return ser.Lang.Translate(text)
}

func (ser *Server) TranslateWithString(key string, args ...string) string {
	return ser.Translate(&lang.Text{
		Key:  key,
		Args: args,
	})
}

func (ser *Server) initCommands() {
	ser.commandMap = NewSimpleCommandMap()

	ser.commandMap.Add(&VersionCommand{})
	ser.commandMap.Add(&StopCommand{})
}

func (ser *Server) handlePanic(err interface{}) {
	text := ""
	switch e := err.(type) {
	case string:
		text = e
	case error: //ummm...
		text = e.Error()
	default:
		text = fmt.Sprintf("%#v", e)
	}

	//TODO: crashdump

	ser.Logger.Fatal("Panic: " + text)
	ser.Logger.Trace(log.Dump(1, log.DefaultTraceLimit))

	ser.forceShutdown()
}