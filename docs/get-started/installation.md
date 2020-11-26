# 安装 IRITA

本文档将说明如何安装irita到您的系统上。

## 安装 Go

按照[官方文档](https://golang.org/doc/install)安装`go`。

记住要设置`$PATH`环境变量，例如：

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
source ~/.bash_profile
```

> _注意_: irita需要`Go 1.14+`。

## 安装软件

接下来，让我们安装最新版本的`irita`。 确保您`git checkout`了正确的发行版本。

```bash
git clone -b <latest-release-tag> https://github.com/bianjieai/irita.git
cd irita && make install
```

这将安装irita二进制文件。 验证安装是否成功：

```bash
irita version --long
```

例如，`irita` 应该输出类似于以下的内容：

```text
name: irita
server_name: irita
version: 0.5.0-40-g404514a
commit: 404514ad296cbda1fa88402942a5ea4e694609d6
build_tags: netgo,ledger
go: go version go1.15.3 darwin/amd64
```
