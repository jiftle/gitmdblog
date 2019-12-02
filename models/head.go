package models

import (
	"encoding/json"
	_ "errors"
	"fmt"
	logger "github.com/ccpaging/log4go"
	"github.com/noaway/dateparse"
	yaml "gopkg.in/yaml.v2"
	"strings"
	"time"
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
	fmt.Println(t)
	var thj tHeadJSON
	var headMeta HeadMeta
	if err := yaml.Unmarshal([]byte(tHeadStr), &headMeta); err != nil {
		logger.Warn("Notice: " + err.Error())
		fmt.Println("Notice: " + err.Error())
		return err, thj
	}
	//2019-11-28 Thu. 日记
	//fmt.Println(headMeta)
	//logger.Debug("标题: %s", headMeta.Title)
	thjTime := headMeta.Title
	if pos := strings.Index(thjTime, " "); pos > 0 {
		thjTime = thjTime[0:pos]
		//logger.Debug("时间: %s", thjTime)
	}

	// 设置时区
	denverLoc, _ := time.LoadLocation("Asia/Shanghai")
	// use time.Local global variable to store location
	time.Local = denverLoc
	tTime, err := dateparse.ParseAny(thjTime)
	if err != nil {
		logger.Warn(thjTime + "时间转换失败," + err.Error())
		thj.URL = headMeta.Title
		//		thj.URL = headMeta.Title
		thj.Time = "1999/01/01 01:01"
		thj.Tag = headMeta.Tags
	} else {
		//logger.Debug("Time: %v", tTime)
		thj.URL = headMeta.Title
		//thj.Time = tTime.Format("1999/01/01")
		//	thj.Time = "1999/01/01 01:01"
		thj.Time = tTime.Format("2006/01/02") + " 01:01"

		//logger.Debug("格式化后的时间: %v", tTime.Format("2006/01/02"))
		thj.Tag = headMeta.Tags
	}

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
