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
	err := service.AddAuthorityDetail(c, &authorityDetailDTO)
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
	content, err := service.GetAuthorityDetailByUserTypeAndModel(c, &dto)
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
	detailList, total, err := service.GetAuthorityDetailListByUserType(c, userType, start, limit)
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
	err := service.UpdateAuthorityDetailById(c, &dto)
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

func SetAdminByStaffId(c *gin.Context) {
	staffId := c.Param("staff_id")
	if staffId == "" {
		log.Printf("[SetAdminByStaffId] staff_id is empty")
		c.JSON(200, gin.H{
			"status": 5001,
			"result": "staff_id is empty",
		})
		return
	}
	if err := service.SetAdminByStaffId(c, staffId); err != nil {
		log.Printf("[SetAdminByStaffId] err = %v", err)
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

func SetNormalByStaffId(c *gin.Context) {
	staffId := c.Param("staff_id")
	if staffId == "" {
		log.Printf("[SetNormalByStaffId] staff_id is empty")
		c.JSON(200, gin.H{
			"status": 5001,
			"result": "staff_id is empty",
		})
		return
	}
	if err := service.SetNormalByStaffId(c, staffId); err != nil {
		log.Printf("[SetNormalByStaffId] err = %v", err)
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
