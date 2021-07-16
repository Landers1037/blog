#!/usr/bin/env bash
# build docker images in alpine

# to use in alpine must with ldflags
NAME=app_blog # binary opt
ENTRY=app.go # go build entry
IMAGE=blog # docker image name
TAG=v3 # docker tag

echo "start build blog binary"
echo "output is ${NAME}"

go_exist=$(which go)
if [[ -z ${go_exist} ]];then
  echo "go tools not exist"
  exit 1
fi

# start build
export GO111MODULE=on
export CGO_ENABLED=1
go build -o ${NAME} --ldflags "-w -s -extldflags -static" -trimpath -tags="nomsgpack go_json" ${ENTRY}

if [[ $? != 0 ]];then
  echo "build failed"
  exit 1
fi

# start npm build
echo -n "if you want to build frontend file into docker? [Y/N]"
read arg
# default build html
if [[ $arg == "n" || $arg == "N" ]];then
  echo "skip build frontend file"
else
  echo "build total frontend file"
  echo "file will be build into /app/html"
  echo "your app.ini should set try_file: /app/html/index.html"
  cd html && npm install && npm run build
  if [[ $? != 0 ]];then
    echo "npm build failed"
    exit 1
  fi
  if [[ -d ../dist ]];then
    rm -rf ../dist
  fi
  echo "start to mv dist"
  mv dist ../dist
  # exit current dir
  cd ..
fi

echo "start build docker image"
if [[ ! -f Dockerfile ]];then
  echo "Dockerfile not exist"
  exit 1
fi

docker_exist=$(which docker)
if [[ -z ${docker_exist} ]];then
  echo "docker not exist"
  exit 1
fi

# start build docker image
# default use alpine:latest
if [[ ! -d ./dist ]];then
  mkdir dist
  touch dist/frontend_file.txt
fi
docker build -t ${IMAGE}:${TAG} .
