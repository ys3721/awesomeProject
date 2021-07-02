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
	content
	data = "{\"msg_type\":\"text\",\"content\":{\"text\":\"" +youNeedSleep+ "\"}}"
)

func main(){
	fmt.Println("Begin main....")
	log.SetPrefix("daily notice service: ")

	c := cron.New()
	c.AddFunc("0 30 * * * *", sendSleepNotice)
	//c.AddFunc("@every 1s", sendSleepNotice)

	c.Start()
	fmt.Println("cron start ok..")
	select {}
}

func sendPunchInNotice() {
	fmt.Println("Begin send punch in notice ..")
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/bbedb9f8-8061-49b4-84e3-f8d8b89e5892"
	var jsonStr = []byte({"1"})
	req := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	client.Do(nil)
}

func sendSleepNotice() {
	fmt.Println("Begin send sleep notice!")
	data := "{\"msg_type\":\"text\",\"content\":{\"text\":\"" +youNeedSleep+ "\"}}"
	b := bytes.Buffer{}
	b.Write([]byte(data))
	http.Post("https://open.feishu.cn/open-apis/bot/v21/hook/bbedb9f8-8061-49b4-84e3-f8d8b89e5892",
		"application/json", &b)
}