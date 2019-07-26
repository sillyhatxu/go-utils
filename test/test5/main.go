package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Row struct {
	ColumnMap map[string]string
}
type Text struct {
	Key   string
	Value string
}

func ReadFile(file string) []Row {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var result []Row
	var colums = make(map[string]string)
	for _, lineStr := range strings.Split(string(b), "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			result = append(result, Row{ColumnMap: colums})
			colums = make(map[string]string)
			continue
		}
		lineArray := strings.Split(lineStr, ":")
		if len(lineArray) < 2 {
			colums[ReplaceAll(lineArray[0])] = ""
			//result = append(result, *&Text{Key: ReplaceAll(lineArray[0]), Value: ""})
		} else {
			colums[ReplaceAll(lineArray[0])] = lineArray[1]
			//result = append(result, *&Text{Key: ReplaceAll(lineArray[0]), Value: lineArray[1]})
		}
	}
	return result
}

//i, ok := m["route"]
func ReplaceAll(c string) string {
	return strings.Replace(strings.Replace(c, " ", "", -1), "	", "", -1)
}

func main() {
	columns := []string{
		"序号", "姓名", "性别", "年龄",
		"学历", "身份证号", "手机号",
		"邮箱号", "身份证地址", "工作单位",
		"单位地址", "单位电话", "申请原因",
		"银行名称", "收卡地址", "申请时间",
		"推荐人姓名", "推荐人手机号", "备注"}
	rowArray := ReadFile("/Users/cookie/go/gopath/src/go-utils/test/test5/test.txt")
	for i, row := range rowArray {
		fmt.Print(i+1, ",")
		columnMap := row.ColumnMap
		for j, c := range columns {
			if j == 0 {
				continue
			}
			if j == len(columns)-1 {
				fmt.Print(columnMap[c])
			} else {
				fmt.Print(columnMap[c], ",")
			}
		}
		fmt.Println()
	}
}
