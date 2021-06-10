<!--
order: 1
-->

# 安装 IRITA

本文档将说明如何安装 `irita` (包含节点和控制台) 。

## 安装 Go

按照[官方文档](https://golang.org/doc/install)安装 `go`。

设置 `$PATH` 环境变量，例如：

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
source ~/.bash_profile
```

> _注意_: IRITA 需要 `Go 1.15+`。

## 安装软件

安装最新版本的 `irita`。确保 `git checkout` 了正确的发行版本。

```bash
git clone https://github.com/bianjieai/irita.git
cd irita && git checkout <latest-release-tag>
make install
```

这将安装 `irita` 二进制文件。验证安装是否成功：

```bash
irita version --long
```

例如，`irita` 应该输出类似于以下的内容：

```text
name: irita
server_name: irita
version: 2.0.0
commit: 5ce9b33ec68c65a5fbcf193ced2c318323218218
build_tags: netgo,ledger
go: go version go1.15 darwin/amd64
```
