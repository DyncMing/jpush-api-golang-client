package main

import (
	"fmt"

	"github.com/DyncMing/jpush-api-golang-client"
)

func main() {
	// QuickApp example
	var pf jpush.Platform
	pf.Add(jpush.QUICKAPP)

	var at jpush.Audience
	at.All()

	var n jpush.Notification
	n.SetAlert("Universal alert")
	n.SetQuickApp(&jpush.QuickAppNotification{Alert: "Hello QuickApp", Title: "QA Title", Page: "/page1"})

	payload := jpush.NewPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&at)
	payload.SetNotification(&n)

	c := jpush.NewJPushClient("appKey", "masterSecret")
	data, err := payload.Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
	res, err := c.Push(data)
	if err != nil {
		fmt.Printf("error: %+v\n", err)
	} else {
		fmt.Printf("ok: %v\n", res)
	}
}
