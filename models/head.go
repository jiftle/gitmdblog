package models

import (
	"encoding/json"
	"fmt"
	logger "github.com/ccpaging/log4go"
	yaml "gopkg.in/yaml.v2"
	"strings"
)

type tHeadJSON struct {
	URL      string
	Time     string
	Tag      string
	IsPublic string `json:"public"`
}

type HeadMeta struct {
	Title string
	Tags  string
}

// 分析文章头 --yaml
func parseTopicHead_YAML(tHeadStr string, t Topic) (error, tHeadJSON) {
	//fmt.Println(tHeadStr)
	var thj tHeadJSON
	var headMeta HeadMeta
	if err := yaml.Unmarshal([]byte(tHeadStr), &headMeta); err != nil {
		fmt.Println("Notice: " + err.Error())
		logger.Warn("Notice: " + err.Error())
		return err, thj
	}
	//fmt.Println(headMeta)

	thj.URL = headMeta.Title
	thj.Time = "1999/01/01 01:01"
	thj.Tag = headMeta.Tags

	//fmt.Println(thj.URL)
	//fmt.Println(thj)
	return nil, thj
}

// 分析文章头
func parseTopicHead_JSON(tHeadStr string, t Topic) (error, tHeadJSON) {
	tHeadStr = strings.Trim(tHeadStr, "```")
	var thj tHeadJSON
	if err := json.Unmarshal([]byte(tHeadStr), &thj); err != nil {
		fmt.Println("Notice: " + err.Error())
		return err, thj
	}

	fmt.Println("---> 解析文章头成功,")
	fmt.Println(thj)
	return nil, thj
}

// 分析文章头
func parseTopicHead(tHeadStr string, t Topic) (error, tHeadJSON) {
	var typ string
	var err error
	var thj tHeadJSON
	typ = "json1"

	if typ == "json" {
		err, thj = parseTopicHead_JSON(tHeadStr, t)
	} else {
		err, thj = parseTopicHead_YAML(tHeadStr, t)
	}

	return err, thj
}
