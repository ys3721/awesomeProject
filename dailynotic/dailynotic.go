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
	youNeedSleep    = "中午不锻炼的话就赶紧睡觉吧！"
	urlStr          = "https://open.feishu.cn/open-apis/bot/v2/hook/bbedb9f8-8061-49b4-84e3-f8d8b89e5892"
	mustSleepOrDead = "半夜失眠，受不了，早点睡吧"
	punchin         = "Do not forgot daka!"
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
	c.AddFunc("*/5 * * * * ?", sendNightSleepNotice)
	c.AddFunc("*/5 * * * * ?", sendPunchInNotice)
	c.AddFunc("*/5 * * * * ?", sendMidSleepNotice)

	c.Start()
	fmt.Println("cron start ok..")
	select {}
}

func sendPunchInNotice() {
	fmt.Println("Begin send punch in notice!")
	sendTextNotice(punchin)
}

func sendMidSleepNotice() {
	fmt.Println("Begin send sleep notice!")
	sendTextNotice(youNeedSleep)
}

func sendNightSleepNotice() {
	log.Printf("Begin send night notice!")
	sendTextNotice(mustSleepOrDead)
}

func sendTextNotice(content string) {
	pd, err := buildPostDataStr(content)
	if err != nil {
		log.Printf("Send text notic err %v !", err)
		return
	}
	http.Post(urlStr, "application/json", bytes.NewBuffer(pd))
}
