package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
	"net/http"
)

func GetNotificationByTitle(c *gin.Context, noticeTitle string, start int, limit int) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Find(&notifications).Error
		}

	} else {
		// 加分页
		if noticeTitle != "all" {
			err = resource.HrmsDB(c).Where("notice_title like ?", "%"+noticeTitle+"%").Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		} else {
			err = resource.HrmsDB(c).Order("date desc").Offset(start).Limit(limit).Find(&notifications).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Notification{}).Count(&total)
	if noticeTitle != "all" {
		total = int64(len(notifications))
	}
	return notifications, total, nil
}

func CreateNotification(c *gin.Context, dto *model.NotificationDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.NoticeId = RandomID("notice")
	notification.Date = Str2Time(dto.Date, 0)
	// 富文本内容base64编码(前端实现)
	//notification.NoticeContent = base64.StdEncoding.EncodeToString([]byte(dto.NoticeContent))
	if err := resource.HrmsDB(c).Create(&notification).Error; err != nil {
		log.Printf("CreateNotification err = %v", err)
		return err
	}

	// 紧急通知，获取公司员工列表，发放短信
	if notification.Type == "紧急通知" {
		var staffs []*model.Staff
		if err := resource.HrmsDB(c).Find(&staffs).Error; err != nil {
			log.Printf("CreateNotification err = %v", err)
			return err
		}
		// 获取员工手机号，发送紧急通知短信
		for _, staff := range staffs {
			sendNoticeMsg("notice", staff.Phone, notification.NoticeTitle)
		}
	}
	return nil
}

//https://api.apishop.net/communication/sms/send?apiKey=IBIMUBn846955ab1be1d10738e67fdb7214c5fef9a626c6&phoneNum=15521306934&templateID= 10713&params=["紧急通知"]
const SMS_URL = "https://api.apishop.net/communication/sms/send?" +
	"apiKey=IBIMUBn846955ab1be1d10738e67fdb7214c5fef9a626c6&phoneNum=%v&templateID=%v"

func sendNoticeMsg(msgType string, phone int64, noticeTitle string) {
	if phone != 15521306934 {
		return
	}
	var err error
	if msgType == "notice" {
		// 紧急通知
		reqUrl := fmt.Sprint(SMS_URL, phone, 10713) + "&params=[" + noticeTitle + "]"
		log.Printf("[sendNoticeMsg] reqUrl = %v", reqUrl)
		_, err = http.Post(reqUrl, "application/json", nil)
	}
	if msgType == "salary" {
		// 发薪通知
		reqUrl := fmt.Sprint(SMS_URL, phone, 10714) + "&params=[" + noticeTitle + "]"
		log.Printf("[sendNoticeMsg] reqUrl = %v", reqUrl)
		_, err = http.Post(reqUrl, "application/json", nil)
	}
	if err != nil {
		log.Printf("[sendNoticeMsg] err = %v", err)
	}
}

func DelNotificationById(c *gin.Context, notice_id string) error {
	if err := resource.HrmsDB(c).Where("notice_id = ?", notice_id).Delete(&model.Notification{}).Error; err != nil {
		log.Printf("DelNotificationById err = %v", err)
		return err
	}
	return nil
}

func UpdateNotificationById(c *gin.Context, dto *model.NotificationEditDTO) error {
	var notification model.Notification
	Transfer(&dto, &notification)
	notification.Date = Str2Time(dto.Date, 0)
	if err := resource.HrmsDB(c).Where("id = ?", notification.ID).
		Updates(&notification).Error; err != nil {
		log.Printf("UpdateNotificationById err = %v", err)
		return err
	}
	return nil
}
