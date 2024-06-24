package models

import (
	"time"
)

// 时间戳转换成日期的方法
func UnixToTime(timeStamp int) string {
	t := time.Unix(int64(timeStamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
