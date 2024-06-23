package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

// 表单传过来的值绑定结构体
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// 时间戳转化为日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

func main() {
	r := gin.Default()
	//加载引擎
	r.Static("/static", "./static")
	//第一个/static表示路由,第二个./static表示路径
	//自定义模板函数，注意这个函数放在加载模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	r.LoadHTMLGlob("templates/**/*")
	//加载模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "前台首页",
			"score": 89,
			"msg":   "msg",
			"hobby": []string{"吃饭", "睡觉", "拉屎"},
			"test":  []string{},
			"news": &Article{
				Title:   "哈哈哈",
				Content: "我我我",
			},
			"data": 1719020499,
			"newslist": []interface{}{
				&Article{
					Title:   "11111",
					Content: "a1111",
				},
				&Article{
					Title:   "2222",
					Content: "a2222",
				},
			},
		})
	})

	//动态路由传值
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(200, "%v", cid)
	})

	r.GET("/path", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		page := c.DefaultQuery("page", "1")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"page":     page,
		})
	})

	//get请求传值query是查询路径后的传值为id的值，defaultQuery是查询路径后的值，没有则自定义一个string类型值
	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	//post演示
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	//获取get,post的值传到结构体上
	//http://localhost:8080/getUser?username=lisi&password=99999，这是get传值
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	})

	//获取post xml数据
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}
		xmlSliceData, _ := c.GetRawData() //获取c.request.Body的请求数据
		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
	})

	//获取表单post来的数据，只有通过post请求才能触发方法,不能刷新刷新的话要重新post
	r.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻",
			"news":  news,
		})
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.GET("/admin/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻",
			"news":  news,
		})
	})
	r.Run(":8080")
}
