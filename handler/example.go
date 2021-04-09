package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
	"net/http"
	"strconv"
)

func ParseExampleContent(c *gin.Context) {
	// 业务处理
	content, err := service.ParseExampleContent(c)
	if err != nil {
		log.Printf("[ParseExampleContent] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5000,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    content,
	})
}

func CreateExample(c *gin.Context) {
	// 参数绑定
	var dto model.ExampleCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateExample] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.CreateExample(c, &dto)
	if err != nil {
		log.Printf("[CreateExample] err = %v", err)
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

func UpdateExampleById(c *gin.Context) {
	// 参数绑定
	var dto model.ExampleEditDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateExampleById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateExampleById(c, &dto)
	if err != nil {
		log.Printf("[UpdateExampleById] err = %v", err)
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

func DelExample(c *gin.Context) {
	// 参数绑定
	exampleId := c.Param("example_id")
	// 业务处理
	err := service.DelExampleByExampleId(c, exampleId)
	if err != nil {
		log.Printf("[DelExample] err = %v", err)
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

func GetExampleByName(c *gin.Context) {
	// 参数绑定
	name := c.Param("name")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetExampleByName(c, name, start, limit)
	if err != nil {
		log.Printf("[GetExampleByName] err = %v", err)
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

func RenderExample(c *gin.Context) {
	// 参数绑定
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	result, err := service.RenderExample(c, int64(id))
	if err != nil {
		log.Printf("[RenderExample] err = %v", err)
		c.Redirect(http.StatusInternalServerError, "login.html")
	}
	c.HTML(http.StatusOK, "example_doing.html", result)
}

func CreateExampleScore(c *gin.Context) {
	// 参数绑定
	var dto model.ExampleScoreCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateExampleScore] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"msg":    err.Error(),
		})
		return
	}
	// 业务处理
	total, err := service.CreateExampleScore(c, &dto)
	if err != nil {
		log.Printf("[CreateExampleScore] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5002,
			"msg":    err.Error(),
			"total":  0,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"total":  total,
	})
}

func GetExampleHistoryByName(c *gin.Context) {
	// 参数绑定
	name := c.Param("name")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetExampleHistoryByName(c, name, start, limit)
	if err != nil {
		log.Printf("[GetExampleHistoryByName] err = %v", err)
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

func GetExampleHistoryByStafId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetExampleHistoryByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetExampleHistoryByStafId] err = %v", err)
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
