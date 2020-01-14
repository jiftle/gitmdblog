package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ---------------- 定义变量 ----------------
var (
	siteName      = "HappeLife"
	blogPostsDir  = "posts"
	refreshSecond = 60 * 5
	host          = "0.0.0.0:8080"
)

// 获取博客文章目录
func GetBlogPostsDir() string {
	readConfig()
	return blogPostsDir
}

func GetSiteName() string {
	return siteName
}

func GetRefreshSecond() int {
	return refreshSecond
}

func GetHost() string {
	return host
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

		//---------- 写入配置文件 -------------
		config.Set("postsdir", "posts")
		config.Set("refreshsecond", 300)
		config.Set("host", "0.0.0.0:8080")
		bakconf := fmt.Sprintf("./conf/gitmdblog.yaml")
		if err := config.WriteConfigAs(bakconf); err != nil {
			fmt.Println("[x] 创建默认配置文件失败")
			panic(err)
		}
	}
	blogPostsDir = config.GetString("postsdir")
	refreshSecond = config.GetInt("refreshSecond")
	listenIp := config.GetString("listen_ip")
	listenPort := config.GetString("listen_port")

	host = fmt.Sprintf("%v:%v", listenIp, listenPort)
	msg := fmt.Sprintf("[v] 读取配置文件成功, posts dir: %s\n", blogPostsDir)
	fmt.Print(msg)
}
