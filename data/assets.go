package data

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets25a0b1481b5ef1da0a69fc5fc4cc4aa70e567123 = "#Advanced config file for Medaka\r\n\r\nsettings:\r\n  language: \"eng\"\r\n  force-language: false\r\n  shutdown-message: \"Server closed\"\r\n  query-plugins: true\r\n  async-workers: auto\r\n  debug: true"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"static"}, "/static": []string{"medaka.yml"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521258460, 1521258460213802700),
		Data:     nil,
	}, "/static": &assets.File{
		Path:     "/static",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521258460, 1521258460213802700),
		Data:     nil,
	}, "/static/medaka.yml": &assets.File{
		Path:     "/static/medaka.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1521258460, 1521258460228000000),
		Data:     []byte(_Assets25a0b1481b5ef1da0a69fc5fc4cc4aa70e567123),
	}}, "")
