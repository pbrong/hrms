package model

import "gorm.io/gorm"

type AttendanceRecord struct {
	gorm.Model
	AttendanceId string `gorm:"column:attendance_id" json:"attendance_id"`
	StaffId      string `gorm:"column:staff_id" json:"staff_id"`
	StaffName    string `gorm:"column:staff_name" json:"staff_name"`
	Date         string `gorm:"column:date" json:"date"`
	WorkDays     int64  `gorm:"column:work_days" json:"work_days"`
	LeaveDays    int64  `gorm:"column:leave_days" json:"leave_days"`
	OvertimeDays int64  `gorm:"column:overtime_days" json:"overtime_days"`
	Approve      int64  `gorm:"column:approve" json:"approve"`
}

type AttendanceRecordCreateDTO struct {
	StaffId      string `gorm:"column:staff_id" json:"staff_id"`
	StaffName    string `gorm:"column:staff_name" json:"staff_name"`
	Date         string `gorm:"column:date" json:"date"`
	WorkDays     int64  `gorm:"column:work_days" json:"work_days"`
	LeaveDays    int64  `gorm:"column:leave_days" json:"leave_days"`
	OvertimeDays int64  `gorm:"column:overtime_days" json:"overtime_days"`
}

type AttendanceRecordEditDTO struct {
	Id           int64
	StaffId      string `gorm:"column:staff_id" json:"staff_id"`
	StaffName    string `gorm:"column:staff_name" json:"staff_name"`
	Date         string `gorm:"column:date" json:"date"`
	WorkDays     int64  `gorm:"column:work_days" json:"work_days"`
	LeaveDays    int64  `gorm:"column:leave_days" json:"leave_days"`
	OvertimeDays int64  `gorm:"column:overtime_days" json:"overtime_days"`
}
