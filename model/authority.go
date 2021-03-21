package model

type AuthorityDetail struct {
	ID               int64  `gorm:"column:id" json:"id"`
	UserType         string `gorm:"column:user_type" json:"user_type"`
	Model            string `gorm:"column:model" json:"model"`
	Name             string `gorm:"column:name" json:"name"`
	AuthorityContent string `gorm:"column:authority_content" json:"authority_content"`
}

type AddAuthorityDetailDTO struct {
	UserType         string `gorm:"column:user_type" json:"user_type" binding:"required"`
	Model            string `gorm:"column:model" json:"model" binding:"required"`
	Name             string `gorm:"column:name" json:"name" binding:"required"`
	AuthorityContent string `gorm:"column:authority_content" json:"authority_content"binding:"required"`
}

type GetAuthorityDetailDTO struct {
	UserType string `gorm:"column:user_type" json:"user_type" binding:"required"`
	Model    string `gorm:"column:model" json:"model" binding:"required"`
}

type UpdateAuthorityDetailDTO struct {
	ID               int64  `gorm:"column:id" json:"id" binding:"required"`
	AuthorityContent string `gorm:"column:authority_content" json:"authority_content" binding:"required"`
}
