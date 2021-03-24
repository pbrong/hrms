package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"log"
)

func RankCreate(c *gin.Context) {
	var rankCreateDto model.RankCreateDTO
	if err := c.BindJSON(&rankCreateDto); err != nil {
		log.Printf("[RankCreate] err = %v", err)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	var exist int64
	resource.HrmsDB(c).Model(&model.Rank{}).Where("rank_name = ?", rankCreateDto.RankName).Count(&exist)
	if exist != 0 {
		c.JSON(200, gin.H{
			"status": 2001,
			"msg":    "职级名称已存在",
		})
		return
	}
	rank := model.Rank{
		RankId:   service.RandomID("rank"),
		RankName: rankCreateDto.RankName,
	}
	resource.HrmsDB(c).Create(&rank)
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    rank,
	})
}

func RankEdit(c *gin.Context) {
	var rankEditDTO model.RankEditDTO
	if err := c.BindJSON(&rankEditDTO); err != nil {
		log.Printf("[RankEdit] err = %v", err)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	resource.HrmsDB(c).Model(&model.Rank{}).Where("rank_id = ?", rankEditDTO.RankId).
		Updates(&model.Rank{RankName: rankEditDTO.RankName})
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func RankQuery(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	rankId := c.Param("rank_id")
	var ranks []model.Rank
	if rankId == "all" {
		// 查询全部
		if start == -1 && start == -1 {
			resource.HrmsDB(c).Find(&ranks)
		} else {
			resource.HrmsDB(c).Offset(start).Limit(limit).Find(&ranks)
		}
		if len(ranks) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB(c).Model(&model.Rank{}).Count(&total)
		c.JSON(200, gin.H{
			"status": code,
			"total":  total,
			"msg":    ranks,
		})
		return
	}
	resource.HrmsDB(c).Where("rank_id = ?", rankId).Find(&ranks)
	if len(ranks) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(ranks))
	c.JSON(200, gin.H{
		"status": code,
		"total":  total,
		"msg":    ranks,
	})
}

func RankDel(c *gin.Context) {
	rankId := c.Param("rank_id")
	if err := resource.HrmsDB(c).Where("rank_id = ?", rankId).Delete(&model.Rank{}).Error; err != nil {
		log.Printf("[RankDel] err = %v", err)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}
