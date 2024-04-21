#! /bin/bash

# 创建构建目录
build_server() {
  bin_path=$1"/dist"
  if [ ! -d "$bin_path" ]; then
    mkdir -p $bin_path
  fi

  export VITE_APP_PATH=$bin_path && npm run build
}

shell_path=$(realpath $0)
project_root=$(dirname "$shell_path")
cd $project_root
bin_root=$project_root
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

npm install
build_server $bin_root
