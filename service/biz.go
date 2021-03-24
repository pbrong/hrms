package service

import (
	"hrms/model"
	"hrms/resource"
	"log"
)

func AddAuthorityDetail(dto *model.AddAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB.Create(&detail).Error; err != nil {
		log.Printf("AddAuthorityDetail err = %v", err)
		return err
	}
	return nil
}

func UpdateAuthorityDetailById(dto *model.UpdateAuthorityDetailDTO) error {
	var detail model.AuthorityDetail
	Transfer(&dto, &detail)
	if err := resource.HrmsDB.Where("id = ?", detail.ID).
		Updates(&detail).Error; err != nil {
		log.Printf("UpdateAuthorityDetailById err = %v", err)
		return err
	}
	return nil
}

func GetAuthorityDetailByUserTypeAndModel(detail *model.GetAuthorityDetailDTO) (string, error) {
	var authorityDetail model.AuthorityDetail
	if err := resource.HrmsDB.Where("user_type = ? and model = ?", detail.UserType, detail.Model).
		Find(&authorityDetail).Error; err != nil {
		log.Printf("GetAuthorityDetailByUserTypeAndModel err = %v", err)
		return "", err
	}
	return authorityDetail.AuthorityContent, nil
}

func GetAuthorityDetailListByUserType(userType string, start int, limit int) ([]*model.AuthorityDetail, int64, error) {
	var authorityDetailList []*model.AuthorityDetail
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		err = resource.HrmsDB.Where("user_type = ?", userType).Find(&authorityDetailList).Error
	} else {
		// 加分页
		err = resource.HrmsDB.Where("user_type = ?", userType).Offset(start).Limit(limit).Find(&authorityDetailList).Error
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB.Model(&model.AuthorityDetail{}).Count(&total)
	return authorityDetailList, total, nil
}

func SetAdminByStaffId(staffId string) error {
	authority := model.Authority{
		UserType: "sys",
	}
	if err := resource.HrmsDB.Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetAdminByStaffId err = %v", err)
		return err
	}
	return nil
}

func SetNormalByStaffId(staffId string) error {
	authority := model.Authority{
		UserType: "normal",
	}
	if err := resource.HrmsDB.Where("staff_id = ?", staffId).Updates(&authority).Error; err != nil {
		log.Printf("SetNormalByStaffId err = %v", err)
		return err
	}
	return nil
}

func GetNotificationByTitle(noticeTitle string, start int, limit int) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB.Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Find(&notifications).Error
		} else {
			err = resource.HrmsDB.Order("date desc").Find(&notifications).Error
		}

	} else {
		// 加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB.Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		} else {
			err = resource.HrmsDB.Order("date desc").Find(&notifications).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB.Model(&model.Notification{}).Count(&total)
	return notifications, total, nil
}

func CreateNotification(dto *model.NotificationDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.NoticeId = RandomID("notice")
	notification.Date = Str2Time(dto.Date, 0)
	// 富文本内容base64编码(前端实现)
	//notification.NoticeContent = base64.StdEncoding.EncodeToString([]byte(dto.NoticeContent))
	if err := resource.HrmsDB.Create(&notification).Error; err != nil {
		log.Printf("CreateNotification err = %v", err)
		return err
	}
	return nil
}

func DelNotificationById(notice_id string) error {
	if err := resource.HrmsDB.Where("notice_id = ?", notice_id).Delete(&model.Notification{}).Error; err != nil {
		log.Printf("DelNotificationById err = %v", err)
		return err
	}
	return nil
}

func UpdateNotificationById(dto *model.NotificationEditDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.Date = Str2Time(dto.Date, 0)
	if err := resource.HrmsDB.Where("id = ?", notification.ID).
		Updates(&notification).Error; err != nil {
		log.Printf("UpdateNotificationById err = %v", err)
		return err
	}
	return nil
}
