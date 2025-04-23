package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world!")
	})

	// 通过Context的Param方法获取API参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/") // 去除字符串前后的/
		c.String(http.StatusOK, name+" is "+action)
	})

	// URL参数通过DefaultQuery()或Query()获取
	// DefaultQuery()若参数不存在则返回默认值，Query()若不存在则返回空串
	r.GET("/custom", func(c *gin.Context) {
		// http://localhost:8000/custom?name=c2
		name := c.DefaultQuery("name", "c1")
		c.String(http.StatusOK, "hello "+name)
	})

	// 表单传输为POST请求，常见传输格式四种
	// - application/json
	// - application/x-www-form-urlencoded
	// - application/xml
	// - multipart/form-data
	// 通过PostForm()方法获取表单参数，默认解析x-www-form-urlencoded或form-data格式参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	// multipart/form-data格式用于文件上传
	r.POST("/upload", func(c *gin.Context) { // 单文件
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, "上传图片出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.MaxMultipartMemory = 8 << 20                // 8MB, 默认32MB
	r.POST("/uploadFiles", func(c *gin.Context) { // 多文件
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// routes group管理相同URL
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	// JSON数据解析和绑定
	// 客户端传参，后端接收并解析到结构体
	type Login struct {
		User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
		Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	}
	// 要求以JSON格式上传
	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	// 表单数据解析和绑定
	r.POST("loginForm", func(c *gin.Context) {
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	// URI数据解析和绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// 各种数据格式的响应
	// 1. json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3. XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	// 4. YAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "suyb"})
	})
	// 5. protobuf，谷歌开发的高效存储读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	// HTML模版渲染
	// 加载HTML模版并根据模版参数配置并返回相应数据
	// LoadHTMLGlob()方法加载模版文件
	// c.HTML()方法响应
	// 引入静态文件需要定义一个静态文件目录, r.Static("/assets", "./assets")

	// 重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 同步异步
	// goroutine实现
	// 1. 异步，需要创建一个副本来执行
	r.GET("/long_async", func(c *gin.Context) {
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行: " + copyContext.Request.URL.Path)
		}()
	})
	// 2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行: " + c.Request.URL.Path)
	})

	// 中间件
	// 全局
	// r.Use(MiddleWare())
	{
		// 局部中间件
		r.GET("/ce2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})

		r.GET("/ce", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(200, gin.H{"request": req})
		})
	}
	// 中间件练习
	r.Use(myTime())
	shoppingG := r.Group("/shopping")
	{
		shoppingG.GET("/index", shopIndexHandler)
		shoppingG.GET("/home", shopHomeHandler)
	}

	r.GET("/cookie/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// name string
			// value string
			// maxAge int 单位秒
			// path string cookie所在目录
			// domain string
			// secure bool 是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie: %s\n", cookie)
	})

	r.GET("/cookie/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})
	r.GET("/cookie/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	// 结构体验证
	type Persion struct {
		// 不能空且大于10
		Age      int       `form:"age" binding:"required,gt=10"`
		Name     string    `form:"name" binding:"required"`
		Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	}
	r.GET("/validateStruct", func(c *gin.Context) {
		var p Persion
		if err := c.ShouldBind(&p); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", p))
	})

	// 日志文件
	// gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/xxx", func(c *gin.Context) {

	})
	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "su")
	c.String(http.StatusOK, "hello "+name+"\n")
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "yb")
	c.String(http.StatusOK, "hello "+name+"\n")
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		// 设置变量到Context的key中，通过Get获取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time: ", t2)
	}
}

func myTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 执行函数
		c.Next()
		t2 := time.Since(t)
		fmt.Println("use time: ", t2)
	}
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续函数处理
		c.Abort()
		return
	}
}
