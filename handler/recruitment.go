package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
)

func CreateRecruitment(c *gin.Context) {
	// 参数绑定
	var dto model.RecruitmentCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateRecruitment] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.CreateRecruitment(c, &dto)
	if err != nil {
		log.Printf("[CreateRecruitment] err = %v", err)
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

func UpdateRecruitmentById(c *gin.Context) {
	// 参数绑定
	var dto model.RecruitmentEditDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateRecruitmentById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateRecruitmentById(c, &dto)
	if err != nil {
		log.Printf("[UpdateRecruitmentById] err = %v", err)
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

func GetRecruitmentByJobName(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("job_name")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetRecruitmentByJobName(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetRecruitmentByJobName] err = %v", err)
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

func DelRecruitmentByRecruitmentId(c *gin.Context) {
	// 参数绑定
	recruitmentId := c.Param("recruitment_id")
	// 业务处理
	err := service.DelRecruitmentByRecruitmentId(c, recruitmentId)
	if err != nil {
		log.Printf("[DelRecruitmentByRecruitmentId] err = %v", err)
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
