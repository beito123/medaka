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
	"errors"
	"io"
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

//ConfigFormat uses for specifying a format.
//a wrapper of viper
type ConfigFormat int

const (
	//Properties is properties format for Config
	Properties ConfigFormat = iota

	//YAML is YAML format for Config
	YAML

	//TOML is TOML format for Config
	TOML

	//JSON is JSON format for Config
	JSON

	//Note: Enum won't support because it is outdated.
)

//NewConfig returns new Config
func NewConfig() *Config {
	return &Config{}
}

//Config is basic config struct in Medaka
type Config struct {
	path    string
	format  ConfigFormat
	Content *viper.Viper

	hasChanged bool
	readOnly   bool
	correct    bool
}

//Path returns file path
func (config *Config) Path() string {
	return config.path
}

//Format returns file format
func (config *Config) Format() ConfigFormat {
	return config.format
}

func (config *Config) HasChanged() bool {
	return config.hasChanged
}

func (config *Config) ReadOnly() bool {
	return config.readOnly
}

//Load loads the file with path and type
func (config *Config) Load(path string, format ConfigFormat, def map[string]interface{}) error {
	if config.correct {
		return errors.New("util.Config.Load: Already loaded.")
	}

	config.path = path
	config.readOnly = false
	config.hasChanged = false

	config.Content = viper.New()

	if def == nil {
		def = make(map[string]interface{})
	}

	config.SetDefaults(def)

	if !config.setFormat(format) {
		return errors.New("util.Config.Load: The config type doesn't support.")
	}

	if !ExistFile(config.path) {
		config.Save()
	}

	file, err := os.Open(config.path)
	if err != nil {
		return err
	}

	defer file.Close()

	err = config.Content.ReadConfig(file)
	if err != nil {
		return err
	}

	config.correct = true

	return nil
}

//LoadReader loads the file with reader and type
func (config *Config) LoadReader(reader io.Reader, format ConfigFormat) error {
	if config.correct {
		return errors.New("util.Config.LoadReader: Already loaded.")
	}

	config.path = ""
	config.readOnly = true
	config.hasChanged = false

	config.Content = viper.New()

	if !config.setFormat(format) {
		return errors.New("util.Config.LoadReader: The config type doesn't support.")
	}

	err := config.Content.ReadConfig(reader)
	if err != nil {
		return err
	}

	config.correct = true

	return nil
}

func (config *Config) setFormat(format ConfigFormat) bool {
	switch format {
	case Properties:
		config.Content.SetConfigType("Properties")
	case YAML:
		config.Content.SetConfigType("YAML")
	case TOML:
		config.Content.SetConfigType("TOML")
	case JSON:
		config.Content.SetConfigType("JSON")
	default:
		return false
	}

	return true
}

//Reset resets the config
func (config *Config) Reset() {
	config = &Config{}
}

//Save saves as file
func (config *Config) Save() error {
	if config.readOnly {
		return errors.New("util.Config.Save: The config is enable readonly!")
	}

	config.hasChanged = false

	return config.Content.WriteConfigAs(config.path)
}

//SetDefault set default value to Config
func (config *Config) SetDefault(key string, value interface{}) {
	if config.readOnly {
		return
	}

	config.hasChanged = true

	config.Content.SetDefault(key, value)
}

//SetDefaults set default sets to Config
func (config *Config) SetDefaults(def map[string]interface{}) {
	if config.readOnly {
		return
	}

	for k, v := range def {
		config.SetDefault(k, v)
	}
}

//Get gets a value with key
func (config *Config) Get(key string, def interface{}) interface{} {
	val := config.Content.Get(key)
	if val == nil {
		return def
	}

	return val
}

//Set sets a value with key
func (config *Config) Set(key string, value interface{}) {
	if config.readOnly {
		return
	}

	config.hasChanged = true

	config.Content.Set(key, value)
}

