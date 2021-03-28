package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
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

func TestExampleExcelParse(t *testing.T) {
	xfile, err := xlsx.OpenFile("/Users/arong/MyFile/毕业设计/试题测试.xlsx")
	if err != nil {
		panic(err)
	}
	var items []*ExampleItem
	for _, sheet := range xfile.Sheets {
		// 遍历行读取
		for number, row := range sheet.Rows[1:] {
			item := &ExampleItem{Num: number + 1}
			// 遍历每行的列读取
			for index, cell := range row.Cells {
				cur := cell.String()
				if index == 0 {
					item.Title = cur
				} else if index > 0 && index < len(row.Cells)-1 {
					item.Items = append(item.Items, cur)
				} else {
					item.Ans = cur
				}
			}
			items = append(items, item)
		}
		bytes, _ := json.Marshal(&items)
		content := string(bytes)
		fmt.Println(content)
		var itemResps []ExampleItem
		_ = json.Unmarshal([]byte(content), &itemResps)
		fmt.Println(itemResps)
	}
}

func TestBase(t *testing.T) {
	name := "彭博荣"
	toString := base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println(toString)
}
