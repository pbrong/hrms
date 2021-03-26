package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
)

func AddAuthorityDetail(c *gin.Context, dto *model.AddAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB(c).Create(&detail).Error; err != nil {
		log.Printf("AddAuthorityDetail err = %v", err)
		return err
	}
	return nil
}

func UpdateAuthorityDetailById(c *gin.Context, dto *model.UpdateAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB(c).Where("id = ?", detail.ID).
		Updates(&detail).Error; err != nil {
		log.Printf("UpdateAuthorityDetailById err = %v", err)
		return err
	}
	return nil
}

func GetAuthorityDetailByUserTypeAndModel(c *gin.Context, detail *model.GetAuthorityDetailDTO) (string, error) {
	var authorityDetail model.AuthorityDetail
	if err := resource.HrmsDB(c).Where("user_type = ? and model = ?", detail.UserType, detail.Model).
		Find(&authorityDetail).Error; err != nil {
		log.Printf("GetAuthorityDetailByUserTypeAndModel err = %v", err)
		return "", err
	}
	return authorityDetail.AuthorityContent, nil
}

func GetAuthorityDetailListByUserType(c *gin.Context, userType string, start int, limit int) ([]*model.AuthorityDetail, int64, error) {
	var authorityDetailList []*model.AuthorityDetail
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		err = resource.HrmsDB(c).Where("user_type = ?", userType).Find(&authorityDetailList).Error
	} else {
		// 加分页
		err = resource.HrmsDB(c).Where("user_type = ?", userType).Offset(start).Limit(limit).Find(&authorityDetailList).Error
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.AuthorityDetail{}).Count(&total)
	return authorityDetailList, total, nil
}

func SetAdminByStaffId(c *gin.Context, staffId string) error {
	authority := model.Authority{
		UserType: "sys",
	}
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetAdminByStaffId err = %v", err)
		return err
	}
	return nil
}

func SetNormalByStaffId(c *gin.Context, staffId string) error {
	authority := model.Authority{
		UserType: "normal",
	}
	if err := resource.HrmsDB(c).Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetNormalByStaffId err = %v", err)
		return err
	}
	return nil
}

func GetNotificationByTitle(c *gin.Context, noticeTitle string, start int, limit int) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Find(&notifications).Error
		}

	} else {
		// 加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Notification{}).Count(&total)
	if noticeTitle != "all" {
		total = int64(len(notifications))
	}
	return notifications, total, nil
}

func CreateNotification(c *gin.Context, dto *model.NotificationDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.NoticeId = RandomID("notice")
	notification.Date = Str2Time(dto.Date, 0)
	// 富文本内容base64编码(前端实现)
	//notification.NoticeContent = base64.StdEncoding.EncodeToString([]byte(dto.NoticeContent))
	if err := resource.HrmsDB(c).Create(&notification).Error; err != nil {
		log.Printf("CreateNotification err = %v", err)
		return err
	}
	return nil
}

func DelNotificationById(c *gin.Context, notice_id string) error {
	if err := resource.HrmsDB(c).Where("notice_id = ?", notice_id).Delete(&model.Notification{}).Error; err != nil {
		log.Printf("DelNotificationById err = %v", err)
		return err
	}
	return nil
}

func UpdateNotificationById(c *gin.Context, dto *model.NotificationEditDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.Date = Str2Time(dto.Date, 0)
	if err := resource.HrmsDB(c).Where("id = ?", notification.ID).
		Updates(&notification).Error; err != nil {
		log.Printf("UpdateNotificationById err = %v", err)
		return err
	}
	return nil
}

func CreateSalary(c *gin.Context, dto *model.SalaryCreateDTO) error {
	var total int64
	resource.HrmsDB(c).Model(&model.Salary{}).Where("staff_id = ?", dto.StaffId).Count(&total)
	if total != 0 {
		return errors.New(fmt.Sprintf("该员工薪资数据已经存在"))
	}
	var salary model.Salary
	Transfer(&dto, &salary)
	salary.SalaryId = RandomID("salary")
	if err := resource.HrmsDB(c).Create(&salary).Error; err != nil {
		log.Printf("CreateSalary err = %v", err)
		return err
	}
	return nil
}

