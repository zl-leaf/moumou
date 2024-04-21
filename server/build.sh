#! /bin/bash

# 创建构建目录
build_server() {
  bin_path=$1"/bootstrap"
  if [ ! -d "$bin_path" ]; then
    mkdir -p $bin_path
  fi

  go build -o $bin_path"/app"
}

# 创建日志目录
build_logdir() {
  log_dir=$1"/bootstrap/log"
  if [ ! -d "$log_dir" ]; then
      mkdir -p $log_dir
    fi
}

# 创建启动shell
build_bootstrap() {
  script_path=$1"/bootstrap/server.sh"
  app_path=$1"/bootstrap/app"
  echo "nohup $app_path 1>/dev/null &" > $script_path
}

bin_root=$(PWD)
while [[ "$1" != "" ]]
do
  case $1 in
    --bin )
      bin_root=$2
      shift 2
      ;;
    * )
      break
      ;;
  esac
done

build_server $bin_root
build_logdir $bin_root
build_bootstrap $bin_root