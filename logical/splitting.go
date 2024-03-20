package logical

import (
	"IP_Detection_Set/bin"
	"IP_Detection_Set/db"
	"IP_Detection_Set/mode"
	"fmt"
	"log"
	"strings"
)

func PingJudgment(url string, token string, starttime string, devicetype string) {
	var successDevices []mode.DeviceCheckInfo
	var failedDevices []mode.DeviceCheckInfo
	failedfilename := "tmp/failed" + devicetype + bin.GetCurrentTimeName() + ".csv"
	// 用于存储循环中的信息的字符串
	var content, failedContent strings.Builder
	deviceInfo := db.PullData(url)

	for _, info := range deviceInfo {
		CheackResult := bin.PingCheack(info.IPInner, 3)
		var status string
		if CheackResult == false {
			status = "Failed"
			// 添加到失败的设备切片中
			failedDevices = append(failedDevices, mode.DeviceCheckInfo{
				Ip:     info.IPInner,
				Name:   info.Name,
				Note:   info.AssMemo,
				Status: status,
			})
		} else {
			status = "Success"
			// 添加到成功的设备切片中
			successDevices = append(successDevices, mode.DeviceCheckInfo{
				Ip:     info.IPInner,
				Name:   info.Name,
				Note:   info.AssMemo,
				Status: status,
			})
		}

	}
	// 打印成功的设备信息
	content.WriteString("检测开始时间： " + starttime + "\n检测设备类型：" + devicetype + "\n检测结束时间： " + bin.GetCurrentTime() + "\n \n状态正常设备:\n")
	for _, device := range successDevices {
		SuccessInfo := fmt.Sprintf("类型: %s, IP: %s, 备注: %s\n", device.Name, device.Ip, device.Note)
		//SuccessInfo := fmt.Sprintf("|类型|IP|备注| \n|-------|-----|-----|\n|%s|%s|%s|\n", device.Name, device.Ip, device.Note)
		content.WriteString(SuccessInfo)
	}

	failedContent.WriteString("检测开始时间： " + starttime + "\n检测设备类型：" + devicetype + "\n检测结束时间： " + bin.GetCurrentTime() + "\n状态异常设备:\n")
	// 打印失败的设备信息
	for _, device := range failedDevices {
		FailedInfo := fmt.Sprintf("类型: %s, IP: %s, 备注: %s\n", device.Name, device.Ip, device.Note)
		//FailedInfo := fmt.Sprintf("|类型|IP|备注| \n|-------|-----|-----|\n|%s|%s|%s|\n", device.Name, device.Ip, device.Note)
		failedContent.WriteString(FailedInfo)
	}
	//// 调用函数将成功和失败设备信息写入 CSV 文件
	//err := bin.WriteDeviceDataToCSV(successDevices, "tmp/sucees"+devicetype+bin.GetCurrentTimeName()+".csv")
	//if err != nil {
	//	log.Println("Failed to write success devices to CSV:", err)
	//}

	err := bin.WriteDeviceDataToCSV(failedDevices, failedfilename)
	if err != nil {
		log.Println("Failed to write failed devices to CSV:", err)
	}
	media, err := bin.UploadMedia(token, failedfilename)
	if err != nil {
		return
	}

	err = bin.SendMessage(token, failedContent.String(), 4096)
	//err = bin.SendMessage(token, content.String()+failedContent.String(), 4096)
	if err != nil {
		return
	}
	err = bin.SendFile(token, media)
	if err != nil {
		fmt.Println(err)
	}

}
