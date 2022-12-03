package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"hrms/model"
	"hrms/resource"
	"hrms/service"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
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
	// 创建员工信息落表
	if staff, err := buildStaffInfoSaveDB(c, staffCreateDto); err != nil {
		log.Printf("[StaffCreate err = %v]", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 2000,
			"msg":    staff,
		})
	}
}

func buildStaffInfoSaveDB(c *gin.Context, staffCreateDto model.StaffCreateDTO) (model.Staff, error) {
	staffID := service.RandomStaffId()
	staff := model.Staff{
		StaffId:       staffID,
		StaffName:     staffCreateDto.StaffName,
		LeaderStaffId: staffCreateDto.LeaderStaffId,
		Phone:         staffCreateDto.Phone,
		Birthday:      service.Str2Time(staffCreateDto.BirthdayStr, 0),
		IdentityNum:   staffCreateDto.IdentityNum,
		Sex:           service.SexStr2Int64(staffCreateDto.SexStr),
		Nation:        staffCreateDto.Nation,
		School:        staffCreateDto.School,
		Major:         staffCreateDto.Major,
		EduLevel:      staffCreateDto.EduLevel,
		BaseSalary:    staffCreateDto.BaseSalary,
		CardNum:       staffCreateDto.CardNum,
		RankId:        staffCreateDto.RankId,
		DepId:         staffCreateDto.DepId,
		Email:         staffCreateDto.Email,
		EntryDate:     service.Str2Time(staffCreateDto.EntryDateStr, 0),
	}
	var exist int64
	resource.HrmsDB(c).Model(&model.Staff{}).Where("identity_num = ? or staff_id = ?", staffCreateDto.IdentityNum, staffID).Count(&exist)
	if exist != 0 {
		return staff, errors.New("已经存在该员工")
	}
	// 查询leader名称
	var leader model.Staff
	resource.HrmsDB(c).Where("staff_id = ?", staffCreateDto.LeaderStaffId).Find(&leader)
	staff.LeaderName = leader.StaffName
	// 创建登陆信息，密码为身份证后六位
	identLen := len(staff.IdentityNum)
	login := model.Authority{
		AuthorityId:  service.RandomID("auth"),
		StaffId:      staffID,
		UserPassword: service.MD5(staff.IdentityNum[identLen-6 : identLen]),
		//Aval:         1,
		UserType: "normal", // 暂时只能创建普通员工
	}
	err := resource.HrmsDB(c).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&staff).Error; err != nil {
			return err
		}
		if err := tx.Create(&login).Error; err != nil {
			return err
		}
		return nil
	})

	return staff, err
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
	staff := model.Staff{
		StaffId:       staffEditDTO.StaffId,
		StaffName:     staffEditDTO.StaffName,
		LeaderStaffId: staffEditDTO.LeaderStaffId,
		Phone:         staffEditDTO.Phone,
		Birthday:      service.Str2Time(staffEditDTO.BirthdayStr, 0),
		IdentityNum:   staffEditDTO.IdentityNum,
		Sex:           service.SexStr2Int64(staffEditDTO.SexStr),
		Nation:        staffEditDTO.Nation,
		School:        staffEditDTO.School,
		Major:         staffEditDTO.Major,
		EduLevel:      staffEditDTO.EduLevel,
		BaseSalary:    staffEditDTO.BaseSalary,
		CardNum:       staffEditDTO.CardNum,
		RankId:        staffEditDTO.RankId,
		DepId:         staffEditDTO.DepId,
		Email:         staffEditDTO.Email,
		EntryDate:     service.Str2Time(staffEditDTO.EntryDateStr, 0),
	}
	// 查询leader名称
	var leader model.Staff
	resource.HrmsDB(c).Where("staff_id = ?", staffEditDTO.LeaderStaffId).Find(&leader)
	staff.LeaderName = leader.StaffName
	resource.HrmsDB(c).Model(&model.Staff{}).Where("staff_id = ?", staffEditDTO.StaffId).
		Updates(&staff)
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
			resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Find(&staffs)
		} else {
			resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Offset(start).Limit(limit).Find(&staffs)
		}
		if len(staffs) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB(c).Model(&model.Staff{}).Where("staff_id != 'root' and staff_id != 'admin'").Count(&total)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"total":  total,
			"msg":    convert2VO(c, staffs),
		})
		return
	}
	resource.HrmsDB(c).Where("staff_id = ? and staff_id != 'root' and staff_id != 'admin'", staffId).Find(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(c, staffs),
	})
}

func getRuleByStaffId(c *gin.Context, staffId string) string {
	var authority model.Authority
	var userTypeName string
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&authority).Error; err == nil {
		switch authority.UserType {
		case "supersys":
			userTypeName = "超级管理员"
		case "sys":
			userTypeName = "系统管理员"
		case "normal":
			userTypeName = "普通员工"
		default:
			userTypeName = "未知"
		}
	}
	return userTypeName
}
func convert2VO(c *gin.Context, staffs []model.Staff) []model.StaffVO {
	var staffVOs []model.StaffVO
	for _, staff := range staffs {
		staffVOs = append(staffVOs, model.StaffVO{
			Staff:        staff,
			DepName:      service.GetDepNameByDepId(c, staff.DepId),
			RankName:     service.GetRankNameRankDepId(c, staff.RankId),
			UserTypeName: getRuleByStaffId(c, staff.StaffId),
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
			resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Find(&staffs)
		} else {
			resource.HrmsDB(c).Where("staff_id != 'root' and staff_id != 'admin'").Offset(start).Limit(limit).Find(&staffs)
		}
		if len(staffs) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		resource.HrmsDB(c).Model(&model.Staff{}).Where("staff_id != 'root' and staff_id != 'admin'").Count(&total)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"total":  total,
			"msg":    convert2VO(c, staffs),
		})
		return
	}
	resource.HrmsDB(c).Where("staff_name like ?", "%"+staffName+"%").Where("staff_id != 'root' and staff_id != 'admin'").Find(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(c, staffs),
	})
}