func DelSalaryBySalaryId(c *gin.Context, salaryId string) error {
	if err := resource.HrmsDB(c).Where("salary_id = ?", salaryId).Delete(&model.Salary{}).
		Error; err != nil {
		log.Printf("DelSalaryBySalaryId err = %v", err)
		return err
	}
	return nil
}

func UpdateSalaryById(c *gin.Context, dto *model.SalaryEditDTO) error {
	var salary model.Salary
	Transfer(&dto, &salary)
	if err := resource.HrmsDB(c).Model(&model.Salary{}).Where("id = ?", salary.ID).
		Update("staff_id", salary.StaffId).
		Update("staff_name", salary.StaffName).
		Update("base", salary.Base).
		Update("subsidy", salary.Subsidy).
		Error; err != nil {
		log.Printf("UpdateSalaryById err = %v", err)
		return err
	}
	return nil
}

func GetSalaryByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.Salary, int64, error) {
	var salarys []*model.Salary
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&salarys).Error
		} else {
			err = resource.HrmsDB(c).Find(&salarys).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&salarys).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Find(&salarys).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Salary{}).Count(&total)
	if staffId != "all" {
		total = int64(len(salarys))
	}
	return salarys, total, nil
}

func CreateSalaryRecord(c *gin.Context, dto *model.SalaryRecordCreateDTO) error {
	var total int64
	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("staff_id = ? and salary_date = ?", dto.StaffId, dto.SalaryDate).Count(&total)
	if total != 0 {
		return errors.New(fmt.Sprintf("该员工薪资数据已经存在"))
	}
	var salaryRecord model.SalaryRecord
	Transfer(&dto, &salaryRecord)
	salaryRecord.SalaryRecordId = RandomID("salary_record")
	salaryRecord.Total = salaryRecord.Base + salaryRecord.Subsidy + salaryRecord.Benifits - salaryRecord.Fine
	salaryRecord.IsPay = 1 // 1未发放 2发放
	if err := resource.HrmsDB(c).Create(&salaryRecord).Error; err != nil {
		log.Printf("CreateSalaryRecord err = %v", err)
		return err
	}
	return nil
}

func DelSalaryRecordBySalaryRecordId(c *gin.Context, salaryRecordId string) error {
	if err := resource.HrmsDB(c).Where("salary_record_id = ?", salaryRecordId).Delete(&model.SalaryRecord{}).
		Error; err != nil {
		log.Printf("DelSalaryRecordBySalaryRecordId err = %v", err)
		return err
	}
	return nil
}

func UpdateSalaryRecordById(c *gin.Context, dto *model.SalaryRecordEditDTO) error {
	var salaryRecord model.SalaryRecord
	Transfer(&dto, &salaryRecord)
	salaryRecord.Total = salaryRecord.Base + salaryRecord.Subsidy + salaryRecord.Benifits - salaryRecord.Fine
	if err := resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("id = ?", salaryRecord.ID).
		Update("staff_id", salaryRecord.StaffId).
		Update("staff_name", salaryRecord.StaffName).
		Update("base", salaryRecord.Base).
		Update("subsidy", salaryRecord.Subsidy).
		Update("benifits", salaryRecord.Benifits).
		Update("fine", salaryRecord.Fine).
		Update("salary_date", salaryRecord.SalaryDate).
		Error; err != nil {
		log.Printf("UpdateSalaryById err = %v", err)
		return err
	}
	return nil
}

func GetSalaryRecordByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.SalaryRecord, int64, error) {
	var salaryRecords []*model.SalaryRecord
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Find(&salaryRecords).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Find(&salaryRecords).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Count(&total)
	if staffId != "all" {
		total = int64(len(salaryRecords))
	}
	return salaryRecords, total, nil
}

// 如果支付过则返回true
func GetSalaryRecordIsPayById(c *gin.Context, id int64) bool {
	var total int64
	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("id = ? and is_pay = 2", id).Count(&total)
	return total != 0
}

