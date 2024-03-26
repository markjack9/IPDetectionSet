package logical

import (
	"IP_Detection_Set/bin"
	"IP_Detection_Set/db"
	"IP_Detection_Set/mode"
	"fmt"
	"log"
	"strings"
)

func PingJudgment(url string) (failedDevices []mode.DeviceCheckInfo, IPs []string) {
	//var successDevices []mode.DeviceCheckInfo
	//var failedDevices []mode.DeviceCheckInfo

	//failedfilename := devicetype + bin.GetCurrentTimeName() + ".csv"
	// 用于存储循环中的信息的字符串
	//var content, failedContent strings.Builder
	deviceInfo := db.PullData(url)

	for _, info := range deviceInfo {
		CheackResult := bin.PingCheack(info.IPInner, 3)
		var status string
		if CheackResult != true {
			status = "Failed"
			// 添加到失败的设备切片中
			failedDevices = append(failedDevices, mode.DeviceCheckInfo{
				Ip:     info.IPInner,
				Name:   info.Name,
				Note:   info.AssMemo,
				Status: status,
			})
			IPs = append(IPs, info.IPInner)
			//} else {
			//	status = "Success"
			//	// 添加到成功的设备切片中
			//	successDevices = append(successDevices, mode.DeviceCheckInfo{
			//		Ip:     info.IPInner,
			//		Name:   info.Name,
			//		Note:   info.AssMemo,
			//		Status: status,
			//	})
		}

	}
	//// 打印成功的设备信息
	//content.WriteString("检测开始时间： " + starttime + "\n检测设备类型：" + devicetype + "\n检测结束时间： " + bin.GetCurrentTime() + "\n \n状态正常设备:\n")
	//for _, device := range successDevices {
	//	SuccessInfo := fmt.Sprintf("类型: %s, IP: %s, 备注: %s\n", device.Name, device.Ip, device.Note)
	//	//SuccessInfo := fmt.Sprintf("|类型|IP|备注| \n|-------|-----|-----|\n|%s|%s|%s|\n", device.Name, device.Ip, device.Note)
	//	content.WriteString(SuccessInfo)
	//}
	//
	//failedContent.WriteString("检测开始时间： " + starttime + "\n检测设备类型：" + devicetype + "\n检测结束时间： " + bin.GetCurrentTime() + "\n状态异常设备:\n")
	//// 打印失败的设备信息
	//for _, device := range failedDevices {
	//	FailedInfo := fmt.Sprintf("类型: %s, IP: %s\n", device.Name, device.Ip)
	//	//FailedInfo := fmt.Sprintf("|类型|IP|备注| \n|-------|-----|-----|\n|%s|%s|%s|\n", device.Name, device.Ip, device.Note)
	//	failedContent.WriteString(FailedInfo)
	//}
	return
	//// 调用函数将成功和失败设备信息写入 CSV 文件
	//err := bin.WriteDeviceDataToCSV(successDevices, "tmp/sucees"+devicetype+bin.GetCurrentTimeName()+".csv")
	//if err != nil {
	//	log.Println("Failed to write success devices to CSV:", err)
	//}
	//if failedDevices != nil {
	//err := bin.WriteDeviceDataToCSV(failedDevices, "tmp/"+failedfilename)
	//if err != nil {
	//	log.Println("Failed to write failed devices to CSV:", err)
	//}
	//media, err := bin.UploadMedia(token, failedfilename)
	//if err != nil {
	//	return
	//}
	//
	//err = bin.SendMessage(token, failedContent.String(), 4096)
	////err = bin.SendMessage(token, content.String()+failedContent.String(), 4096)
	//if err != nil {
	//	return
	//}
	//err = bin.SendFile(token, media)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//} else {

	//err := bin.SendMessage(token, "检测开始时间： "+starttime+"\n检测设备类型："+devicetype+"\n检测结束时间： "+bin.GetCurrentTime()+"\n设备检测无异常\n", 4096)
	////err = bin.SendMessage(token, content.String()+failedContent.String(), 4096)
	//if err != nil {
	//	return
	//}
	//}

}

func Infomerging(starttime string, token string, faileinfo []mode.DeviceCheckInfo, longtxet string) error {
	failedfilename := bin.GetCurrentTimeName() + ".csv"
	var failedContent strings.Builder
	failedContent.WriteString("IT\n" + "检测开始时间： " + starttime + "\n检测结束时间： " + bin.GetCurrentTime() + "\n" + "\n状态异常设备:\n" + longtxet)
	//for _, device := range faileinfo {
	//	FailedInfo := fmt.Sprintf("类型: %s, IP: %s\n", device.Name, device.Ip)
	//	//FailedInfo := fmt.Sprintf("|类型|IP|备注| \n|-------|-----|-----|\n|%s|%s|%s|\n", device.Name, device.Ip, device.Note)
	//	failedContent.WriteString(FailedInfo)
	//}
	err := bin.WriteDeviceDataToCSV(faileinfo, "tmp/"+failedfilename)
	if err != nil {
		log.Println("Failed to write failed devices to CSV:", err)
	}
	media, err := bin.UploadMedia(token, failedfilename)
	if err != nil {
		return err
	}

	err = bin.SendMessage(token, failedContent.String(), 4096)
	//err = bin.SendMessage(token, content.String()+failedContent.String(), 4096)
	if err != nil {
		return err
	}
	err = bin.SendFile(token, media)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
