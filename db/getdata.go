package db

import (
	"IP_Detection_Set/mode"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PullData(url string) (info []mode.DataDetail) {
	// 发送GET请求

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// 读取返回的JSON数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	// 将JSON数据解码为结构体
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("解析 JSON 数据时出错:", err)
		return
	}

	// 提取 data 字段中的字符串
	dataString, ok := responseData["data"].(string)
	if !ok {
		fmt.Println("无法获取 data 字段")
		return
	}
	// 解析 data 字段中的 JSON 数据
	var data []mode.DataDetail
	err = json.Unmarshal([]byte(dataString), &data)
	if err != nil {
		fmt.Println("解析 JSON 数据时出错:", err)
		return
	}
	info = data
	return info
}
