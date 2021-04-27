package handler

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"log"
	"net/http"
)

func PasswordQuery(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	staffId := c.Param("staff_id")
	var psws []model.PasswordQueryVO
	result, err := buildPasswordQueryResult(c, staffId, start, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5000,
			"total":  0,
			"msg":    err,
		})
		return
	}
	// 总记录数
	resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Model(&model.Staff{}).Count(&total)
	psws = result
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    psws,
	})
}

func buildPasswordQueryResult(c *gin.Context, staffId string, start int, limit int) ([]model.PasswordQueryVO, error) {
	var loginList []model.Authority
	var err error
	if staffId == "all" {
		// 查询全部
		if start == -1 && limit == -1 {
			// 不加分页
			err = resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Find(&loginList).Error
		} else {
			// 加分页
			err = resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Offset(start).Limit(limit).Find(&loginList).Error
		}
	} else {
		// 查询单个用户
		err = resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Where("staff_id = ?", staffId).First(&loginList).Error
	}
	if err != nil {
		log.Printf("[buildPasswordQueryResult] err = %v", err)
		return nil, err
	}
	var queryVOs []model.PasswordQueryVO
	for _, loginData := range loginList {
		queryVO := model.PasswordQueryVO{
			Id:        int64(loginData.ID),
			StaffId:   loginData.StaffId,
			StaffName: convertStaffIdToName(c, loginData.StaffId),
			Password:  loginData.UserPassword,
		}
		queryVOs = append(queryVOs, queryVO)
	}
	return queryVOs, nil
}

type Result struct {
	StaffName string `json:"staff_name"`
}

func convertStaffIdToName(c *gin.Context, staffId string) string {
	var result Result
	resource.HrmsDB(c).Raw("select staff_name from staff where staff_id = ?", staffId).Scan(&result)
	return result.StaffName
}

func PasswordEdit(c *gin.Context) {
	var passwordEditDTO model.PasswordEditDTO
	if err := c.Bind(&passwordEditDTO); err != nil {
		log.Printf("[PasswordEdit] err = %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5000,
			"msg":    err,
		})
		return
	}
	staffId := passwordEditDTO.StaffId
	password := service.MD5(passwordEditDTO.Password)
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Updates(&model.Authority{
		UserPassword: password,
	}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 5000,
			"msg":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
	})
}
