package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
)

func Template(c *gin.Context) {
	// 参数绑定
	var authorityDetailDTO model.AddAuthorityDetailDTO
	if err := c.ShouldBindJSON(&authorityDetailDTO); err != nil {
		log.Printf("[Template] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.AddAuthorityDetail(c, &authorityDetailDTO)
	if err != nil {
		log.Printf("[Template] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}
