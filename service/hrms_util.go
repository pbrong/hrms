package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hrms/model"
	"hrms/resource"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func AcceptPage(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	if pageStr == "" {
		log.Printf("未传入分页参数page，查询全部")
		return -1, -1
	}
	page, _ := strconv.Atoi(pageStr)
	limitStr := c.Query("limit")
	if limitStr == "" {
		log.Printf("未传入分页参数limit，查询全部")
		return -1, -1
	}
	limit, _ := strconv.Atoi(limitStr)
	startIndex := (page - 1) * limit
	return startIndex, limit
}

func RandomID(pre string) string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%v_%v", pre, rand.Uint32())
}

func RandomStaffId() string {
	rand.Seed(time.Now().UnixNano())
	randStaffStr := fmt.Sprintf("H%v", rand.Uint32())
	return randStaffStr[0:6]
}

func Str2Time(timeStr string, typ int) time.Time {
	var curTime time.Time
	var err error
	if typ == 0 {
		curTime, err = time.Parse("2006-01-02", timeStr)
		if err != nil {
			fmt.Printf("err = %v", err)
		}
	}
	if typ == 1 {
		curTime, err = time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			fmt.Printf("err = %v", err)
		}
	}
	return curTime
}

func Time2Str(curTime time.Time, typ int) string {
	var timeStr string
	if typ == 0 {
		timeStr = curTime.Format("2006-01-02")
	}
	if typ == 1 {
		timeStr = curTime.Format("2006-01-02 15:04:05")
	}
	return timeStr
}

func SexStr2Int64(sexStr string) int64 {
	var sex int64
	if sexStr == "男" {
		sex = 1
	}
	if sexStr == "女" {
		sex = 2
	}
	return sex
}

func SexInt2Str(sex int64) string {
	var sexStr = "Null"
	if sex == 1 {
		sexStr = "男"
	}
	if sex == 2 {
		sexStr = "女"
	}
	return sexStr
}

func GetDepNameByDepId(depId string) string {
	var dep model.Department
	resource.HrmsDB.Where("dep_id = ?", depId).Find(&dep)
	return dep.DepName
}

func GetRankNameRankDepId(rankId string) string {
	var rank model.Rank
	resource.HrmsDB.Where("rank_id = ?", rankId).Find(&rank)
	return rank.RankName
}

func Transfer(from, to interface{}) error {
	bytes, err := json.Marshal(&from)
	if err != nil {
		log.Println("Transfer json err = %v", err)
		return err
	}
	err = json.Unmarshal(bytes, &to)
	if err != nil {
		log.Println("Transfer json err = %v", err)
		return err
	}
	return nil
}
