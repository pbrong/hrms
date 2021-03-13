package model

import (
	"fmt"
	"math/rand"
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

func TestConstructStaffId(t *testing.T) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		randStaffStr := fmt.Sprintf("H%v", rand.Uint32())
		fmt.Println("staffId = " + randStaffStr[0:6])
	}
}

func TestIdenLenSplit(t *testing.T) {
	ident := "460034199905215518"
	identLen := len(ident)
	fmt.Println("pass: " + ident[identLen-6:identLen])
}
