#!/bin/bash
# -----------------------------------------------------------------
# FileName: sh-publish.sh
# Date: 2020-11-09
# Author: jiftle
# Description: 
# -----------------------------------------------------------------

echo "  |--> 编译项目"
cd ..
go build

# ----------------- 发布到vim插件目录 ------------------------
echo "  |--> 发布到vim插件目录"
cp -f ./gitmdblog ~/.vim/plugged/vim-jiftle-gitmdblog/blogserv/

