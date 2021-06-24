package blog

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	fmt.Println("blog postHandler")
}

func commentHandler(c *gin.Context) {
	fmt.Println("blog commentHandler")
}