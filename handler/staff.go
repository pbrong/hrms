package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"log"
	"net/http"
)

func StaffCreate(c *gin.Context) {
	var staffCreateDto model.StaffCreateDTO
	if err := c.BindJSON(&staffCreateDto); err != nil {
		log.Printf("[StaffCreate] err = %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err.Error(),
		})
		return
	}
	log.Printf("[StaffCreate staff = %v]", staffCreateDto)
	staffId := service.RandomStaffId()
	// 创建员工信息
	staff := model.Staff{
		StaffId:     staffId,
		StaffName:   staffCreateDto.StaffName,
		Birthday:    service.Str2Time(staffCreateDto.BirthdayStr, 0),
		IdentityNum: staffCreateDto.IdentityNum,
		Sex:         service.SexStr2Int64(staffCreateDto.SexStr),
		Nation:      staffCreateDto.Nation,
		School:      staffCreateDto.School,
		Major:       staffCreateDto.Major,
		EduLevel:    staffCreateDto.EduLevel,
		BaseSalary:  staffCreateDto.BaseSalary,
		CardNum:     staffCreateDto.CardNum,
		RankId:      staffCreateDto.RankId,
		DepId:       staffCreateDto.DepId,
		Email:       staffCreateDto.Email,
		EntryDate:   service.Str2Time(staffCreateDto.EntryDateStr, 0),
	}
	var exist int64
	resource.HrmsDB.Model(&model.Staff{}).Where("identity_num = ? or staff_id = ?", staffCreateDto.IdentityNum, staffId).Count(&exist)
	if exist != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"msg":    "已经存在",
		})
		return
	}
	// 创建登陆信息，密码为身份证后六位
	identLen := len(staff.IdentityNum)
	login := model.Login{
		LoginId:      service.RandomID("pass"),
		StaffId:      staffId,
		UserPassword: staff.IdentityNum[identLen-6 : identLen],
		Aval:         1,
		UserType:     "normal", // 暂时只能创建普通员工
	}
	err := resource.HrmsDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&staff).Error; err != nil {
			return err
		}
		if err := tx.Create(&login).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("[StaffCreate err = %v]", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
		"msg":    staff,
	})
}

func StaffEdit(c *gin.Context) {
	var staffEditDTO model.StaffEditDTO
	if err := c.BindJSON(&staffEditDTO); err != nil {
		log.Printf("[StaffEdit] err = %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	log.Printf("[StaffEdit staff = %v]", staffEditDTO)
	resource.HrmsDB.Model(&model.Staff{}).Where("staff_id = ?", staffEditDTO.StaffId).
		Updates(&model.Staff{
			StaffId:     staffEditDTO.StaffId,
			StaffName:   staffEditDTO.StaffName,
			Birthday:    service.Str2Time(staffEditDTO.BirthdayStr, 0),
			IdentityNum: staffEditDTO.IdentityNum,
			Sex:         service.SexStr2Int64(staffEditDTO.SexStr),
			Nation:      staffEditDTO.Nation,
			School:      staffEditDTO.School,
			Major:       staffEditDTO.Major,
			EduLevel:    staffEditDTO.EduLevel,
			BaseSalary:  staffEditDTO.BaseSalary,
			CardNum:     staffEditDTO.CardNum,
			RankId:      staffEditDTO.RankId,
			DepId:       staffEditDTO.DepId,
			Email:       staffEditDTO.Email,
			EntryDate:   service.Str2Time(staffEditDTO.EntryDateStr, 0),
		})
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func StaffQuery(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	staffId := c.Param("staff_id")
	var staffs []model.Staff
	if staffId == "all" {
		// 查询全部
		if start == -1 && start == -1 {
			resource.HrmsDB.Find(&staffs)
		} else {
			resource.HrmsDB.Offset(start).Limit(limit).Find(&staffs)
		}
		if len(staffs) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB.Model(&model.Staff{}).Count(&total)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"total":  total,
			"msg":    convert2VO(staffs),
		})
		return
	}
	resource.HrmsDB.Where("staff_id = ?", staffId).Find(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(staffs),
	})
}

func convert2VO(staffs []model.Staff) []model.StaffVO {
	var staffVOs []model.StaffVO
	for _, staff := range staffs {
		staffVOs = append(staffVOs, model.StaffVO{
			Staff:    staff,
			DepName:  service.GetDepNameByDepId(staff.DepId),
			RankName: service.GetRankNameRankDepId(staff.RankId),
		})
	}
	return staffVOs
}

func StaffQueryByName(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	staffName := c.Param("staff_name")
	var staffs []model.Staff
	if staffName == "all" {
		// 查询全部
		if start == -1 && start == -1 {
			resource.HrmsDB.Find(&staffs)
		} else {
			resource.HrmsDB.Offset(start).Limit(limit).Find(&staffs)
		}
		if len(staffs) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB.Model(&model.Staff{}).Count(&total)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"total":  total,
			"msg":    convert2VO(staffs),
		})
		return
	}
	resource.HrmsDB.Where("staff_name like ?", "%"+staffName+"%").Find(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(staffs),
	})
}

func StaffDel(c *gin.Context) {
	rankId := c.Param("staff_id")
	if err := resource.HrmsDB.Where("staff_id = ?", rankId).Delete(&model.Staff{}).Error; err != nil {
		log.Printf("[StaffDel] err = %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
	})
}