func PaySalaryRecordById(c *gin.Context, id int64) error {
	if err := resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("id = ?", id).
		Update("is_pay", 2).Error; err != nil {
		log.Printf("PaySalaryRecordById err = %v", err)
		return err
	}
	return nil
}

func GetHadPaySalaryRecordByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.SalaryRecord, int64, error) {
	var salaryRecords []*model.SalaryRecord
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ? and is_pay = 2", staffId).Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Where("is_pay = 2").Find(&salaryRecords).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ? and is_pay = 2", staffId).Offset(start).Limit(limit).Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Where("is_pay = 2").Offset(start).Limit(limit).Find(&salaryRecords).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("is_pay = 2").Count(&total)
	if staffId != "all" {
		total = int64(len(salaryRecords))
	}
	return salaryRecords, total, nil
}

func CreateAttendanceRecord(c *gin.Context, dto *model.AttendanceRecordCreateDTO) error {
	var total int64
	resource.HrmsDB(c).Model(&model.AttendanceRecord{}).Where("staff_id = ? and date = ?", dto.StaffId, dto.Date).Count(&total)
	if total != 0 {
		return errors.New(fmt.Sprintf("该月考勤数据已经存在"))
	}
	var attendanceRecord model.AttendanceRecord
	Transfer(&dto, &attendanceRecord)
	attendanceRecord.AttendanceId = RandomID("attendance_record")
	if err := resource.HrmsDB(c).Create(&attendanceRecord).Error; err != nil {
		log.Printf("CreateAttendanceRecord err = %v", err)
		return err
	}
	return nil
}

func DelAttendRecordByAttendId(c *gin.Context, attendanceId string) error {
	if err := resource.HrmsDB(c).Where("attendance_id = ?", attendanceId).Delete(&model.AttendanceRecord{}).
		Error; err != nil {
		log.Printf("DelAttendRecordByAttendId err = %v", err)
		return err
	}
	return nil
}

func UpdateAttendRecordById(c *gin.Context, dto *model.AttendanceRecordEditDTO) error {
	var attentRecord model.AttendanceRecord
	Transfer(&dto, &attentRecord)
	if err := resource.HrmsDB(c).Model(&model.AttendanceRecord{}).Where("id = ?", attentRecord.ID).
		Update("staff_id", attentRecord.StaffId).
		Update("staff_name", attentRecord.StaffName).
		Update("overtime_days", attentRecord.OvertimeDays).
		Update("leave_days", attentRecord.LeaveDays).
		Update("work_days", attentRecord.WorkDays).
		Update("date", attentRecord.Date).
		Error; err != nil {
		log.Printf("UpdateAttendRecordById err = %v", err)
		return err
	}
	return nil
}

func GetAttendRecordByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.AttendanceRecord, int64, error) {
	var records []*model.AttendanceRecord
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Find(&records).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Find(&records).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.AttendanceRecord{}).Count(&total)
	if staffId != "all" {
		total = int64(len(records))
	}
	return records, total, nil
}

func GetAttendRecordHistoryByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.AttendanceRecord, int64, error) {
	var records []*model.AttendanceRecord
	var err error
	sqlReq1 := `select * from attendance_record as attend left join salary_record as salary on attend.staff_id = salary.staff_id
and attend.date = salary.salary_date where salary.is_pay = 2 and attend.staff_id = ?`
	sqlReq2 := `select * from attendance_record as attend left join salary_record as salary on attend.staff_id = salary.staff_id
and attend.date = salary.salary_date where salary.is_pay = 2`
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Raw(sqlReq1, staffId).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Raw(sqlReq2).Find(&records).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Raw(sqlReq1, staffId).Offset(start).Limit(limit).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Raw(sqlReq2).Offset(start).Limit(limit).Find(&records).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.AttendanceRecord{}).Count(&total)
	if staffId != "all" {
		total = int64(len(records))
	}
	return records, total, nil
}

// 如果支付过则返回true
func GetAttendRecordIsPayByStaffIdAndDate(c *gin.Context, staffId string, date string) bool {
	var total int64
	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("staff_id = ? and salary_date = ? and is_pay = 2", staffId, date).Count(&total)
	return total != 0
}
