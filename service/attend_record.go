package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		Update("approve", 0).
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
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Order("date desc").Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Find(&records).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Order("date desc").Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Order("date desc").Find(&records).Error
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
and attend.date = salary.salary_date where salary.is_pay = 2 and attend.staff_id = ? order by attend.date desc`
	sqlReq2 := `select * from attendance_record as attend left join salary_record as salary on attend.staff_id = salary.staff_id
and attend.date = salary.salary_date where salary.is_pay = 2 order by attend.date desc`
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

// 通过leader_staff_id查询下属提交的考勤上报数据进行审批
func GetAttendRecordApproveByLeaderStaffId(c *gin.Context, leaderStaffId string) ([]*model.AttendanceRecord, int64, error) {
	// 查询下属staff_id
	var staffs []*model.Staff
	resource.HrmsDB(c).Where("leader_staff_id = ?", leaderStaffId).Find(&staffs)
	if len(staffs) == 0 {
		return nil, 0, nil
	}
	// 查询下属是否有未审批的考勤申请
	var err error
	var attends []*model.AttendanceRecord
	for _, staff := range staffs {
		var attend []*model.AttendanceRecord
		staffId := staff.StaffId
		resource.HrmsDB(c).Where("staff_id = ? and approve = 0", staffId).Find(&attend)
		if attend != nil {
			attends = append(attends, attend...)
		}
	}
	if err != nil {
		return nil, 0, err
	}
	total := int64(len(attends))
	return attends, total, nil
}

// 通过考勤审批信息，修改考勤信息为通过，并且按该员工工资套账进行相应的薪资详情计算，得到五险一金税后薪资
func Compute(c *gin.Context, attendId string) error {
	err := resource.HrmsDB(c).Transaction(func(tx *gorm.DB) error {
		// 更新考勤信息为审批通过状态
		if err := tx.Model(&model.AttendanceRecord{}).Where("attendance_id = ?", attendId).Update("approve", 1).Error; err != nil {
			return err
		}
		// 根据考勤信息及该员工薪资套账进行当月薪资计算

		// 获取当月出勤信息
		attendInfo, err := getAttendInfoByAttendId(tx, attendId)
		if err != nil {
			return err
		}
		// 获取员工工号及当月出勤天数、缺勤天数、加班天数及月份
		staffId := attendInfo.StaffId
		workDays := attendInfo.WorkDays
		leaveDays := attendInfo.LeaveDays
		overtimeDays := attendInfo.OvertimeDays
		month := attendInfo.Date
		// 获取该员工薪资套账
		salaryInfo, err := getSalaryInfoByStaffId(tx, staffId)
		if err != nil {
			return err
		}
		// 获员工姓名、基本薪资、住房补贴、绩效奖金、提成薪资、其他薪资、是否缴纳五险一金
		staffName := salaryInfo.StaffName
		base := salaryInfo.Base
		subsidy := salaryInfo.Subsidy
		bonus := salaryInfo.Bonus
		commission := salaryInfo.Commission
		other := salaryInfo.Other
		fund := salaryInfo.Fund
		// 按出勤天数更新基本工资
		base = int64((float64(base) / getCurMonthWorkdays()) * float64(workDays))
		// 更新绩效奖金,每缺勤一天扣1/5
		if leaveDays > 5 {
			bonus = 0
		}
		x := float64(5-leaveDays) / 5.0
		bonus = int64(float64(bonus) * x)
		// 更新加班工资，按国家法定节假日2倍加班费计算
		overtimeSalary := int64((float64(base) / getCurMonthWorkdays()) * 2.0 * float64(overtimeDays))
		// 判断是否交五险一金，不交的话不计算三险
		salaryRecord := model.SalaryRecord{}
		amount := float64(overtimeSalary + base + subsidy + bonus + commission + other)
		if fund == 1 {
			// 缴纳五险一金，计算个人需缴纳养老保险、失业保险和医疗保险及住房公积金
			//养老保险金：   800.00 (8%)   1900.00    (19%)
			//医疗保险金：   200.00 (2%)   1000.00    (10%)
			//失业保险金：   20.00  (0.2%) 80.00  (0.8%)
			//基本住房公积金： 1200  (12%)  1200    (12%)
			//补充住房公积金： 0.00   (0%)   0.00   (0%)
			//工伤保险金：       0         40.00  (0.4%)
			///生育保险金：      0         80.00  (0.8%)
			salaryRecord.PensionInsurance = amount * 0.08
			salaryRecord.MedicalInsurance = amount * 0.02
			salaryRecord.UnemploymentInsurance = amount * 0.002
			salaryRecord.HousingFund = amount * 0.12
		}
		// 计算扣除三险及住房公积金后薪资，并以此计算扣税金额
		amount = amount - salaryRecord.PensionInsurance - salaryRecord.MedicalInsurance -
			salaryRecord.UnemploymentInsurance - salaryRecord.HousingFund
		// 以最新税法，起征点5000元计算，七级扣税
		var tax float64 = 0
		if amount > 5000 {
			total := amount - 5000
			// 按法定税率扣税
			if total <= 3000 {
				tax = total * 0.03
			} else if total <= 12000 {
				tax = total*0.10 - 210
			} else if total <= 25000 {
				tax = total*0.20 - 1410
			} else if total <= 35000 {
				tax = total*0.25 - 2660
			} else if total <= 55000 {
				tax = total*0.30 - 4410
			} else if total <= 80000 {
				tax = total*0.35 - 7160
			} else if total > 80000 {
				tax = total*0.45 - 15160
			}
		}
		// 计算税后工资
		total := amount - tax
		// 创建薪资记录用于发放
		salaryRecord.SalaryRecordId = RandomID("salary_record")
		salaryRecord.StaffId = staffId
		salaryRecord.StaffName = staffName
		salaryRecord.Base = base
		salaryRecord.Subsidy = subsidy
		salaryRecord.Bonus = bonus
		salaryRecord.Commission = commission
		salaryRecord.Overtime = overtimeSalary
		salaryRecord.Other = other
		salaryRecord.Tax = tax
		salaryRecord.Total = total
		salaryRecord.IsPay = 1
		salaryRecord.SalaryDate = month
		// 创建或更新薪资记录
		affected := tx.Where("staff_id = ? and salary_date = ?", staffId, month).Updates(&salaryRecord).RowsAffected
		if affected != 0 {
			// 已更新记录
			return nil
		}
		if err = tx.Create(&salaryRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func getCurMonthWorkdays() float64 {
	return 22.25
}

func getSalaryInfoByStaffId(tx *gorm.DB, staffId string) (*model.Salary, error) {
	var salarys []*model.Salary
	tx.Where("staff_id = ?", staffId).Find(&salarys)
	if len(salarys) == 0 {
		return nil, errors.New("不存在该薪资套账")
	}
	return salarys[0], nil
}

func getAttendInfoByAttendId(tx *gorm.DB, attendId string) (*model.AttendanceRecord, error) {
	var records []*model.AttendanceRecord
	tx.Where("attendance_id = ?", attendId).Find(&records)
	if len(records) == 0 {
		return nil, errors.New("不存在该考勤信息")
	}
	return records[0], nil
}
