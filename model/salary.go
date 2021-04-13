package model

import (
	"gorm.io/gorm"
)

type Salary struct {
	gorm.Model
	SalaryId   string `gorm:"column:salary_id" json:"salary_id"`
	StaffId    string `gorm:"column:staff_id" json:"staff_id"`
	StaffName  string `gorm:"column:staff_name" json:"staff_name"`
	Base       int64  `gorm:"column:base" json:"base"`
	Subsidy    int64  `gorm:"column:subsidy" json:"subsidy"`
	Bonus      int64  `gorm:"column:bonus" json:"bonus"`
	Commission int64  `gorm:"column:commission" json:"commission"`
	Other      int64  `gorm:"column:other" json:"other"`
	Fund       int64  `gorm:"column:fund" json:"fund"`
}

type SalaryCreateDTO struct {
	StaffId    string `gorm:"column:staff_id" json:"staff_id"`
	StaffName  string `gorm:"column:staff_name" json:"staff_name"`
	Base       int64  `gorm:"column:base" json:"base"`
	Subsidy    int64  `gorm:"column:subsidy" json:"subsidy"`
	Bonus      int64  `gorm:"column:bonus" json:"bonus"`
	Commission int64  `gorm:"column:commission" json:"commission"`
	Other      int64  `gorm:"column:other" json:"other"`
	Fund       int64  `gorm:"column:fund" json:"fund"`
}

type SalaryEditDTO struct {
	Id         int64
	StaffId    string `gorm:"column:staff_id" json:"staff_id"`
	StaffName  string `gorm:"column:staff_name" json:"staff_name"`
	Base       int64  `gorm:"column:base" json:"base"`
	Subsidy    int64  `gorm:"column:subsidy" json:"subsidy"`
	Bonus      int64  `gorm:"column:bonus" json:"bonus"`
	Commission int64  `gorm:"column:commission" json:"commission"`
	Other      int64  `gorm:"column:other" json:"other"`
	Fund       int64  `gorm:"column:fund" json:"fund"`
}

type SalaryRecord struct {
	gorm.Model
	SalaryRecordId string `gorm:"column:salary_record_id" json:"salary_record_id"`
	StaffId        string `gorm:"column:staff_id" json:"staff_id"`
	StaffName      string `gorm:"column:staff_name" json:"staff_name"`
	Base           int64  `gorm:"column:base" json:"base"`
	Subsidy        int64  `gorm:"column:subsidy" json:"subsidy"`
	Benifits       int64  `gorm:"column:benifits" json:"benifits"`
	Fine           int64  `gorm:"column:fine" json:"fine"`
	Total          int64  `gorm:"column:total" json:"total"`
	IsPay          int64  `gorm:"column:is_pay" json:"is_pay"`
	SalaryDate     string `gorm:"column:salary_date" json:"salary_date"`
}

type SalaryRecordCreateDTO struct {
	StaffId   string `gorm:"column:staff_id" json:"staff_id"`
	StaffName string `gorm:"column:staff_name" json:"staff_name"`
	Base      int64  `gorm:"column:base" json:"base"`
	Subsidy   int64  `gorm:"column:subsidy" json:"subsidy"`
	Benifits  int64  `gorm:"column:benifits" json:"benifits"`
	Fine      int64  `gorm:"column:fine" json:"fine"`
	//Total      int64  `gorm:"column:total" json:"total"`
	//IsPay      int64  `gorm:"column:is_pay" json:"is_pay"`
	SalaryDate string `gorm:"column:salary_date" json:"salary_date"`
}

type SalaryRecordEditDTO struct {
	Id        int64
	StaffId   string `gorm:"column:staff_id" json:"staff_id"`
	StaffName string `gorm:"column:staff_name" json:"staff_name"`
	Base      int64  `gorm:"column:base" json:"base"`
	Subsidy   int64  `gorm:"column:subsidy" json:"subsidy"`
	Benifits  int64  `gorm:"column:benifits" json:"benifits"`
	Fine      int64  `gorm:"column:fine" json:"fine"`
	//Total      int64  `gorm:"column:total" json:"total"`
	//IsPay      int64  `gorm:"column:is_pay" json:"is_pay"`
	SalaryDate string `gorm:"column:salary_date" json:"salary_date"`
}
