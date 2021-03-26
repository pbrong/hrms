package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
)

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
