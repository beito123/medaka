# Medaka

![Image](https://cldup.com/dWMoUCl7CU.png)

    This is a project to learn Go for me
    So This is very unstable!

Medaka is a MCBE(mcpe) server software written in Go.

These codes were rewritten for Go on the basis of [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp))

## Build

If you haven't dep, Run the under cmd to install dep.

    $ go get -u github.com/golang/dep/cmd/dep

Run the under commands to build.

    $ go get github.com/beito123/medaka
    $ cd $GOPATH/github.com/beito123/medaka
    make

### Support environment

    Go(v1.10 >), Dep

## License

There is this project the under LGPLv3 License.

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

## Development Plan

### Policy

#### Actively uses libraries

I can't develop alone.(no time & no skill)

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
- Embed assets : [go-assets](https://github.com/jessevdk/go-assets)

### Design reference

- [Go app of Kayac](https://github.com/kayac?language=go)

### Tools

- Dependency resolution : [Dep](https://github.com/golang/dep)

### Other

- Other, Developers who are providing codes for Medaka.
- And you

# Medaka (Japanese)

![Image](https://cldup.com/dWMoUCl7CU.png)

    このプロジェクトは私がGo言語を学習するためのものです。
    そのため、非常に不安定です！

このソフトウェアは、Go言語で書かれたサーバーソフトウェアです。

これらのコードは [PocketMine-MP](https://www.github.com/pmmp) (by [PocketMine-Team](https://www.github.com/pmmp)) をGoに書き直したものです。

## ビルド

Depがない場合は、以下のコマンドを実行してDepをインストールください。

    $ go get -u github.com/golang/dep/cmd/dep

ビルドするには以下のコマンドを実行してください。

    $ go get github.com/beito123/medaka
    $ cd $GOPATH/github.com/beito123/medaka
    make

### サポート環境

    Go(v1.10 >), Dep

## ライセンス

このプロジェクトはLGPLv3ライセンス下にあります。

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

## 開発プラン

### 方針

#### 積極的なライブラリの使用

私一人で全て書くのは限界があるためです。(時間も技術もない)

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
- リソースの埋込み : [go-assets](https://github.com/jessevdk/go-assets)

### 設計の参考

- [KayacのGo製のアプリケーション](https://github.com/kayac?language=go)

### Tools

- 依存関係の解決 : [Dep](https://github.com/golang/dep)

### その他

- その他、コードを提供してくださっている開発者の方々
- And you
