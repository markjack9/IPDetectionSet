package main

import (
	"IP_Detection_Set/bin"
	"IP_Detection_Set/logical"
	"IP_Detection_Set/mode"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var DeviceMerging []mode.DeviceCheckInfo
	var IpMerging []string
	var failedContentMerging strings.Builder
	qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	starttime := bin.GetCurrentTime()
	NewDeviceInfo, ipsinfo := checkPrinter()
	DeviceMerging = append(DeviceMerging, NewDeviceInfo...)
	IpMerging = append(IpMerging, ipsinfo...)
	if NewDeviceInfo != nil {
		failedContentMerging.WriteString("\n打印机:\n" + strings.Join(IpMerging, ","))
	}
	IpMerging = nil
	time.Sleep(3 * time.Second)
	NewDeviceInfo, ipsinfo = checkAP()
	DeviceMerging = append(DeviceMerging, NewDeviceInfo...)
	IpMerging = append(IpMerging, ipsinfo...)
	if NewDeviceInfo != nil {
		failedContentMerging.WriteString("\nAP:\n" + strings.Join(IpMerging, ","))
	}

	IpMerging = nil
	time.Sleep(3 * time.Second)
	NewDeviceInfo, ipsinfo = checkBroadcast()
	DeviceMerging = append(DeviceMerging, NewDeviceInfo...)
	IpMerging = append(IpMerging, ipsinfo...)
	if NewDeviceInfo != nil {
		failedContentMerging.WriteString("\n广播:\n" + strings.Join(IpMerging, ","))
	}
	IpMerging = nil
	time.Sleep(3 * time.Second)
	NewDeviceInfo, ipsinfo = checkCamera()
	DeviceMerging = append(DeviceMerging, NewDeviceInfo...)
	IpMerging = append(IpMerging, ipsinfo...)
	if NewDeviceInfo != nil {
		failedContentMerging.WriteString("\n摄像头:\n" + strings.Join(IpMerging, ","))
	}
	IpMerging = nil
	time.Sleep(3 * time.Second)
	NewDeviceInfo, ipsinfo = checkAccessControl()
	fmt.Println("所有任务已完成")
	DeviceMerging = append(DeviceMerging, NewDeviceInfo...)
	IpMerging = append(IpMerging, ipsinfo...)
	if NewDeviceInfo != nil {
		failedContentMerging.WriteString("\n门禁:\n" + strings.Join(IpMerging, ","))
	}
	IpMerging = nil
	err := logical.Infomerging(starttime, qyapitoken, DeviceMerging, failedContentMerging.String())
	if err != nil {
		log.Println(err)
	}
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
func checkPrinter() (infomerging []mode.DeviceCheckInfo, ips []string) {
	fmt.Println("正在检测打印机...")
	url := "http://simba.gyyx.cn/specific_assets?atype=32"
	//qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	//currentTime := bin.GetCurrentTime()
	//devicetype := "打印机"
	infomerging, ips = logical.PingJudgment(url)
	return
	// 这里放置打印机检测的具体逻辑
}

// 检测AP任务
func checkAP() (infomerging []mode.DeviceCheckInfo, ips []string) {
	fmt.Println("正在检测AP...")
	url := "http://simba.gyyx.cn/specific_assets?atype=47"
	//qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	//currentTime := bin.GetCurrentTime()
	//devicetype := "AP"
	infomerging, ips = logical.PingJudgment(url)
	return
	// 这里放置AP检测的具体逻辑
}

// 检测广播任务
func checkBroadcast() (infomerging []mode.DeviceCheckInfo, ips []string) {
	fmt.Println("正在检测广播...")
	url := "http://simba.gyyx.cn/specific_assets?atype=50"
	//qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	//currentTime := bin.GetCurrentTime()
	//devicetype := "广播"
	infomerging, ips = logical.PingJudgment(url)
	return
	// 这里放置广播检测的具体逻辑
}

// 检测摄像头任务
func checkCamera() (infomerging []mode.DeviceCheckInfo, ips []string) {
	fmt.Println("正在检测摄像头...")
	url := "http://simba.gyyx.cn/specific_assets?atype=49"
	//qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	//currentTime := bin.GetCurrentTime()
	//devicetype := "摄像头"
	infomerging, ips = logical.PingJudgment(url)
	return
	// 这里放置摄像头检测的具体逻辑
}

// 打印机检测任务
func checkAccessControl() (infomerging []mode.DeviceCheckInfo, ips []string) {
	fmt.Println("正在检测门禁...")
	url := "http://simba.gyyx.cn/specific_assets?atype=51"
	//qyapitoken := "c2ae10dc-c831-4d08-a573-f1515caedc6b"
	//currentTime := bin.GetCurrentTime()
	//devicetype := "门禁"
	infomerging, ips = logical.PingJudgment(url)
	return
	// 模拟打印机检测过程
}
