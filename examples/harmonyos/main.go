package main

import (
	"fmt"

	"github.com/DyncMing/jpush-api-golang-client"
)

func main() {
	// HarmonyOS (HMOS) example
	var pf jpush.Platform
	pf.Add(jpush.HMOS)

	var at jpush.Audience
	at.All()

	var n jpush.Notification
	n.SetAlert("Universal alert")
	n.SetHmos(&jpush.HmosNotification{
		Alert:  "Hello HMOS",
		Title:  "HMOS Title",
		Intent: map[string]interface{}{"url": "scheme://test?key=1"},
	})

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
