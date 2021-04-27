package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	httpReq "github.com/kirinlabs/HttpRequest"
	"github.com/tealeg/xlsx"
	"log"
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

func TestSMS(t *testing.T) {
	reqJSON := map[string]interface{}{
		"apiKey":     "IBIMUBn846955ab1be1d10738e67fdb7214c5fef9a626c6",
		"phoneNum":   15521306934,
		"templateID": "10713",
		"params":     "[\"测试通知\"]",
	}
	datas, _ := json.Marshal(&reqJSON)
	log.Printf("[sendNoticeMsg] req data = %v", string(datas))
	resp, err := httpReq.Post("https://api.apishop.net/communication/sms/send", reqJSON)
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	body, _ := resp.Body()
	log.Printf("[sendNoticeMsg] resp = %v", string(body))
}

func TestComputeSalary(t *testing.T) {
	data := []byte("215517test")
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))
}
