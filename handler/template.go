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

func GetTemplate(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetSalaryByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[Template] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5000,
			"total":  0,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"total":  total,
		"msg":    list,
	})
}
