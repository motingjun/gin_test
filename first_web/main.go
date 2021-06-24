package main

import (
	// "net/http"
	// "strings"
	// "fmt"

	"github.com/gin-gonic/gin"
	// "first_web/routers"
	// "first_web/app/blog"
	// "first_web/app/shop"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello gin!")
// 	})
// 	r.Run(":8010")
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")

// 		action = strings.Trim(action, "/")
// 		c.String(http.StatusOK, name+" is "+action)
// 	})
// 	r.Run(":8010")
// }


// func main() {
// 	r := gin.Default()

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/login", login)
// 		v1.GET("submit", submit)
// 	}
// 	v2 := r.Group("/v2")
// 	{
// 		v2.POST("/login", login)
// 		v2.POST("/submit", submit)
// 	}
// 	r.Run(":8010")
// }

// func login(c *gin.Context) {
// 	name := c.DefaultQuery("name", "jack")
// 	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
// }

// func submit(c *gin.Context) {
// 	name := c.DefaultQuery("name", "lily")
// 	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
// }

// func main() {
// 	r := routers.SetupRouter()
// 	if err := r.Run(":8010"); err != nil {
// 		fmt.Println("startup service failed, err:%v\n", err)
// 	} 
// }

// func main() {
// 	// 加载多个APP的路由配置
// 	routers.Include(shop.Routers, blog.Routers)
// 	//初始化路由
// 	r := routers.Init()
// 	if err := r.Run(); err != nil {
// 		fmt.Println("startup service failed, err:%v\n", err)
// 	}
// }


// // 定义接收数据的结构体
// type Login struct {
// 	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
// 	User	string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
// 	Pssword	string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
// }

// func main() {
// 	r := gin.Default()
// 	r.POST("loginJSON", func(c *gin.Context) {
// 		//声明接收的变量
// 		var json Login
// 		//将reque的body中的数据，自动按照json格式解析到结构体
// 		if err := c.ShouldBindJSON(&json); err != nil {
// 			// 返回错误信息
// 			// gin.H封装了生成json数据的工具
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		// 判断用户名密码是否正确
// 		if json.User != "root" || json.Pssword != "admin" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status": "304",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"status": "200",})
// 	})
// 	r.Run(":8010")
// }


// 多种响应方式
func main() {
	r := gin.Default()
	// 1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "someJSON",
			"status": "200",
		})
	})
	// 2.结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name 	string
			Message string
			Number 	int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3.XML
	r.GET("someXML", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "abc",
		})
	})
	// 4.YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"name": "zhangsan",
		})
	})
	// 5.protobuf格式，谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(200, data)
	})

	r.Run(":8010")
}