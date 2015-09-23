package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("date_converter | ")
}
func main() {
	var seconds = flag.Int64("s", 0, "指定utc秒数")
	var layer = flag.String("fmt", "2006-01-02 15:04:05", "指定日期格式, 如果分别指定输入输出日期，可用\"|\"分隔")
	var date = flag.String("d", "", "指定日期")
	var method = flag.String("m", "conv", "输出方法：\n\t\tconv: 转换成fmt格式\n\t\tdiff: 计算指定时间和当前时间差值")

	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		return
	}

	curTime := time.Now()
	var gotTime time.Time
	var err error
	seg := strings.Split(*layer, "|")
	fmt_in := *layer
	fmt_out := *layer
	if len(seg) == 2 {
		fmt_in = seg[0]
		fmt_out = seg[1]
	}
	if *seconds > 0 {
		gotTime = time.Unix(*seconds, 0)
	} else if *date != "" {
		gotTime, err = time.Parse(fmt_in, *date)
		if err != nil {
			log.Println("parse date error:", err)
			return
		}
	} else {
		flag.Usage()
		return
	}
	switch *method {
	case "conv":
		fmt.Println(gotTime.Format(fmt_out))
	case "diff":
		dur := curTime.Sub(gotTime)
		fmt.Printf("duration = %.2fh = %.2fs\n", dur.Hours(), dur.Seconds())
	default:
		log.Println("不支持输出格式")
	}
}
