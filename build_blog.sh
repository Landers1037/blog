#!/bin/bash
# build blog for yourself

# 使用go 普通方式编译而非mod方式
# 需要使用gopath
# 默认保证本项目放置在$GOPATH/src下 否则无法编译
current_dir=$(dirname $(pwd))
gopath=$(dirname "$current_dir")

export GOPATH=$gopath
export CGO_ENABLED=1
export GO111MODULE=on
# 输出全部的编译参数
echo "GOPATH: $(go env GOPATH)"
echo "GOPROXY: $(go env GOPROXY)"
echo "CGO_ENABLED: $(go env CGO_ENABLED)"
echo "GO111MODULE: $(go env GO111MODULE)"

echo "start to build app"
go build --mod=mod -a -ldflags="-w -s -extldflags -static" -trimpath -tags="nomsgpack go_json" -o app_blog app.go
echo "done"