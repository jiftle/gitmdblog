package models

import (
	"bufio"
	"bytes"
	"errors"
	"gitmdblog/config"
	"os"
	"path/filepath"
	"strings"
	"time"

	logger "github.com/ccpaging/log4go"
	"gopkg.in/russross/blackfriday.v2"
	//"github.com/jiftle/blackfriday"
)

//GetTopicByPath read the topic by path
// 读取文章内容
func GetTopicByPath(path string) (*Topic, error) {
	// 打开文件
	fp, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, errors.New(path + "：" + err.Error())
	}
	defer fp.Close()

	// 初始化结构体
	t := &Topic{
		TopicID:  strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Title:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		IsPublic: true,
	}

	var tHeadStr string
	var i int
	i = 0
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		i += 1
		s := scanner.Text()
		if i == 1 && s != "---" {
			return t, nil
		}
		if i != 1 && s == "---" {
			break
		}
		if s == "---" {
			continue
		}
		tHeadStr += s
		tHeadStr += "\n"
		if len(s) == 0 {
			break
		}
	}

	var thj tHeadJSON

	// ---------------- 截取头部信息 ----------------
	err, thj = parseTopicHead(tHeadStr, *t)
	if err != nil {
		return t, nil
	}

	// ------------ 赋值 ------------------
	//t.TopicID = thj.URL
	// if t.TopicID == "" {
	// 	return nil, errors.New(t.Title + "：" + err.Error())
	// }
	t.Time, err = time.Parse("2006/01/02 15:04", thj.Time)
	if err != nil {
		logger.Error(t.Title + "：" + err.Error())
		return nil, errors.New(t.Title + "：" + err.Error())
	}
	if strings.Compare(thj.IsPublic, "no") == 0 {
		t.IsPublic = false
	}
	tagArray := strings.Split(thj.Tag, ",")
	var isFind bool
	for _, tagName := range tagArray {
		tagName = strings.TrimSpace(tagName)
		if len(tagName) == 0 {
			continue
		}
		isFind = false
		for kc := range TopicsGroupByTag {
			if strings.Compare(strings.ToLower(tagName), TopicsGroupByTag[kc].TagID) == 0 {
				t.Tag = append(t.Tag, TopicsGroupByTag[kc])
				isFind = true
				break
			}
		}
		if isFind == false {
			tt := &TopicTag{TagID: strings.ToLower(tagName), TagName: tagName}
			t.Tag = append(t.Tag, tt)
			TopicsGroupByTag = append(TopicsGroupByTag, tt)
		}
	}
	var content bytes.Buffer
	for scanner.Scan() {
		content.Write(scanner.Bytes())
		content.WriteString("\n")
	}

	// ------------------ 使用blackfriday引擎喧嚷markdown ------------------
	input := content.Bytes()
	output := blackfriday.Run(input)
	stext := string(output)

	//fmt.Println(stext)

	// ------------------ 直接输出到页面 -----------------------
	t.Content = string(stext)
	t.TopicPath = path

	finfo, _ := os.Stat(path)
	lastModTime := finfo.ModTime()
	if lastModTime.Unix()-t.Time.Unix() > 7*86400 && time.Now().Unix()-lastModTime.Unix() < 365*86400 {
		t.LastModifyTime = lastModTime
	}

	// ----------------- 站点配置 -----------------
	t.SiteName = config.GetSiteName()

	return t, nil
}

//InitTopicList load all the topic on init
// 初始化文章列表
func InitTopicList() error {
	// 清空文章存储
	Topics = Topics[:0]

	// 按月份分组
	TopicsGroupByMonth = TopicsGroupByMonth[:0]
	// 按标签分组
	TopicsGroupByTag = TopicsGroupByTag[:0]

	// 遍历文件
	return filepath.Walk(topicMarkdownFolder, func(path string, info os.FileInfo, err error) error {
		// 过滤目录或无关文件
		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		// 根据路径获取文章
		t, err := GetTopicByPath(path)
		if err != nil {
			return err
		}
		// 文件路径为空，直接返回
		if t.TopicPath == "" {
			return nil
		}

		// 按标签分组
		SetTopicToTag(t)
		// 按月份分组
		SetTopicToMonth(t)

		//append topics desc
		for i := range Topics {
			if t.Time.After(Topics[i].Time) {
				Topics = append(Topics, nil)
				copy(Topics[i+1:], Topics[i:])
				Topics[i] = t
				return nil
			}
		}

		// 加入队列末尾
		Topics = append(Topics, t)
		return nil
	})
}

//SetTopicToTag set topic to tag struct 设置标签结构
func SetTopicToTag(t *Topic) {
	if t.IsPublic == false {
		return
	}
	for i := range t.Tag {
		for k := range TopicsGroupByTag {
			if TopicsGroupByTag[k].TagID != t.Tag[i].TagID {
				continue
			}
			isFind := false
			for j := range TopicsGroupByTag[k].Topics {
				if t.Time.After(TopicsGroupByTag[k].Topics[j].Time) {
					TopicsGroupByTag[k].Topics = append(TopicsGroupByTag[k].Topics, nil)
					copy(TopicsGroupByTag[k].Topics[j+1:], TopicsGroupByTag[k].Topics[j:])
					TopicsGroupByTag[k].Topics[j] = t
					isFind = true
					break
				}
			}
			if isFind == false {
				TopicsGroupByTag[k].Topics = append(TopicsGroupByTag[k].Topics, t)
			}
			break
		}
	}
}

//SetTopicToMonth set topic to month struct 月份结构
func SetTopicToMonth(t *Topic) {
	if t.IsPublic == false {
		return
	}

	month := t.Time.Format("2006-01")
	logger.Debug("---- >>> --- >>> 月份分组 %v, topic=%v", month, t.Title)
	tm := &TopicMonth{}
	for _, m := range TopicsGroupByMonth {
		//	logger.Debug("---- >>> --- >>> xx 进度循环[月份分组] %v", m)
		if m.Month == month {
			tm = m
		}
	}
	if tm.Month == "" {
		//logger.Debug("---- >>> --- >>> xx 进度循环[月份分组]  month != nil")
		tm.Month = month
		isFind := false
		for i := range TopicsGroupByMonth {
			if strings.Compare(tm.Month, TopicsGroupByMonth[i].Month) > 0 {
				//		logger.Debug("月份分组: %v", TopicsGroupByMonth[i].Month)
				TopicsGroupByMonth = append(TopicsGroupByMonth, nil)
				copy(TopicsGroupByMonth[i+1:], TopicsGroupByMonth[i:])
				TopicsGroupByMonth[i] = tm
				isFind = true
				break
			}
		}
		if isFind == false {
			TopicsGroupByMonth = append(TopicsGroupByMonth, tm)
		}
	}

	// 遍历分组文章列表
	for i := range tm.Topics {
		//	logger.Debug("---- >>> --- >>> xx 进度循环[月份分组]  tm.Topics")
		if t.Time.After(tm.Topics[i].Time) {
			//	logger.Debug("---- >>> --- >>>  %v After %v", t.Time, tm.Topics[i].Time)
			// 追加到列表里
			tm.Topics = append(tm.Topics, nil)
			// 后继元素整体后移
			copy(tm.Topics[i+1:], tm.Topics[i:])
			// 插入新文章
			tm.Topics[i] = t
			return
		}
	}
	tm.Topics = append(tm.Topics, t)
}
