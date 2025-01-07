#!/bin/bash

# 输出文件夹
OUTPUT_DIR="build"
mkdir -p $OUTPUT_DIR

# 目标平台和架构
TARGETS=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64")

# 程序名称
APP_NAME="imagepdf"

echo "开始编译程序..."

# 循环编译目标
for TARGET in "${TARGETS[@]}"; do
    OS=$(echo $TARGET | cut -d'/' -f1)
    ARCH=$(echo $TARGET | cut -d'/' -f2)
    
    OUTPUT_FILE="${OUTPUT_DIR}/${APP_NAME}-${OS}-${ARCH}"
    if [ "$OS" == "windows" ]; then
        OUTPUT_FILE="${OUTPUT_FILE}.exe"
    fi

    echo "编译 $OS/$ARCH..."
    GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_FILE ./
    if [ $? -ne 0 ]; then
        echo "编译 $OS/$ARCH 失败!"
    else
        echo "编译完成: $OUTPUT_FILE"
    fi
done

echo "所有目标平台编译完成！"
