package types

import ()

// DelayMsg DelayMsg
type DelayMsg struct {
	Delay    int    `json:"delay"`    // 延迟时间
	Callback string `json:"callback"` // 回调
	Body     string `json:"body"`     // 回调json body
}
