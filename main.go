package main

import (
	"IP_Detection_Set/bin"
	"IP_Detection_Set/logical"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	checkPrinter()
	time.Sleep(3 * time.Second)
	checkAP()
	time.Sleep(3 * time.Second)
	checkBroadcast()
	time.Sleep(3 * time.Second)
	checkCamera()
	time.Sleep(3 * time.Second)
	checkAccessControl()
	fmt.Println("所有任务已完成")

	// 创建日志文件
	file, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// 设置日志输出到文件
	log.SetOutput(file)

	// 修改标准输出和标准错误输出
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[INFO] ")

	// 记录日志
	log.Println("This is a log message.")
}

// 检测打印机任务
func checkPrinter() {
	fmt.Println("正在检测打印机...")
	url := "http://simba.gyyx.cn/specific_assets?atype=32"
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	currentTime := bin.GetCurrentTime()
	devicetype := "打印机"
	logical.PingJudgment(url, qyapitoken, currentTime, devicetype)
	// 这里放置打印机检测的具体逻辑
}

// 检测AP任务
func checkAP() {
	fmt.Println("正在检测AP...")
	url := "http://simba.gyyx.cn/specific_assets?atype=47"
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	currentTime := bin.GetCurrentTime()
	devicetype := "AP"
	logical.PingJudgment(url, qyapitoken, currentTime, devicetype)
	// 这里放置AP检测的具体逻辑
}

// 检测广播任务
func checkBroadcast() {
	fmt.Println("正在检测广播...")
	url := "http://simba.gyyx.cn/specific_assets?atype=50"
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	currentTime := bin.GetCurrentTime()
	devicetype := "广播"
	logical.PingJudgment(url, qyapitoken, currentTime, devicetype)
	// 这里放置广播检测的具体逻辑
}

// 检测摄像头任务
func checkCamera() {
	fmt.Println("正在检测摄像头...")
	url := "http://simba.gyyx.cn/specific_assets?atype=49"
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	currentTime := bin.GetCurrentTime()
	devicetype := "摄像头"
	logical.PingJudgment(url, qyapitoken, currentTime, devicetype)
	// 这里放置摄像头检测的具体逻辑
}

// 打印机检测任务
func checkAccessControl() {
	fmt.Println("正在检测门禁...")
	url := "http://simba.gyyx.cn/specific_assets?atype=51"
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	currentTime := bin.GetCurrentTime()
	devicetype := "门禁"
	logical.PingJudgment(url, qyapitoken, currentTime, devicetype)
	// 模拟打印机检测过程
}
