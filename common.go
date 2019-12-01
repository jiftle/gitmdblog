package main

import (
	"fmt"
	logger "github.com/ccpaging/log4go"
	"os"
)

func PrintLOGO() {
	//--------------- 自定义文字LOGO ----------------
	fmt.Println()
	fmt.Println(`  ________  .__    __                  .___ __________  .__                     `)
	fmt.Println(` /  _____/  |__| _/  |_    _____     __| _/ \______   \ |  |     ____      ____ `)
	fmt.Println(`/   \  ___  |  | \   __\  /     \   / __ |   |    |  _/ |  |    /  _ \    / ___\`)
	fmt.Println(`\    \_\  \ |  |  |  |   |  Y Y  \ / /_/ |   |    |   \ |  |__ (  <_> )  / /_/  >`)
	fmt.Println(` \______  / |__|  |__|   |__|_|  / \____ |   |______  / |____/  \____/   \___  /`)
	fmt.Println(`        \/                     \/       \/          \/                  /_____/ `)
	fmt.Println(`                                                         --writed by jiftle     `)
	fmt.Println()
}
func DirIsExisted(filename string) bool {
	file := filename

	// 判断文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		logger.Info("file: %v isn't exist.\n", file)
		return false
	}
	return true
}
func init() {
	// 日志初始化
	logger.LoadConfiguration("conf/log4go.xml")
}
