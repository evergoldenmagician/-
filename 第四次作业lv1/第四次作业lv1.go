package main

import(
	"fmt"
	"time"
)

func main() {
	var a string
	fmt.Scanf("%f",a)
	datetime:=a //待转化为时间戳的字符串

	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64
	fmt.Println(timestamp)

	//时间戳转化为日期
	datetime = time.Unix(timestamp, 0).Format(timeLayout)
	fmt.Println(datetime)
}