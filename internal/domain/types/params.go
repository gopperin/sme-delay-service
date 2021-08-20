package types

import ()

// DelayMsg 延迟消息必要信息
type DelayMsg struct {
	ID       string `json:"id"`       // 唯一校验
	Delay    int    `json:"delay"`    // 延迟时间，秒
	Callback string `json:"callback"` // 回调
}
