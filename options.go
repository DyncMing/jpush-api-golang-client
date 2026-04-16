package jpush

type Options struct {
	SendNo            int               `json:"sendno,omitempty"`              //推送序号
	TimeToLive        int               `json:"time_to_live,omitempty"`        //离线消息保留时长(秒)
	OverrideMsgId     int               `json:"override_msg_id,omitempty"`     //要覆盖的消息 ID
	ApnsProduction    bool              `json:"apns_production"`               //APNs 是否生产环境
	ApnsCollapseId    string            `json:"apns_collapse_id,omitempty"`    //更新 iOS 通知的标识符
	BigPushDuration   int               `json:"big_push_duration,omitempty"`   //定速推送时长(分钟)
	ThirdPartyChannel ThirdPartyChannel `json:"third_party_channel,omitempty"` //推送请求下发通道
}

type ThirdChannelType string

func (t ThirdChannelType) String() string {
	return string(t)
}

const (
	XIAOMI       ThirdChannelType = "xiaomi"
	HUAWEI       ThirdChannelType = "huawei"
	MEIZU        ThirdChannelType = "meizu"
	OPPO         ThirdChannelType = "oppo"
	VIVO         ThirdChannelType = "vivo"
	FCM          ThirdChannelType = "fcm"
	HMOS_CHANNEL ThirdChannelType = "hmos"
)

type ThirdPartyOptions struct {
	Distribution          string      `json:"distribution,omitempty"`           //通知栏消息下发逻辑
	DistributionFcm       string      `json:"distribution_fcm,omitempty"`       //通知栏消息fcm+国内厂商组合类型下发逻辑
	DistributionCustomize string      `json:"distribution_customize,omitempty"` //自定义消息国内厂商类型下发逻辑
	ChannelId             string      `json:"channel_id"`                       //通知栏消息分类
	SkipQuota             bool        `json:"skip_quota"`                       //配额判断及扣除, 目前仅对小米和oppo有效
	Classification        int         `json:"classification,omitempty"`         //通知栏消息分类, 为了适配 vivo 手机厂商通知栏消息分类,“0”代表运营消息，“1”代表系统消息
	PushMode              int         `json:"push_mode,omitempty"`              //通知栏消息类型, 对应 vivo 的 pushMode 字段,值分别是：“0”表示正式推送；“1”表示测试推送，不填默认为0
	Importance            string      `json:"importance,omitempty"`             //华为通知栏消息智能分类, 为了适配华为手机厂商的通知栏消息智能分类
	Urgency               string      `json:"urgency,omitempty"`                //华为厂商自定义消息优先级, 为了适配华为手机厂商自定义消息的优先级
	Category              string      `json:"category,omitempty"`               //华为厂商自定义消息场景标识
	LargeIcon             string      `json:"large_icon,omitempty"`             //厂商消息大图标样式, 目前支持小米/华为/oppo三个厂商
	SmallIconUri          string      `json:"small_icon_uri,omitempty"`         //厂商消息小图标样式, 目前支持小米/华为两个厂商
	SmallIconColor        string      `json:"small_icon_color,omitempty"`       //小米厂商小图标样式颜色
	Style                 int         `json:"style,omitempty"`                  //厂商消息大文本/inbox/大图片样式
	BigText               string      `json:"big_text,omitempty"`               //厂商消息大文本样式
	Inbox                 interface{} `json:"inbox,omitempty"`                  //厂商消息inbox样式, 目前支持华为厂商
	BigPicPath            string      `json:"big_pic_path,omitempty"`           //厂商big_pic_path, 为了适配厂商的消息大图片样式,目前支持小米/oppo两个厂商
	OnlyUseVendorStyle    bool        `json:"only_use_vendor_style,omitempty"`  //是否是否使用自身通道设置样式
	CallbackId            string      `json:"callback_id,omitempty"`            // vivo厂商通道回调ID
	// HMOS specific
	HmPayload         interface{} `json:"hm_payload,omitempty"`         // 鸿蒙卡片/实况窗消息体
	ReceiptId         string      `json:"receipt_id,omitempty"`         // 华为/鸿蒙回执 ID
	PushType          int         `json:"push_type,omitempty"`          // 鸿蒙 push_type
	TestMessage       bool        `json:"test_message,omitempty"`       // 鸿蒙 test_message 标识
	BadgeAddNumHmos   int         `json:"badge_add_num,omitempty"`      // 鸿蒙角标累加值
	BadgeSetNumHmos   int         `json:"badge_set_num,omitempty"`      // 鸿蒙角标设置值
	ExtraData         string      `json:"extra_data,omitempty"`         // 鸿蒙 extra_data 字段
	DisplayForeground string      `json:"display_foreground,omitempty"` // 鸿蒙前台显示策略
	Sound             string      `json:"sound,omitempty"`              // 鸿蒙自定义铃声
	SoundDuration     int         `json:"sound_duration,omitempty"`     // 鸿蒙自定义铃声时长（秒）
}

type ThirdPartyChannel map[string]ThirdPartyOptions

// SetSendNo 设置消息的发送编号，用来覆盖推送时由 JPush 生成的编号。
func (o *Options) SetSendNo(sendNo int) {
	o.SendNo = sendNo
}

// SetTimeToLive 设置消息的有效期，单位为秒。
func (o *Options) SetTimeToLive(timeToLive int) {
	o.TimeToLive = timeToLive
}

// SetOverrideMsgId 设置覆盖推送时由 JPush 生成的消息 ID。
func (o *Options) SetOverrideMsgId(overrideMsgId int) {
	o.OverrideMsgId = overrideMsgId
}

// SetApnsProduction 设置推送时 APNs 是否生产环境。
func (o *Options) SetApnsProduction(apnsProduction bool) {
	o.ApnsProduction = apnsProduction
}

// SetBigPushDuration 设置大推送时长，单位为秒。
func (o *Options) SetBigPushDuration(bigPushDuration int) {
	o.BigPushDuration = bigPushDuration
}

// AddThirdPartyChannel 添加第三方渠道。
func (o *Options) AddThirdPartyChannel(channel ThirdChannelType, value ThirdPartyOptions) {
	if o.ThirdPartyChannel == nil {
		o.ThirdPartyChannel = make(ThirdPartyChannel)
	}
	o.ThirdPartyChannel[channel.String()] = value
}

// AddHmosChannel is a convenience wrapper to add HMOS (HarmonyOS) third-party options.
func (o *Options) AddHmosChannel(value ThirdPartyOptions) {
	o.AddThirdPartyChannel(HMOS_CHANNEL, value)
}
