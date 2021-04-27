package service

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
)

//func CreateSalaryRecord(c *gin.Context, dto *model.SalaryRecordCreateDTO) error {
//	var total int64
//	resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("staff_id = ? and salary_date = ?", dto.StaffId, dto.SalaryDate).Count(&total)
//	if total != 0 {
//		return errors.New(fmt.Sprintf("该员工薪资数据已经存在"))
//	}
//	var salaryRecord model.SalaryRecord
//	Transfer(&dto, &salaryRecord)
//	salaryRecord.SalaryRecordId = RandomID("salary_record")
//	salaryRecord.Total = salaryRecord.Base + salaryRecord.Subsidy + salaryRecord.Benifits - salaryRecord.Fine
//	salaryRecord.IsPay = 1 // 1未发放 2发放
//	if err := resource.HrmsDB(c).Create(&salaryRecord).Error; err != nil {
//		log.Printf("CreateSalaryRecord err = %v", err)
//		return err
//	}
//	return nil
//}

func DelSalaryRecordBySalaryRecordId(c *gin.Context, salaryRecordId string) error {
	if err := resource.HrmsDB(c).Where("salary_record_id = ?", salaryRecordId).Delete(&model.SalaryRecord{}).
		Error; err != nil {
		log.Printf("DelSalaryRecordBySalaryRecordId err = %v", err)
		return err
	}
	return nil
}

//func UpdateSalaryRecordById(c *gin.Context, dto *model.SalaryRecordEditDTO) error {
//	var salaryRecord model.SalaryRecord
//	Transfer(&dto, &salaryRecord)
//	salaryRecord.Total = salaryRecord.Base + salaryRecord.Subsidy + salaryRecord.Benifits - salaryRecord.Fine
//	if err := resource.HrmsDB(c).Model(&model.SalaryRecord{}).Where("id = ?", salaryRecord.ID).
//		Update("staff_id", salaryRecord.StaffId).
//		Update("staff_name", salaryRecord.StaffName).
//		Update("base", salaryRecord.Base).
//		Update("subsidy", salaryRecord.Subsidy).
//		Update("benifits", salaryRecord.Benifits).
//		Update("fine", salaryRecord.Fine).
//		Update("salary_date", salaryRecord.SalaryDate).
//		Error; err != nil {
//		log.Printf("UpdateSalaryById err = %v", err)
//		return err
//	}
//	return nil
//}

func GetSalaryRecordByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.SalaryRecord, int64, error) {
	var salaryRecords []*model.SalaryRecord
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Order("salary_date desc").Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Order("salary_date desc").Find(&salaryRecords).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Order("salary_date desc").Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Order("salary_date desc").Find(&salaryRecords).Error
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
	var salarys []*model.SalaryRecord
	resource.HrmsDB(c).Where("id = ?", id).Find(&salarys)
	if len(salarys) > 0 && salarys[0].IsPay == 2 {
		// 发送短信通知员工薪资已发放
		date := salarys[0].SalaryDate
		sendNoticeMsg("salary", getStaffPhoneByStaffId(c, salarys[0].StaffId), []string{date})
	}
	// 通过税务系统上报税款并将工资发放至员工银行卡中
	payStaffSalaryAndTax(salarys[0])
	return nil
}

func payStaffSalaryAndTax(record *model.SalaryRecord) {
	// 需要对接银行系统及税务系统
}

func getStaffPhoneByStaffId(c *gin.Context, staffId string) int64 {
	var staffs []*model.Staff
	resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&staffs)
	if len(staffs) == 0 {
		return 15521306934
	}
	return staffs[0].Phone
}

func GetHadPaySalaryRecordByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.SalaryRecord, int64, error) {
	var salaryRecords []*model.SalaryRecord
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ? and is_pay = 2", staffId).Order("salary_date desc").Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Where("is_pay = 2").Order("salary_date desc").Find(&salaryRecords).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ? and is_pay = 2", staffId).Offset(start).Limit(limit).Order("salary_date desc").Find(&salaryRecords).Error
		} else {
			err = resource.HrmsDB(c).Where("is_pay = 2").Offset(start).Limit(limit).Order("salary_date desc").Find(&salaryRecords).Error
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
