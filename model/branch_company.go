package model

type BranchCompany struct {
	ID       int64  `gorm:"column:id" json:"id"`
	BranchId string `gorm:"column:branch_id" json:"branch_id"`
	Name     string `gorm:"column:name" json:"name"`
	Desc     string `gorm:"column:desc" json:"desc"`
}
