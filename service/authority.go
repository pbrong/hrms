package service

import (
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
	resource.HrmsDB(c).Model(&model.AuthorityDetail{}).Where("user_type = ?", userType).Count(&total)
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
