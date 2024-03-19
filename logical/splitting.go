package logical

import (
	"IP_Detection_Set/bin"
	"IP_Detection_Set/db"
	"IP_Detection_Set/mode"
	"fmt"
	"strings"
)

func PingJudgment(url string, token string, starttime string, devicetype string) {
	var successDevices []mode.DeviceCheckInfo
	var failedDevices []mode.DeviceCheckInfo
	// 用于存储循环中的信息的字符串
	var content, failedContent strings.Builder
	deviceInfo := db.PullData(url)
	failedContent.WriteString("\n状态异常设备:\n")
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
		SuccessInfo := fmt.Sprintf("Name: %s, IP: %s, Note: %s\n", device.Name, device.Ip, device.Note)
		content.WriteString(SuccessInfo)
	}

	// 打印失败的设备信息
	for _, device := range failedDevices {
		FailedInfo := fmt.Sprintf("类型: %s, IP: %s, 备注: %s\n", device.Name, device.Ip, device.Note)
		failedContent.WriteString(FailedInfo)
	}

	err := bin.SendMessage(token, content.String()+failedContent.String(), 2048)
	if err != nil {
		return
	}
}
