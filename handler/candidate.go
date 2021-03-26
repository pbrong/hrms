package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
	"strconv"
)

func CreateCandidate(c *gin.Context) {
	// 参数绑定
	var dto model.CandidateCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateCandidate] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.CreateCandidate(c, &dto)
	if err != nil {
		log.Printf("[CreateCandidate] err = %v", err)
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

func DelCandidateByCandidateId(c *gin.Context) {
	// 参数绑定
	candidateId := c.Param("candidate_id")
	// 业务处理
	err := service.DelCandidateByCandidateId(c, candidateId)
	if err != nil {
		log.Printf("[DelCandidateByCandidateId] err = %v", err)
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

func UpdateCandidateById(c *gin.Context) {
	// 参数绑定
	var dto model.CandidateEditDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateCandidateById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateCandidateById(c, &dto)
	if err != nil {
		log.Printf("[UpdateCandidateById] err = %v", err)
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

func GetCandidateByStaffId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetCandidateByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetCandidateByStaffId] err = %v", err)
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

func GetCandidateByName(c *gin.Context) {
	// 参数绑定
	name := c.Param("name")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetCandidateByName(c, name, start, limit)
	if err != nil {
		log.Printf("[GetCandidateByName] err = %v", err)
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

func SetCandidateRejectById(c *gin.Context) {
	// 参数绑定
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	// 业务处理
	err = service.SetCandidateRejectById(c, int64(id))
	if err != nil {
		log.Printf("[SetCandidateRejectById] err = %v", err)
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

func SetCandidateAcceptById(c *gin.Context) {
	// 参数绑定
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	// 业务处理
	err = service.SetCandidateAcceptById(c, int64(id))
	if err != nil {
		log.Printf("[SetCandidateAcceptById] err = %v", err)
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
