package util

import (
	"log"
	"time"
)


func TransTime(s1 string, s2 string) (f float64, err error) {
	timeFmt := "2006-01-02 15:04:05" //这个时间很特殊，
	parse1, err := time.Parse(timeFmt, s1)
	if err != nil {
		log.Println("时间转换失败",err)
		return 0, err
	}
	if len(s2) > 1{
		parse2, err := time.Parse(timeFmt, s2)
		if err != nil {
			log.Println("时间转换失败",err)
			return 0, err
		}
		d := parse2.Sub(parse1)
		f = d.Seconds()
		return f,nil
	}
	return 0,err
}

func GetTime()string{
	timeFmt := "2006-01-02 15:04:05" //这个
	timeUnix:=time.Now().Unix()//已知的时间戳

	formatTimeStr:=time.Unix(timeUnix,0).Format(timeFmt)

	return formatTimeStr
}
