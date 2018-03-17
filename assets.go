package medaka

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets795df3d638631b1a7721f68251ac97447c402b64 = "#Advanced config file for Medaka\r\n\r\nsettings:\r\n  language: \"eng\"\r\n  force-language: false\r\n  shutdown-message: \"Server closed\"\r\n  query-plugins: true\r\n  async-workers: auto\r\n  debug: true"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"data"}, "/data": []string{"medaka.yml"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521204910, 1521204910435016900),
		Data:     nil,
	}, "/data": &assets.File{
		Path:     "/data",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1521087574, 1521087574877021600),
		Data:     nil,
	}, "/data/medaka.yml": &assets.File{
		Path:     "/data/medaka.yml",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1521087788, 1521087788065667600),
		Data:     []byte(_Assets795df3d638631b1a7721f68251ac97447c402b64),
	}}, "")