//Exist returns existing data with key
func (config *Config) Exist(key string) bool {
	return config.Content.IsSet(key)
}

//GetString gets a value as string with key
func (config *Config) GetString(key string, def string) string {
	return cast.ToString(config.Get(key, def))
}

//GetBool gets a value as boolean with key
func (config *Config) GetBool(key string, def bool) bool {
	return cast.ToBool(config.Get(key, def))
}

//GetInt gets a value as int with key
func (config *Config) GetInt(key string, def int) int {
	return cast.ToInt(config.Get(key, def))
}

//GetUInt gets a value as uint with key
func (config *Config) GetUInt(key string, def uint) uint {
	return cast.ToUint(config.Get(key, def))
}

//GetInt64 gets a value as int64 with key
func (config *Config) GetInt64(key string, def int64) int64 {
	return cast.ToInt64(config.Get(key, def))
}

//GetUInt64 gets a value as uint64 with key
func (config *Config) GetUInt64(key string, def uint64) uint64 {
	return cast.ToUint64(config.Get(key, def))
}

//GetFloat32 gets a value as float32 with key
func (config *Config) GetFloat32(key string, def float32) float32 {
	return cast.ToFloat32(config.Get(key, def))
}

//GetFloat64 gets a value as float64 with key
func (config *Config) GetFloat64(key string, def float64) float64 {
	return cast.ToFloat64(config.Get(key, def))
}

//GetStringSlice gets a value as []string with key
func (config *Config) GetStringSlice(key string, def []string) []string {
	return cast.ToStringSlice(config.Get(key, def))
}

//GetStringMap gets a value as map[string]interface{} with key
func (config *Config) GetStringMap(key string, def map[string]interface{}) map[string]interface{} {
	return cast.ToStringMap(config.Get(key, def))
}

//GetStringMapString gets a value as map[string]string with key
func (config *Config) GetStringMapString(key string, def map[string]string) map[string]string {
	return cast.ToStringMapString(config.Get(key, def))
}

//GetStringMapStringSlice gets a value as map[string][]string with key
func (config *Config) GetStringMapStringSlice(key string, def map[string][]string) map[string][]string {
	return cast.ToStringMapStringSlice(config.Get(key, def))
}

//SetString sets a value as string with key
func (config *Config) SetString(key string, val string) {
	config.Set(key, val)
}

//SetBool sets a value as boolean with key
func (config *Config) SetBool(key string, val bool) {
	config.Set(key, val)
}

//SetInt sets a value as int with key
func (config *Config) SetInt(key string, val int) {
	config.Set(key, val)
}

//SetUInt sets a value as uint with key
func (config *Config) SetUInt(key string, val uint) {
	config.Set(key, val)
}

//SetInt64 sets a value as int64 with key
func (config *Config) SetInt64(key string, val int64) {
	config.Set(key, val)
}

//SetUInt64 sets a value as uint64 with key
func (config *Config) SetUInt64(key string, val uint64) {
	config.Set(key, val)
}

//SetFloat32 sets a value as float32 with key
func (config *Config) SetFloat32(key string, val float32) {
	config.Set(key, val)
}

//SetFloat64 sets a value as float64 with key
func (config *Config) SetFloat64(key string, val float64) {
	config.Set(key, val)
}

//SetStringSlice sets a value as []string with key
func (config *Config) SetStringSlice(key string, val []string) {
	config.Set(key, val)
}

//SetStringMap sets a value as map[string]interface{} with key
func (config *Config) SetStringMap(key string, val map[string]interface{}) {
	config.Set(key, val)
}

//SetStringMapString sets a value as map[string]string with key
func (config *Config) SetStringMapString(key string, val map[string]string) {
	config.Set(key, val)
}

//SetStringMapStringSlice sets a value as map[string][]string with key
func (config *Config) SetStringMapStringSlice(key string, val map[string][]string) {
	config.Set(key, val)
}
