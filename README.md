# GitMarkdownBlog

采用`Go`语言和`Markdown`实现的一个简易博客系统(/posts/)，主要包括以下功能：

- 按日期、按标签展现文章列表
- 首页、文章详情页

功能比较简单，不需要依赖数据库，不需要管理后台，使用者只需要关注文章内容的书写即可，同时写好的文章可直接在`Github`上查看。


## 如何写文章？

文章采用`Markdown`的写法，需要先了解`Markdown`的写法，基本用法与文档示例写法可查看[Markdown基本用法](/posts/Markdown基本用法.md)。除此之外有几点需要注意：

1、文章放在`posts`目录中，文件夹可多层嵌套（无影响），文件需以md为后缀，文件名即文章标题。
2、`md`文件头部需写入文章头部，文章头部和文章正文以换行区分，示例如下：

 ```
---
title: 2019-11-25 Mon. 日记
tags: Markdown,diary
notebook: diary
---
```

 文章正文

文章头部采用`json`来描述文章信息，字段定义如下：

字段   | 必选 | 说明
---    | --- | ---
url    | 是  | 文章URL
time   | 是  |  文章发表时间
tag    | 是  | 标签，多个标签用英文逗号分隔
public | 否  | 为no的时候表示文章不可被浏览器访问到

## 启动

```
go run main.go
```

说明：

- 不需要生成静态页面时将`main.go`中`isCreateHTML`设置为false，默认为false。
- `domain`用来定义模版页链接前缀，可设置为空

## 效果截图

![截图](https://raw.githubusercontent.com/jiftle/gitmdblog/master/static/uploads/screenshot01.png)



## 感谢(创造者)

- github.com/pengbotao/itopic.go

