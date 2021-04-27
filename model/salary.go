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
	SalaryRecordId        string  `gorm:"column:salary_record_id" json:"salary_record_id"`
	StaffId               string  `gorm:"column:staff_id" json:"staff_id"`
	StaffName             string  `gorm:"column:staff_name" json:"staff_name"`
	Base                  int64   `gorm:"column:base" json:"base"`
	Subsidy               int64   `gorm:"column:subsidy" json:"subsidy"`
	Bonus                 int64   `gorm:"column:bonus" json:"bonus"`
	Commission            int64   `gorm:"column:commission" json:"commission"`
	Other                 int64   `gorm:"column:other" json:"other"`
	PensionInsurance      float64 `gorm:"column:pension_insurance" json:"pension_insurance"`
	UnemploymentInsurance float64 `gorm:"column:unemployment_insurance" json:"unemployment_insurance"`
	MedicalInsurance      float64 `gorm:"column:medical_insurance" json:"medical_insurance"`
	HousingFund           float64 `gorm:"column:housing_fund" json:"housing_fund"`
	Tax                   float64 `gorm:"column:tax" json:"tax"`
	Overtime              int64   `gorm:"column:overtime" json:"overtime"`
	Total                 float64 `gorm:"column:total" json:"total"`
	IsPay                 int64   `gorm:"column:is_pay" json:"is_pay"`
	SalaryDate            string  `gorm:"column:salary_date" json:"salary_date"`
}
