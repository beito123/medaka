package app

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

	"github.com/beito123/medaka"
	"github.com/beito123/medaka/command"
	"github.com/beito123/medaka/lang"
	"github.com/beito123/medaka/util"
)

var instance *Server

//Instance returns generated server already
func Instance() *Server {
	return instance
}

type Server struct {
	Path string

	Logger        medaka.Logger
	CommandReader command.Reader
	CommandSender command.Sender
	Lang          *lang.Lang

	settings *util.Config

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
	ser.Logger.SetLogDebug(ser.GetSettingsBool("settings.debug"))

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

	supportedLang := []string{
		"eng",
		"jpn",
	}

	language := ser.GetSettingString("settings.language")
	supported := false
	for i := range supportedLang {
		if supportedLang[i] == language {
			supported = true
			break
		}
	}

	if !supported {
		language = "eng"
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

	ser.Logger.Info(ser.Lang.Message("medaka.lang.loaded"))

	//start message
	ser.Logger.Info(ser.Lang.Message("medaka.server.start", medaka.SupportMCBEVersion))

	//ready async task//decide pool size//workers

	//Ready Network
	//Set Threshold for batching packets
	//Set CompressionLevel
	//Set CompressionAsync

	//Ready Server scheduler

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

	ser.checkConsole()

	return true
}

func (ser *Server) checkConsole() {
	//TODO: writes simple command system
	text := ser.CommandReader.Line()
	ser.Logger.Notice("Command:" + fmt.Sprintf("%#v", text))
	if strings.Index(text, "stop") >= 0 {
		ser.Logger.Info("Stopping the server...")

		ser.Shutdown()
	}

	if strings.Index(text, "version") >= 0 {
		ser.Logger.Notice("version:" + medaka.Version + ", revision:" + medaka.Revision)
	}
}

func (ser *Server) sendCommand(sender command.Sender, cmd string) {
	//
}

func (ser *Server) Shutdown() {
	ser.running = false
}

func (ser *Server) forceShutdown() {
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

	//Save and close...

	//all kick
	//network
	//plugin
	//world
}

func (ser *Server) GetSettings() *util.Config {
	return ser.settings
}

func (ser *Server) GetSettingsBool(key string) bool {
	return ser.settings.GetBool(key)
}

func (ser *Server) GetSettingString(key string) string {
	return ser.settings.GetString(key)
}

func (ser *Server) GetSettingInt(key string) int {
	return ser.settings.GetInt(key)
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
	ser.Logger.Trace(medaka.Dump(1, medaka.TraceLimit))

	ser.forceShutdown()
}
