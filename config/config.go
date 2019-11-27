package config

// ---------------- 定义变量 ----------------
var (
	siteName = "HappeLife"
	//blogPostsDir = "posts"
	blogPostsDir = "/media/jiftle/work/work/git/coding_net/grocery/dailylog"
)

func GetBlogPostsDir() string {
	return blogPostsDir
}

func GetSiteName() string {
	return siteName
}
