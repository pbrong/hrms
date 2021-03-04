package model

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateStaff(t *testing.T) {
	timeStr := "2021-01-02"
	//curTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	curTime, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		fmt.Printf("err = %v", err)
		return
	}
	fmt.Printf(curTime.String())
}
