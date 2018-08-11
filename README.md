# Medaka

![Image](https://cldup.com/WxrSZhzMED.png)

```
This is a project to learn Go for me
So This is very unstable!
```

Medaka is a MCBE(mcpe) server software written in Go.

These codes are written in Go based on [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp)).

## Build

If you haven't dep, Run the under command to get dep.

```
$ go get -u github.com/golang/dep/cmd/dep
```

Run the under commands to build.

```
$ go get -u github.com/beito123/medaka
$ cd $GOPATH/src/github.com/beito123/medaka
$ make deps
$ make
```

- Manual build: [Link](https://gist.github.com/beito123/609f4bf2f25f8c24541e8bb47c78cb92)

### Support environment

Windows, Linux, (maybe it work on Mac...)


```
Go(v1.10^), Dep(v0.5.0), GNU make(v4.2) or GnuWin32(v3.81)
```

## License

There is this project the under LGPLv3 License.

```
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```

## Special Thanks

Thank you so much!

### Based on

- [PocketMine-MP](https://www.github.com/pmmp) (by [PMMP-Team](https://github.com/pmmp) License:[LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.ja.html)): A MCBE server software in PHP

### Library

- [Logrus](https://github.com/Sirupsen/logrus) (by [sirupsen](https://github.com/sirupsen) License: [MIT License](https://opensource.org/licenses/mit-license.php)): A logger library

- [go-colorable](https://github.com/mattn/go-colorable) (by [mattn](https://github.com/mattn) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Color Console for Windows

- [Viper](https://github.com/spf13/viper) (by [spf13](https://github.com/spf13) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Internal process in Config

- [cast](https://github.com/spf13/cast) (by [spf13](https://github.com/spf13) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Safe casting for Config

- [go-assets](https://github.com/jessevdk/go-assets) (by [jessevdk](https://github.com/jessevdk) License: [BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause)): Embed assets

- [orderedmap](https://github.com/secnot/orderedmap) (by [secnot](https://github.com/secnot) License: [MIT License](https://opensource.org/licenses/mit-license.php)): A ordered map

### Design reference

- [Apps of Kayac](https://github.com/kayac?language=go)

### Tools

- [Dep](https://github.com/golang/dep) (by [golang-team](https://github.com/golang/)): Dependency resolution

### Other

- Developers who are providing codes for Medaka.

- And you

## Medaka (Japanese)

![Image](https://cldup.com/WxrSZhzMED.png)

```
このプロジェクトは私がGo言語を学習するためのものです。
そのため、非常に不安定です！
```

このソフトウェアは、Go言語で書かれたサーバーソフトウェアです。

これらのコードは [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp)) を元にGoで書かれています。

### ビルド

Depがない場合は、以下のコマンドを実行してDepをインストールください。

```
$ go get -u github.com/golang/dep/cmd/dep
```

ビルドするには以下のコマンドを実行してください。

```
$ go get -u github.com/beito123/medaka
$ cd $GOPATH/src/github.com/beito123/medaka
$ make deps
$ make
```

- 手動ビルド: [Link](https://gist.github.com/beito123/609f4bf2f25f8c24541e8bb47c78cb92)

#### サポート環境

Windows, Linux, (多分Macでも動くはず...)

```
Go(v1.10^), Dep(v0.5.0), GNU make(v4.2) or GnuWin32(v3.81)
```

### ライセンス

このプロジェクトはLGPLv3ライセンス下にあります。

```
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```

## 謝辞

Thank you so much!

### Base

- [PocketMine-MP](https://www.github.com/pmmp) (by [PMMP-Team](https://github.com/pmmp) License:[LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.ja.html)): PHPで書かれたMCBEサーバー

### ライブラリ

- [Logrus](https://github.com/Sirupsen/logrus) (by [sirupsen](https://github.com/sirupsen) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Loggerライブラリ

- [go-colorable](https://github.com/mattn/go-colorable) (by [mattn](https://github.com/mattn) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Windowsコンソールに色を付けるライブラリ

- [Viper](https://github.com/spf13/viper) (by [spf13](https://github.com/spf13) License: [MIT License](https://opensource.org/licenses/mit-license.php)): Configの内部処理

- [cast](https://github.com/spf13/cast) (by [spf13](https://github.com/spf13) License: [MIT License](https://opensource.org/licenses/mit-license.php)): 安全なキャストができるライブラリ

- [go-assets](https://github.com/jessevdk/go-assets) (by [jessevdk](https://github.com/jessevdk) License: [BSD-3-Clause](https://opensource.org/licenses/BSD-3-Clause)): ファイルの埋め込み

- [orderedmap](https://github.com/secnot/orderedmap) (by [secnot](https://github.com/secnot) License: [MIT License](https://opensource.org/licenses/mit-license.php)): 順序付きマップ

### 設計の参考

- [Apps of Kayac](https://github.com/kayac?language=go)

### ツール

- [Dep](https://github.com/golang/dep) (by [golang-team](https://github.com/golang/)): 依存関係の解決

### その他

- コードを提供してくださっている開発者の方々

- And you
