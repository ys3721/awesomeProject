package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
)

const (
	youNeedSleep = "中午不锻炼的话就赶紧睡觉吧！"
	urlStr       = "https://open.feishu.cn/open-apis/bot/v2/hook/bbedb9f8-8061-49b4-84e3-f8d8b89e5892"
)

type MsgInfo struct {
	MsgType string    `json:"msg_type"`
	Content *TextInfo `json:"content"`
}

type TextInfo struct {
	Text string `json:"text"`
}

var testInfo = &TextInfo{Text: youNeedSleep}

var msg = MsgInfo{MsgType: "text", Content: testInfo}

func buildPostDataStr(sendStr string) ([]byte, error) {
	testInfo = &TextInfo{Text: sendStr}
	msgdata := MsgInfo{MsgType: "text", Content: testInfo}
	jsonB, err := json.Marshal(msgdata)
	return jsonB, err
}

func main() {
	fmt.Println("Begin main....")
	log.SetPrefix("daily notice service: ")

	c := cron.New()
	c.AddFunc("0 30 * * * *", sendSleepNotice)
	c.AddFunc("0 0 23 * * *", sendSleepNotice)
	c.AddFunc("0 57 9 * * *", sendPunchInNotice)
	//c.AddFunc("@every 1s", sendSleepNotice)

	c.Start()
	fmt.Println("cron start ok..")
	select {}
}

func sendPunchInNotice() {
	fmt.Println("Begin send punch in notice ..")
	by, err := buildPostDataStr("打卡打打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡打卡卡")
	if err != nil {
		return
	}
	http.Post(urlStr, "application/json", bytes.NewBuffer(by))
}

func sendSleepNotice() {
	fmt.Println("Begin send sleep notice!")
	by, err := buildPostDataStr(youNeedSleep)
	if err != nil {
		return
	}
	ioB := bytes.NewBuffer(by)
	http.Post(urlStr, "application/json", ioB)
}
