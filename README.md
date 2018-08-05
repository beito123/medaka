# Medaka

![Image](https://cldup.com/WxrSZhzMED.png)

```
This is a project to learn Go for me
So This is very unstable!
```

Medaka is a MCBE(mcpe) server software written in Go.

These codes are written in Go based on [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp)).

## Build

If you haven't dep, Run the under cmd to install dep.

```
$ go get -u github.com/golang/dep/cmd/dep
```

Run the under commands to build.

```
$ go get github.com/beito123/medaka
$ cd $GOPATH/github.com/beito123/medaka
make
```

### Support environment

Windows, Linux, (maybe it work on Mac...)

```
Go(>v1.10), Dep(latest)
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

### PocketMine-MP

This project is based on PocketMine-MP!

- Server [PocketMine-MP](https://www.github.com/pmmp)

- Network for Raknet : [Raklib](https://github.com/pmmp/RakLib)

### Library

- Logger : [Logrus](https://github.com/Sirupsen/logrus)

- Color Console for Windows : [go-colorable](https://github.com/mattn/go-colorable)

- Load and Save configs : [Viper](https://github.com/spf13/viper)

- Safe cast for Config : [cast](https://github.com/spf13/cast)

- Embed assets : [go-assets](https://github.com/jessevdk/go-assets)

- Ordered Map : [orderedmap](https://github.com/secnot/orderedmap)

### Design reference

- [Go app of Kayac](https://github.com/kayac?language=go)

### Tools

- Dependency resolution : [Dep](https://github.com/golang/dep)

### Other

- Other, Developers who are providing codes for Medaka.

- And you

# Medaka (Japanese)

![Image](https://cldup.com/WxrSZhzMED.png)

```
このプロジェクトは私がGo言語を学習するためのものです。
そのため、非常に不安定です！
```

このソフトウェアは、Go言語で書かれたサーバーソフトウェアです。

これらのコードは [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp)) を元にGoで書かれています。

## ビルド

Depがない場合は、以下のコマンドを実行してDepをインストールください。

```
$ go get -u github.com/golang/dep/cmd/dep
```

ビルドするには以下のコマンドを実行してください。

```
$ go get github.com/beito123/medaka
$ cd $GOPATH/github.com/beito123/medaka
make
```

### サポート環境

Windows, Linux, (多分Macでも動くはず...)

```
Go(>v1.10), Dep(latest)
```

## ライセンス

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

### PocketMine-MP

このプロジェクトはPocketMine-MPを元にできています！

- サーバー : [PocketMine-MP](https://www.github.com/pmmp)

- Raknetのためのネットワーク : [Raklib](https://github.com/pmmp/RakLib)

### ライブラリ

- コンソールへのログの管理 : [Logrus](https://github.com/Sirupsen/logrus)

- Windowsでのコンソールの色付け : [go-colorable](https://github.com/mattn/go-colorable)

- Configの読み書き : [Viper](https://github.com/spf13/viper)

- Configの安全なキャスト : [cast](https://github.com/spf13/cast)

- リソースの埋込み : [go-assets](https://github.com/jessevdk/go-assets)

- 順序付きMap : [orderedmap](https://github.com/secnot/orderedmap)

### 設計の参考

- [KayacのGo製のアプリケーション](https://github.com/kayac?language=go)

### Tools

- 依存関係の解決 : [Dep](https://github.com/golang/dep)

### その他

- その他、コードを提供してくださっている開発者の方々

- And you
