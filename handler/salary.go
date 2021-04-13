package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/service"
	"log"
	"strconv"
)

func DelSalary(c *gin.Context) {
	// 参数绑定
	salaryId := c.Param("salary_id")
	// 业务处理
	err := service.DelSalaryBySalaryId(c, salaryId)
	if err != nil {
		log.Printf("[DelSalary] err = %v", err)
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

func CreateSalary(c *gin.Context) {
	// 参数绑定
	var dto model.SalaryCreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[CreateSalary] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.CreateSalary(c, &dto)
	if err != nil {
		log.Printf("[CreateSalary] err = %v", err)
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

func UpdateSalaryById(c *gin.Context) {
	// 参数绑定
	var dto model.SalaryEditDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[UpdateSalaryById] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	// 业务处理
	err := service.UpdateSalaryById(c, &dto)
	if err != nil {
		log.Printf("[UpdateSalaryById] err = %v", err)
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

func GetSalaryByStaffId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetSalaryByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetSalaryByStaffId] err = %v", err)
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

//func DelSalaryRecord(c *gin.Context) {
//	// 参数绑定
//	salaryId := c.Param("salary_record_id")
//	// 业务处理
//	err := service.DelSalaryRecordBySalaryRecordId(c, salaryId)
//	if err != nil {
//		log.Printf("[DelSalaryRecord] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5002,
//			"result": err.Error(),
//		})
//		return
//	}
//	c.JSON(200, gin.H{
//		"status": 2000,
//	})
//}
//
//func CreateSalaryRecord(c *gin.Context) {
//	// 参数绑定
//	var dto model.SalaryRecordCreateDTO
//	if err := c.ShouldBindJSON(&dto); err != nil {
//		log.Printf("[CreateSalaryRecord] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5001,
//			"result": err.Error(),
//		})
//		return
//	}
//	// 业务处理
//	err := service.CreateSalaryRecord(c, &dto)
//	if err != nil {
//		log.Printf("[CreateSalaryRecord] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5002,
//			"result": err.Error(),
//		})
//		return
//	}
//	c.JSON(200, gin.H{
//		"status": 2000,
//	})
//}
//
//func UpdateSalaryRecordById(c *gin.Context) {
//	// 参数绑定
//	var dto model.SalaryRecordEditDTO
//	if err := c.ShouldBindJSON(&dto); err != nil {
//		log.Printf("[UpdateSalaryRecordById] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5001,
//			"result": err.Error(),
//		})
//		return
//	}
//	// 业务处理
//	err := service.UpdateSalaryRecordById(c, &dto)
//	if err != nil {
//		log.Printf("[UpdateSalaryRecordById] err = %v", err)
//		c.JSON(200, gin.H{
//			"status": 5002,
//			"result": err.Error(),
//		})
//		return
//	}
//	c.JSON(200, gin.H{
//		"status": 2000,
//	})
//}

func GetSalaryRecordByStaffId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetSalaryRecordByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetSalaryRecordByStaffId] err = %v", err)
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

func GetSalaryRecordIsPayById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 5000,
			"msg":    err,
		})
		return
	}
	isPay := service.GetSalaryRecordIsPayById(c, int64(id))
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    isPay,
	})
}

func PaySalaryRecordById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	err = service.PaySalaryRecordById(c, int64(id))
	if err != nil {
		c.JSON(200, gin.H{
			"status": 5002,
			"msg":    err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func GetHadPaySalaryRecordByStaffId(c *gin.Context) {
	// 参数绑定
	staffId := c.Param("staff_id")
	start, limit := service.AcceptPage(c)
	// 业务处理
	list, total, err := service.GetHadPaySalaryRecordByStaffId(c, staffId, start, limit)
	if err != nil {
		log.Printf("[GetHadPaySalaryRecordByStaffId] err = %v", err)
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
