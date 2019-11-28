package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ---------------- 定义变量 ----------------
var (
	siteName     = "HappeLife"
	blogPostsDir = "posts"
	//blogPostsDir = "/media/jiftle/work/work/git/coding_net/grocery/dailylog"
//	blogPostsDir = "/usrlocal/git/coding_net/grocery/dailylog"
)

// 获取博客文章目录
func GetBlogPostsDir() string {
	readConfig()
	return blogPostsDir
}

func GetSiteName() string {
	return siteName
}

func readConfig() {

	// 读取文件
	//-------------- properties 文件 ----------------
	config := viper.New()
	config.AddConfigPath("./conf")
	config.SetConfigType("yaml")
	config.SetConfigName("gitmdblog")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("[x] 配置文件不存在，创建默认配置文件")
	}

	//---------- 写入配置文件 -------------
	bakconf := fmt.Sprintf("./conf/gitmdblog.yaml")
	if err := config.WriteConfigAs(bakconf); err != nil {
		fmt.Println("[x] 创建默认配置文件失败")
		panic(err)
		config.Set("postsdir", "posts")
		config.WriteConfig()
	}
	blogPostsDir = config.GetString("postsdir")
	fmt.Println("[v] 读取配置文件成功")
	//	msg := fmt.Sprintf("posts dir: %s\n", blogPostsDir)
	//	fmt.Print(msg)
}