func StaffQueryByDep(c *gin.Context) {
	var total int64 = 1
	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	depName := c.Param("dep_name")
	var staffs []model.Staff
	reqSql := `select * from staff as staff left join department as dep on staff.dep_id = dep.dep_id where staff.deleted_at is null and dep.dep_name like "%v"`
	if start != -1 && limit != -1 {
		reqSql += fmt.Sprintf(` limit %v,%v`, start, limit)
	}
	reqSql = fmt.Sprintf(reqSql, "%"+depName+"%")
	resource.HrmsDB(c).Raw(reqSql).Scan(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(c, staffs),
	})
}

func StaffDel(c *gin.Context) {
	staffId := c.Param("staff_id")
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Delete(&model.Staff{}).Error; err != nil {
		log.Printf("[StaffDel] err = %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 5001,
			"msg":    err,
		})
		return
	}
	// 密码删除
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Delete(&model.Authority{}).Error; err != nil {
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

func StaffQueryByStaffId(c *gin.Context) {
	var total int64 = 1
	code := 2000
	staffId := c.Param("staff_id")
	var staffs []model.Staff
	resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&staffs)
	if len(staffs) == 0 {
		// 不存在
		code = 2001
	}
	total = int64(len(staffs))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    convert2VO(c, staffs),
	})
}

func ExcelExport(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ExcelExport err = %v]", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 5001,
				"msg":    err.Error(),
			})
			return
		}
	}()
	file, err := c.FormFile("excel_staffs")
	if err != nil {
		log.Printf("ExcelExport err = %v", err)
		return
	}
	if strings.Split(file.Filename, ".")[1] != "xlsx" {
		log.Printf("ExcelExport 只可上传xlsx格式文件")
		return
	}
	fileOpen, err := file.Open()
	if err != nil {
		log.Printf("ExcelExport err = %v", err)
		return
	}
	defer fileOpen.Close()
	bytes, err := ioutil.ReadAll(fileOpen)
	if err != nil {
		log.Printf("ExcelExport err = %v", err)
		return
	}
	xfile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		log.Printf("ExcelExport err = %v", err)
		return
	}
	var exportStaffList []model.StaffCreateDTO
	for _, sheet := range xfile.Sheets {
		headers := sheet.Rows[0]
		for _, r := range sheet.Rows[1:] {
			staff := model.StaffCreateDTO{}
			for i, v := range r.Cells {
				switch headers.Cells[i].String() {
				case "员工姓名":
					staff.StaffName = v.String()
				case "指定上级":
					staff.LeaderName = v.String()
				case "上级工号":
					staff.LeaderStaffId = v.String()
				case "员工性别":
					staff.SexStr = v.String()
				case "身份证号":
					staff.IdentityNum = v.String()
				case "出生日期":
					staff.BirthdayStr = v.String()
				case "民族":
					staff.Nation = v.String()
				case "毕业院校":
					staff.School = v.String()
				case "毕业专业":
					staff.Major = v.String()
				case "最高学历":
					staff.EduLevel = v.String()
				case "基本薪资":
					if s, err := v.Int64(); err != nil {
						staff.BaseSalary = -1
					} else {
						staff.BaseSalary = s
					}
				case "银行卡号":
					staff.CardNum = v.String()
				case "职位":
					staff.RankId = getRankID(c, v.String())
				case "部门":
					staff.DepId = getDepID(c, v.String())
				case "电子邮箱":
					staff.Email = v.String()
				case "手机号":
					if s, err := v.Int64(); err != nil {
						staff.Phone = -1
					} else {
						staff.Phone = s
					}
				case "入职日期":
					staff.EntryDateStr = v.String()
				}
			}
			exportStaffList = append(exportStaffList, staff)
		}
	}

	var (
		eg         errgroup.Group
		successNum int64
		errNum     int64
	)
	for _, s := range exportStaffList {
		var s = s
		eg.Go(func() error {
			if _, err := buildStaffInfoSaveDB(c, s); err != nil {
				atomic.AddInt64(&errNum, 1)
				return err
			}
			atomic.AddInt64(&successNum, 1)
			return nil
		})
	}
	eg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
		"msg":    fmt.Sprintf("完成员工信息导入，成功%v条，失败%v条", successNum, errNum),
	})
}

func getDepID(c *gin.Context, depName string) string {
	var dep model.Department
	if err := resource.HrmsDB(c).Model(&model.Department{}).Where("dep_name = ?", depName).Take(&dep).Error; err != nil {
		return "-1"
	}
	return dep.DepId
}

func getRankID(c *gin.Context, rankName string) string {
	var rank model.Rank
	if err := resource.HrmsDB(c).Model(&model.Rank{}).Where("rank_name = ?", rankName).Take(&rank).Error; err != nil {
		return "-1"
	}
	return rank.RankId
}
