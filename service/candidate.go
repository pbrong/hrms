package service

import (
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
)

func CreateCandidate(c *gin.Context, dto *model.CandidateCreateDTO) error {
	var candidateRecord model.Candidate
	Transfer(&dto, &candidateRecord)
	candidateRecord.CandidateId = RandomID("candidate")
	if err := resource.HrmsDB(c).Create(&candidateRecord).Error; err != nil {
		log.Printf("CreateCandidate err = %v", err)
		return err
	}
	return nil
}

func DelCandidateByCandidateId(c *gin.Context, candidateId string) error {
	if err := resource.HrmsDB(c).Where("candidate_id = ?", candidateId).Delete(&model.Candidate{}).
		Error; err != nil {
		log.Printf("DelCandidateByCandidateId err = %v", err)
		return err
	}
	return nil
}

func UpdateCandidateById(c *gin.Context, dto *model.CandidateEditDTO) error {
	var candidate model.Candidate
	Transfer(&dto, &candidate)
	if err := resource.HrmsDB(c).Model(&model.Candidate{}).Where("id = ?", candidate.ID).
		Updates(&candidate).Error; err != nil {
		log.Printf("UpdateCandidateById err = %v", err)
		return err
	}
	return nil
}

func GetCandidateByName(c *gin.Context, name string, start int, limit int) ([]*model.Candidate, int64, error) {
	var records []*model.Candidate
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if name != "all" {
			err = resource.HrmsDB(c).Where("name like ?", "%"+name+"%").Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Find(&records).Error
		}

	} else {
		// 加分页
		if name != "all" {
			err = resource.HrmsDB(c).Where("name like ?", "%"+name+"%").Offset(start).Limit(limit).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Find(&records).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Candidate{}).Count(&total)
	if name != "all" {
		total = int64(len(records))
	}
	return records, total, nil
}

func GetCandidateByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.Candidate, int64, error) {
	var records []*model.Candidate
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Find(&records).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = resource.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&records).Error
		} else {
			err = resource.HrmsDB(c).Offset(start).Limit(limit).Find(&records).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	resource.HrmsDB(c).Model(&model.Candidate{}).Count(&total)
	if staffId != "all" {
		total = int64(len(records))
	}
	return records, total, nil
}

// 0面试中、1拒绝、2录取

// 拒绝
func SetCandidateRejectById(c *gin.Context, id int64) error {
	if err := resource.HrmsDB(c).Where("id = ?", id).
		Updates(&model.Candidate{Status: 1}).Error; err != nil {
		log.Printf("SetCandidateRejectById err = %v", err)
		return err
	}
	return nil
}

// 录取
func SetCandidateAcceptById(c *gin.Context, id int64) error {
	if err := resource.HrmsDB(c).Where("id = ?", id).
		Updates(&model.Candidate{Status: 2}).Error; err != nil {
		log.Printf("SetCandidateAcceptById err = %v", err)
		return err
	}
	return nil
}
