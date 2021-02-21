package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func AcceptPage(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	if pageStr == "" {
		log.Printf("未传入分页参数page，查询全部")
		return -1, -1
	}
	page, _ := strconv.Atoi(pageStr)
	limitStr := c.Query("limit")
	if limitStr == "" {
		log.Printf("未传入分页参数limit，查询全部")
		return -1, -1
	}
	limit, _ := strconv.Atoi(limitStr)
	startIndex := (page - 1) * limit
	return startIndex, limit
}

func RandomID(pre string) string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%v_%v", pre, rand.Uint32())
}
