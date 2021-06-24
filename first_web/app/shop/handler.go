package shop

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func goodsHandler(c *gin.Context) {
	fmt.Println("shop goodsHandler")
}

func checkoutHandler(c *gin.Context) {
	fmt.Println("shop checkoutHandler")
}