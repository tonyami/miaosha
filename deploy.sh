#!/bin/bash

NAME=miaosha
VERSION=latest

# 编译
build() {
  rm -f ./bin/${NAME}
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/${NAME} ./main.go
  echo -e '\033[32m编译成功... \033[0m'
}

# 构建镜像
build_image() {
  docker build -t ${NAME}:${VERSION} .
  echo -e '\033[32m构建镜像成功... \033[0m'
}

# 删除镜像
delete_image() {
  p=`docker images | grep ${NAME} | awk '{print $3}'`
  if [[ $p != '' ]];then
    docker rmi $p
    echo -e '\033[31m镜像已删除... \033[0m'
  fi
}


# 停止并删除容器
stop_container() {
  p=`docker ps | grep ${NAME} | awk '{print $1}'`
  if [[ $p != '' ]];then
    docker stop $p
    docker rm $p
    echo -e '\033[31m容器已停止并删除... \033[0m'
  fi
}

# 启动容器
start_container () {
  docker run --name ${NAME} -p 8080:8080 -v /etc/localtime:/etc/localtime -v `pwd`/conf.ini:/conf.ini -e GIN_MODE=release -d ${NAME}:${VERSION}
  echo -e '\033[32m容器已启动... \033[0m'
}

stop_container && delete_image && build && build_image && start_container
