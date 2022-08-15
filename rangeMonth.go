package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(rangeDay([]int{2022, 11, 3}, []int{2023, 11, 3}))
}

// coding coding
// 给定开始日期，结束日期，按照自然月拆分，按时间顺序每行输出一个自然月的开始日期和结束日期。
func rangeDay(beg []int, end []int) []map[string]int {
	yearDay := make([]map[string]int, 0) //存放结果集
	begYear := beg[0]
	endYear := end[0]
	begMonth := beg[1]
	endMonth := end[1]

	if begYear == endYear {
		//1、开始年份等于结束年份
		oneYearDay := getYearBegEndDay(begYear, begMonth, endMonth)
		yearDay = append(yearDay, oneYearDay...)
	}

	for i := 0; i <= endYear-begYear; i++ { //year循环
		//开始结束时间年份一样，年份不一样   开始年份需要判断月份生成月    中间年份所有月   结束年份考虑结束年份的月
		curYear := begYear + i

		//2不等于
		//if begYear < endYear {
		if curYear == begYear { //考虑起始月份
			oneYearDay := getYearBegEndDay(curYear, begMonth, 12)
			yearDay = append(yearDay, oneYearDay...)
		} else if curYear == endYear { //考虑结束年月份
			oneYearDay := getYearBegEndDay(curYear, 1, endMonth)
			yearDay = append(yearDay, oneYearDay...)
		} else {
			oneYearDay := getYearBegEndDay(curYear, 1, 12)
			yearDay = append(yearDay, oneYearDay...)
		}

		//}
	}
	return yearDay

}

func getYearBegEndDay(year int, begMonth int, endMonth int) []map[string]int {
	oneYearDay := make([]map[string]int, 0)
	for j := 0; j <= endMonth-begMonth; j++ {
		oneBeg := make(map[string]int)
		oneBeg["year"] = year
		oneBeg["month"] = begMonth + j
		oneBeg["day"] = 1
		oneEnd := make(map[string]int)
		oneEnd["year"] = year
		oneEnd["month"] = oneBeg["month"]
		//nextMonth := oneEnd["month"] + 1
		oneEnd["day"] = time.Date(year, time.Month(oneEnd["month"]), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, -1).Day()
		//拿到该年月的开始结束时间
		oneYearDay = append(oneYearDay, oneBeg)
		oneYearDay = append(oneYearDay, oneEnd)
	}
	return oneYearDay
}
