package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	tr := TrunkNumber{
		ID:           1,
		EntId:        3,
		SipToUser:    "dfd不能更新",
		NumberType:   1,
		ExpirationAt: time.Now(),
		Purpose:      3,
		IsDefault:    0,
		ChargeMode:   2,
		Prefix:       "Prefix",
		Gateway:      "Gateway",
		Region:       "Region",
		EffectiveAt:  time.Now(),
	}
	mb, _ := json.Marshal(tr)
	fmt.Println(string(mb))

}

// 中继号码
type TrunkNumber struct {
	ID           int64     `json:"id"`
	EntId        int64     `json:"ent_id"`
	SipToUser    string    `json:"sip_to_user"`
	NumberType   int       `json:"number_type"`
	ExpirationAt time.Time `json:"expiration_at"`
	Purpose      int       `json:"purpose"`
	IsDefault    int8      `json:"is_default"`
	ChargeMode   int       `json:"charge_mode"` // 计费方式
	Prefix       string    `json:"prefix"`
	Gateway      string    `json:"gateway"`
	Region       string    `json:"region"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	EffectiveAt  time.Time `json:"effective_at"`
}
