package config

// ---------------- 定义变量 ----------------
var (
	siteName     = "HappeLife"
	blogPostsDir = "posts"
)

func GetBlogPostsDir() string {
	return blogPostsDir
}

func GetSiteName() string {
	return siteName
}
