package bin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func PingCheack(host string, count int) (hostcheck bool) {
	log.Printf("Pinging %s with %d packets:\n", host, count)
	var failCount, successCount int
	// PingCheck 使用系统命令执行 ICMP Ping 操作
	for i := 1; i <= count; i++ {
		cmd := exec.Command("ping", "-c", "1", host) // 在 Windows 上可能需要使用 "-n" 替代 "-c"
		output, err := cmd.Output()

		if err != nil {
			log.Printf("Ping failed (%d): %v\n", i, err)
			failCount++
		} else {
			// 检查输出中是否包含 "ttl"，如果包含则认为 Ping 成功
			if strings.Contains(string(output), "ttl") {
				log.Printf("Ping succeeded (%d)\n", i)
				successCount++
			} else {
				log.Printf("Ping failed (%d): No response\n", i)
				failCount++
			}
		}
	}
	// 如果成功次数等于循环次数，则判定为成功
	if successCount <= count && failCount >= count {
		hostcheck = false
	} else {
		hostcheck = true
	}

	return hostcheck
}

// SplitMessage 将消息内容按照指定的长度拆分成多个小片段
func SplitMessage(content string, maxBytes int) []string {
	var segments []string
	runes := []rune(content)

	for len(runes) > 0 {
		// 计算当前片段的长度
		var segmentLength int
		for i, r := range runes {
			segmentLength += len(string(r))
			if segmentLength > maxBytes {
				// 当前片段长度超过了最大字节数，截取并保存当前片段
				segments = append(segments, string(runes[:i]))
				runes = runes[i:]
				break
			}
		}
		if segmentLength <= maxBytes {
			// 当前片段长度未超过最大字节数，保存整个内容并退出循环
			segments = append(segments, string(runes))
			break
		}
	}

	return segments
}

// SendMessage 发送消息，超出限制的部分拆分发送
func SendMessage(token, content string, maxBytes int) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", token)

	// 检查消息内容是否超出限制
	if len(content) > maxBytes {
		// 超出限制，拆分消息内容并依次发送
		segments := SplitMessage(content, maxBytes)
		for _, segment := range segments {
			err := sendMessageSegment(url, segment)
			if err != nil {
				return err
			}
		}
	} else {
		// 未超出限制，直接发送消息
		err := sendMessageSegment(url, content)
		if err != nil {
			return err
		}
	}

	return nil
}

// sendMessageSegment 发送消息片段
func sendMessageSegment(url, content string) error {
	// 构造消息结构体
	message := struct {
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("reboot mag err:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
	// 发送消息
	// 这里省略发送消息的逻辑
	// 假设已经发送成功
}

//func SendMessage(token, content string) error {
//	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", token)
//
//	message := mode.TextMessage{
//		MsgType: "text",
//		Text: struct {
//			Content string `json:"content"`
//		}{
//			Content: content,
//		},
//	}
//
//	payload, err := json.Marshal(message)
//	if err != nil {
//		return err
//	}
//
//	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
//	if err != nil {
//		return err
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			log.Println("reboot mag err:", err)
//		}
//	}(resp.Body)
//
//	if resp.StatusCode != http.StatusOK {
//		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
//	}
//
//	return nil
//}

// GetCurrentTime 获取当前时间并返回格式化后的字符串
func GetCurrentTime() string {
	// 获取当前时间
	currentTime := time.Now()

	// 格式化时间为指定格式
	timeFormat := "2006年01月02日15点04分"
	formattedTime := currentTime.Format(timeFormat)

	return formattedTime
}
