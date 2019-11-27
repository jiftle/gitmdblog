package main

import (
	"bytes"
	"fmt"
	"gitmdblog/config"
	"gitmdblog/models"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	logger "github.com/ccpaging/log4go"
)

var (
	siteName     = config.GetSiteName()
	host         = "127.0.0.1:8001"
	isCreateHTML = false
	htmlPrefix   = "../itopic.org" //without last slash
	domain       = ""
	githubURL    = "https://github.com/pengbotao/itopic.go"
)

func init() {
	logger.LoadConfiguration("conf/log4go.xml")
}

func main() {
	logger.Info("%s", "--- in main ---")
	router := loadHTTPRouter()

	// 定时器，定时刷新
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for range ticker.C {
			fmt.Println("---> ticker round **>>**>>")
			// 初始化文章列表
			models.InitTopicList()
			hr := loadHTTPRouter()
			router = hr
		}
	}()

	// ------------------ 网站图标，静态资源 -----------------------
	http.Handle("/favicon.ico", http.FileServer(http.Dir("static")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if pos := strings.LastIndex(path, "."); pos > 0 {
			path = path[0:pos]
		}
		if bf, ok := router[path]; ok {
			if strings.Compare("/sitemap", path) == 0 {
				w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
			} else {
				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			}
			w.Write(bf.Bytes())
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		models.InitTopicList()
		hr := loadHTTPRouter()
		router = hr
		http.NotFound(w, r)
	})

	fmt.Printf("The topic server is running at http://%s\n", host)
	fmt.Printf("Quit the server with Control-C\n\n")
	if err := http.ListenAndServe(host, nil); err != nil {
		fmt.Print(err)
	}
}

// 加载路由
func loadHTTPRouter() map[string]bytes.Buffer {
	router := make(map[string]bytes.Buffer)
	var tpl *template.Template

	// 指定模板文件存放目录
	tpl, err := template.ParseGlob("views/*.tpl")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pages []map[string]string

	//首页 homepage router
	var buff bytes.Buffer
	topicCnt := len(models.Topics)
	topicDivCnt := topicCnt / 2

	var topicsLeft []*models.TopicMonth  //左侧
	var topicsRight []*models.TopicMonth //右侧

	if topicDivCnt > 0 {
		t := 0
		isSplit := false
		for i := range models.TopicsGroupByMonth {
			if t > topicDivCnt {
				isSplit = true
				topicsLeft = models.TopicsGroupByMonth[0:i]
				topicsRight = models.TopicsGroupByMonth[i:]
				break
			}
			t += len(models.TopicsGroupByMonth[i].Topics)
		}
		if isSplit == false {
			topicsLeft = models.TopicsGroupByMonth
		}
	} else {
		topicsLeft = models.TopicsGroupByMonth
	}

	// 网页模板
	if err := tpl.ExecuteTemplate(&buff, "index.tpl", map[string]interface{}{
		"siteName":  siteName,
		"topics_l":  topicsLeft,
		"topics_r":  topicsRight,
		"domain":    domain,
		"time":      time.Now(),
		"githubURL": githubURL,
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router["/"] = buff
	router["/index"] = buff

	pages = append(pages, map[string]string{
		"loc":        domain + "/",
		"lastmod":    time.Now().Format("2006-01-02"),
		"changefreq": "weekly",
		"priority":   "1",
	})

	// 文章页 topic router
	for i := range models.Topics {
		if models.Topics[i].IsPublic == false {
			continue
		}
		var buff bytes.Buffer
		err := tpl.ExecuteTemplate(&buff, "topic.tpl", map[string]interface{}{
			"topic":     models.Topics[i],
			"domain":    domain,
			"time":      time.Now(),
			"githubURL": githubURL,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		router["/"+models.Topics[i].TopicID] = buff
		pages = append(pages, map[string]string{
			"loc":        domain + "/" + models.Topics[i].TopicID + ".html",
			"lastmod":    models.Topics[i].Time.Format("2006-01-02"),
			"changefreq": "monthly",
			"priority":   "0.9",
		})
	}
	//month router
	for i := range models.TopicsGroupByMonth {
		var buff bytes.Buffer
		err := tpl.ExecuteTemplate(&buff, "list.tpl", map[string]interface{}{
			"title":  models.TopicsGroupByMonth[i].Month,
			"topics": models.TopicsGroupByMonth[i].Topics,
			"domain": domain,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		router["/"+models.TopicsGroupByMonth[i].Month] = buff
		pages = append(pages, map[string]string{
			"loc":        domain + "/" + models.TopicsGroupByMonth[i].Month + ".html",
			"lastmod":    time.Now().Format("2006-01-02"),
			"changefreq": "monthly",
			"priority":   "0.2",
		})
	}

	//tag 标签 router
	for i := range models.TopicsGroupByTag {
		var buff bytes.Buffer
		err := tpl.ExecuteTemplate(&buff, "list.tpl", map[string]interface{}{
			"title":  models.TopicsGroupByTag[i].TagName,
			"topics": models.TopicsGroupByTag[i].Topics,
			"domain": domain,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		router["/tag/"+models.TopicsGroupByTag[i].TagID] = buff
		pages = append(pages, map[string]string{
			"loc":        domain + "/tag/" + url.QueryEscape(models.TopicsGroupByTag[i].TagID) + ".html",
			"lastmod":    time.Now().Format("2006-01-02"),
			"changefreq": "monthly",
			"priority":   "0.2",
		})
	}
	//sitemap
	var sitemapBuff bytes.Buffer
	if err := tpl.ExecuteTemplate(&sitemapBuff, "sitemap.tpl", map[string]interface{}{
		"pages": pages,
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router["/sitemap"] = sitemapBuff
	//create html
	if isCreateHTML == true {
		go generateHTML(router)
	}
	return router
}

func generateHTML(router map[string]bytes.Buffer) {
	for k, v := range router {
		if k == "/" {
			writeFile(htmlPrefix+k+"index.html", v)
		} else {
			if k == "/sitemap" {
				writeFile(htmlPrefix+k+".xml", v)
			} else {
				writeFile(htmlPrefix+k+".html", v)
			}
		}
	}
	//copy static folder
	err := copyDir("./static", htmlPrefix+"/static")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func writeFile(filename string, content bytes.Buffer) {
	_, err := os.Stat(path.Dir(filename))
	if os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(filename), 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	file, _ := os.Create(filename)
	content.WriteTo(file)
}

func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	err = out.Sync()
	if err != nil {
		return err
	}

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return err
	}

	return
}

func copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dst, si.Mode())
		if err != nil {
			return err
		}
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return
}
