package utils

import "log"

// Logging 打印日志
func Logging(info string) {
	if GlobalObject.Log {
		log.SetFlags(log.Ldate | log.Ltime)
		log.Println(info)
	}
}
