package config

// ---------------- 定义变量 ----------------
var (
	blogPostsDir = "posts"
)

func GetBlogPostsDir() string {
	return blogPostsDir
}
