package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
)

func AddAuthorityDetail(c *gin.Context) {
	var authorityDetailDTO model.AddAuthorityDetailDTO
	if err := c.ShouldBindJSON(&authorityDetailDTO); err != nil {
		log.Printf("[AddAuthorityDetail] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	err := service.AddAuthorityDetail(&authorityDetailDTO)
	if err != nil {
		log.Printf("[AddAuthorityDetail] err = %v", err)
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

func GetAuthorityDetailByUserTypeAndModel(c *gin.Context) {
	var dto model.GetAuthorityDetailDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[GetAuthorityDetailByUserTypeAndModel] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	content, err := service.GetAuthorityDetailByUserTypeAndModel(&dto)
	if err != nil {
		log.Printf("[GetAuthorityDetailByUserTypeAndModel] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    content,
	})
}

func GetAuthorityDetailListByUserType(c *gin.Context) {
	// 分页
	start, limit := service.AcceptPage(c)
	userType := c.Param("user_type")
	detailList, total, err := service.GetAuthorityDetailListByUserType(userType, start, limit)
	if err != nil {
		log.Printf("[GetAuthorityDetailByUserTypeAndModel] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"total":  total,
		"msg":    detailList,
	})
}

func UpdateAuthorityDetailById(c *gin.Context) {
	// 参数绑定
	var dto model.UpdateAuthorityDetailDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateAuthorityDetailById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateAuthorityDetailById(&dto)
	if err != nil {
		log.Printf("[UpdateAuthorityDetailById] err = %v", err)
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
