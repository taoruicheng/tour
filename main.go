package main

import (
	"log"

	"github.com/taoruicheng/tour/cmd"
)

func main() {
	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
