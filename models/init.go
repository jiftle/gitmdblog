package models

import (
	"fmt"
	"gitmdblog/config"
	"os"
	"time"
)

// ---------------- 定义变量 ----------------
var (
	// 博客文章目录
	topicMarkdownFolder = config.GetBlogPostsDir()

	//Topics store all the topic 存储所有的文章
	Topics []*Topic

	//TopicsGroupByMonth store the topic by month
	TopicsGroupByMonth []*TopicMonth
	//TopicsGroupByTag store all the tag
	TopicsGroupByTag []*TopicTag

	TopicsGroupMonthListGroup []*TopicGroup
)

//Topic struct
type Topic struct {
	SiteName       string
	TopicID        string
	Title          string
	Time           time.Time
	LastModifyTime time.Time
	Tag            []*TopicTag
	Content        string
	TopicPath      string
	IsPublic       bool //true for public，false for protected
}

//TopicTag struct
type TopicTag struct {
	TagID   string
	TagName string
	Topics  []*Topic
}

//TopicMonth show the topic group by month
type TopicMonth struct {
	Month  string
	Topics []*Topic
}

//文章分组
type TopicGroup struct {
	TopicMonthLists []*TopicMonthList
}

// 文章月分组列表
type TopicMonthList struct {
	Title string
	Url   string
}

func init() {
	if err := InitTopicList(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
