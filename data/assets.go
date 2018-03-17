package data

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsd9dd3b7e23cf21210562389f49941f358763f9b1 = "# Lang file in English\r\n\r\nlanguage=English\r\nlanguage.name=English\r\n\r\nmedaka.lang.loaded=Loaded a language file in English\r\n\r\nmedaka.server.start=Starting minecraft server version {%0}"
var _Assets01626ec553c190a58d74919f914b2131f2025f75 = "# Lang file in Japanese\r\n\r\nlanguage=Japanese\r\nlanguage.name=日本語\r\n\r\nmedaka.lang.loaded=日本語の言語ファイルが読み込まれました。\r\n\r\nmedaka.server.start=minecraftサーバー バージョン {%0} を開始します。"
var _Assets25a0b1481b5ef1da0a69fc5fc4cc4aa70e567123 = "#Advanced config file for Medaka\r\n\r\nsettings:\r\n  language: \"eng\"\r\n  force-language: false\r\n  shutdown-message: \"Server closed\"\r\n  query-plugins: true\r\n  async-workers: auto\r\n  debug: true"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/static": []string{"medaka.yml"}, "/static/lang": []string{"eng.ini", "jpn.ini"}, "/": []string{"static"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521258460, 1521258460213802700),
		Data:     nil,
	}, "/static": &assets.File{
		Path:     "/static",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521268425, 1521268425775515000),
		Data:     nil,
	}, "/static/lang": &assets.File{
		Path:     "/static/lang",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521277615, 1521277615720293400),
		Data:     nil,
	}, "/static/lang/eng.ini": &assets.File{
		Path:     "/static/lang/eng.ini",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1521278716, 1521278716094174500),
		Data:     []byte(_Assetsd9dd3b7e23cf21210562389f49941f358763f9b1),
	}, "/static/lang/jpn.ini": &assets.File{
		Path:     "/static/lang/jpn.ini",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1521278707, 1521278707115585900),
		Data:     []byte(_Assets01626ec553c190a58d74919f914b2131f2025f75),
	}, "/static/medaka.yml": &assets.File{
		Path:     "/static/medaka.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1521258460, 1521258460228000000),
		Data:     []byte(_Assets25a0b1481b5ef1da0a69fc5fc4cc4aa70e567123),
	}}, "")
