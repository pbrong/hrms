package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"log"
)

func DepartCreate(c *gin.Context) {
	var departmentCreateDTO model.DepartmentCreateDTO
	if err := c.BindJSON(&departmentCreateDTO); err != nil {
		log.Printf("[handler.DepartCreate] err = %v", err)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    err.Error(),
		})
		return
	}
	var departmentCheck model.Department
	var result *gorm.DB
	// 找不到记录也会抛出ErrRecordNotFound错误，但这其实不算错误情况
	resource.HrmsDB(c).Where("dep_name = ?", departmentCreateDTO.DepName).First(&departmentCheck)
	if departmentCheck.DepName == departmentCreateDTO.DepName {
		log.Printf("[HrmsDB.Create] 部门已存在, dep = %v", departmentCheck)
		c.JSON(200, gin.H{
			"status": 2001,
			"msg":    "部门名称已存在",
		})
		return
	}
	departmentCreate := model.Department{
		DepId:       service.RandomID("dep"),
		DepDescribe: departmentCreateDTO.DepDescribe,
		DepName:     departmentCreateDTO.DepName,
	}
	if result = resource.HrmsDB(c).Create(&departmentCreate); result.Error != nil {
		result.Rollback()
		log.Printf("[HrmsDB.Create] err = %v", result.Error)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    result.Error.Error(),
		})
		return
	}
	if result = resource.HrmsDB(c).Where("id = ?", departmentCreate.ID); result.Error != nil {
		log.Printf("[HrmsDB.Create] 插入数据失败， departmentCreate = %v", departmentCreate)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    "插入数据失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    departmentCreate,
	})
}

func DepartEdit(c *gin.Context) {
	var departmentEditDTO model.DepartmentEditDTO
	if err := c.BindJSON(&departmentEditDTO); err != nil {
		log.Printf("[DepartEdit] err = %v", err)
		c.JSON(500, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	resource.HrmsDB(c).Model(&model.Department{}).Where("dep_id = ?", departmentEditDTO.DepId).
		Updates(&model.Department{DepDescribe: departmentEditDTO.DepDescribe, DepName: departmentEditDTO.DepName})
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func DepartQuery(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	depId := c.Param("dep_id")
	var deps []model.Department
	if depId == "all" {
		// 查询全部
		if start == -1 && start == -1 {
			resource.HrmsDB(c).Find(&deps)
		} else {
			resource.HrmsDB(c).Offset(start).Limit(limit).Find(&deps)
		}
		if len(deps) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB(c).Model(&model.Department{}).Count(&total)
		c.JSON(200, gin.H{
			"status": code,
			"total":  total,
			"msg":    deps,
		})
		return
	}
	resource.HrmsDB(c).Where("dep_id = ?", depId).Find(&deps)
	if len(deps) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(deps))
	c.JSON(200, gin.H{
		"status": code,
		"total":  total,
		"msg":    deps,
	})
}

func DepartDel(c *gin.Context) {
	depId := c.Param("dep_id")
	if err := resource.HrmsDB(c).Where("dep_id = ?", depId).Delete(&model.Department{}).Error; err != nil {
		log.Printf("[DepartDel] err = %v", err)
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
