# jpush-api-golang-client

> Heads-up: this library is being actively refactored to modernize its APIs and structure. Expect changes to land soon; keep an eye on releases/changelog.

English is the default README. 中文说明请见 [README.zh.md](README.zh.md)。

## Overview
Golang SDK for JPush REST APIs (push, schedule, report, CID, SMS). The current codebase mirrors the official docs with light helpers for HTTP requests and payload builders.

Supported today:
- ✅ Push API v3
- ✅ Report API v3
- ✅ Schedule API v3
- ✅ SMS API v1 (template SMS send)
- ⏳ Not yet: Device API v3, File API v3, Image API v3, Admin API v3

## Install
```bash
go get github.com/DyncMing/jpush-api-golang-client
```

## Quickstart (Push)
1) Platform
```go
var pf jpush.Platform
pf.Add(jpush.ANDROID)
pf.Add(jpush.IOS)
pf.Add(jpush.WINPHONE)
// pf.All()
```

2) Audience
```go
var at jpush.Audience
at.SetTag([]string{"tag1", "tag2"})
at.SetID([]string{"1", "2"})
// at.All()
```

3) Notification or Message
```go
var n jpush.Notification
n.SetAlert("alert")
n.SetAndroid(&jpush.AndroidNotification{Alert: "alert", Title: "title"})
n.SetIos(&jpush.IosNotification{Alert: "alert", Badge: 1})
n.SetWinPhone(&jpush.WinPhoneNotification{Alert: "alert"})

var m jpush.Message
m.MsgContent = "This is a message"
m.Title = "Hello"
```

4) Payload
```go
payload := jpush.NewPayLoad()
payload.SetPlatform(&pf)
payload.SetAudience(&at)
payload.SetNotification(&n)
payload.SetMessage(&m)
```

5) Send
```go
c := jpush.NewJPushClient("appKey", "masterSecret") // obtain from https://www.jiguang.cn/
data, err := payload.Bytes()
if err != nil {
	panic(err)
}
res, err := c.Push(data)
if err != nil {
	fmt.Printf("%+v\n", err)
} else {
	fmt.Printf("ok: %v\n", res)
}
```

See `examples/` for full push and CID samples.

## Third-party channel (厂商通道)

You can set per-vendor delivery options via `Options.ThirdPartyChannel`. The Go constant for HarmonyOS channel is `HMOS_CHANNEL`.

Example:

```go
opts := jpush.ThirdPartyOptions{
	Distribution: "jpush",
	LargeIcon:    "https://example.com/icon.png",
}
o := &jpush.Options{}
o.AddHmosChannel(opts) // convenience wrapper for HMOS channel
// or: o.AddThirdPartyChannel(jpush.HMOS_CHANNEL, opts)
```

## SMS
Template SMS send is available (see `sms_test.go` for an end-to-end example). Fill your own `appKey`/`masterSecret` and template params before running tests.

