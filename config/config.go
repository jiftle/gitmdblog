package config

// ---------------- 定义变量 ----------------
var (
	siteName     = "HappeLife"
	blogPostsDir = "posts"
	//blogPostsDir = "/media/jiftle/work/work/git/coding_net/grocery/dailylog"
//	blogPostsDir = "/usrlocal/git/coding_net/grocery/dailylog"
)

// 获取博客文章目录
func GetBlogPostsDir() string {
	return blogPostsDir
}

func GetSiteName() string {
	return siteName
}
